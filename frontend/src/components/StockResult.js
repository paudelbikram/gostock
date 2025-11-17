import React from 'react';
import { useSelector } from 'react-redux';
import {
  Box,
  Typography,
  CircularProgress,
} from '@mui/material';
import OverviewCard from './OverviewCard';
import StockChart from './StockChart';
import StockTable from './StockTable';
import StockHealth from './StockHealth';
import DYOR from './DYOR';
import DCFCalculator from './DCFCalculator';

const determineYears = (data) => {
  const five = 5;
  const three = 3;
  const { balancesheet } = data;
  // let's based this off of balancesheet 
  const balanceSheetYears = balancesheet.annualReports.length;
  if (balanceSheetYears > 0) {
    if (balanceSheetYears >= five) {
      return [five, three];
    } else if (balanceSheetYears >= three) {
      return [three];
    } else {
      return [balanceSheetYears];
    }
  }
  return [];
}

const calculateStockHealth = (years, data) => {
  if (!data) return [];

  const { balancesheet, cashflow, income, overview } = data;

  const annualReportsCashflow = cashflow.annualReports.slice(0, years).reverse();
  const annualReportsIncome = income.annualReports.slice(0, years).reverse();
  const annualReportsBalance = balancesheet.annualReports.slice(0, years).reverse();

  // 1. Year PE
  const totalNetIncome = annualReportsIncome.reduce((sum, r) => sum + Number(r.netIncome || 0), 0);
  const PE = overview && totalNetIncome ? overview.MarketCapitalization / totalNetIncome : null;

  // 2. ROIC = Avg(FCF / (Equity + Debt))
  const roicValues = annualReportsCashflow.map((r, i) => {
    const fcf = Number(r.operatingCashflow || 0) - Math.abs(Number(r.capitalExpenditures || 0));
    const equity = Number(annualReportsBalance[i]?.totalShareholderEquity || 0);
    const debt = Number(annualReportsBalance[i]?.longTermDebt || 0);
    return equity + debt ? fcf / (equity + debt) : 0;
  });
  const ROIC = roicValues.length ? roicValues.reduce((a,b)=>a+b,0)/roicValues.length : null;

  // 3. Shares Outstanding Trend
  const sharesCurrent = Number(annualReportsBalance[years-1]?.commonStockSharesOutstanding || 0);
  const sharesPast = Number(annualReportsBalance[0]?.commonStockSharesOutstanding || 0);
  const sharesTrend = sharesCurrent / sharesPast;

  // 4. Cash Flow Growth
  const fcfCurrent = Number(annualReportsCashflow[years-1]?.operatingCashflow || 0) - Math.abs(Number(annualReportsCashflow[years-1]?.capitalExpenditures || 0));
  const fcfPast = Number(annualReportsCashflow[0]?.operatingCashflow || 0) - Math.abs(Number(annualReportsCashflow[0]?.capitalExpenditures || 0));
  const cashFlowGrowth = fcfPast ? (fcfCurrent - fcfPast)/fcfPast*100 : null;

  // 5. Net Income Growth
  const netIncomeCurrent = Number(annualReportsIncome[years-1]?.netIncome || 0);
  const netIncomePast = Number(annualReportsIncome[0]?.netIncome || 0);
  const netIncomeGrowth = netIncomePast ? (netIncomeCurrent - netIncomePast)/netIncomePast*100 : null;

  // 6. Revenue Growth
  const revenueCurrent = Number(annualReportsIncome[years-1]?.totalRevenue || 0);
  const revenuePast = Number(annualReportsIncome[0]?.totalRevenue || 0);
  const revenueGrowth = revenuePast ? (revenueCurrent - revenuePast)/revenuePast*100 : null;

  // 7. Debt vs FCF
  const longTermDebt = (annualReportsBalance[years - 1]?.longTermDebt === "None") ? 0
                      : Number(annualReportsBalance[years - 1]?.longTermDebt || 0);
  const sumFCF = annualReportsCashflow.reduce((sum, r) => sum + (Number(r.operatingCashflow || 0) - Math.abs(Number(r.capitalExpenditures || 0))), 0);
  const debtVsFCF = sumFCF ? longTermDebt / sumFCF : null;

  // 8. P/FCF
  const marketCap = Number(overview?.MarketCapitalization || 0);
  const pfcf = sumFCF ? marketCap / sumFCF : null;

  return [
    { name: `${years}-Year PE Ratio`, 
      description: `Sum of net income over ${years} years divided by current market capitalization.
                    Anything below 20 (historical average) is considered undervalued.`,
      value: PE?.toFixed(2), 
      check: PE < 20 
    },
    { name: `${years}-Year ROIC`, 
      value: (ROIC*100)?.toFixed(2) + '%', 
      description: `Return on invested capital over ${years} years.
                     ROIC over 10 considers a value creator. `, 
      check: ROIC > 0.1 
    },
    { name: 'Shares Outstanding Trend', 
      description: `Number of shares outstanding compared to ${years} years ago.
                     Anything above 1 indicates shares dilution.`,
      value: sharesTrend.toFixed(2), 
      check: sharesTrend <= 1 
    },
    { name: `Cash Flow Growth (${years}Y)`, 
      value: cashFlowGrowth?.toFixed(2)+'%', 
      description: `Cash flow compared to ${years} years ago.
                    Anything above 0 indicates increasing cash flow.`,
      check: cashFlowGrowth > 0 
    },
    { name: `Net Income Growth (${years}Y)`, 
      description: `Net Income compared to ${years} years ago.
                    Anything above 0 indicates increasing net income.`,
      value: netIncomeGrowth?.toFixed(2)+'%', 
      check: netIncomeGrowth > 0 
    },
    { name: `Revenue Growth (${years}Y)`, 
      description: `Total revenue compared to ${years} years ago.
                    Anything above 0 indicates increasing revenue.`,
      value: revenueGrowth?.toFixed(2)+'%', 
      check: revenueGrowth > 0 
    },
    { name: 'Debt vs FCF', 
      value: debtVsFCF?.toFixed(2), 
      description: `Current long term debt devided by sum of free cash flow over ${years} years.
                    Anything equal to or above 5 is considered risky.`,
      check: debtVsFCF < 5 
    },
    { name: `${years}-Year P/FCF`,
      description: `Current market capitalization divided by sum of free cash flow over ${years} years.
                    Anything below 20 considered undervalued.`, 
      value: pfcf?.toFixed(2), 
      check: pfcf < 20 
    },
  ];
};

const StockResult = () => {
  const { data, loading, error } = useSelector((state) => state.stock);
  if (loading) return <Box p={2}><CircularProgress /></Box>;
  if (error) return <Box p={2}><Typography color="error">{error}</Typography></Box>;
  if (!data) return null;
  const balanceSheetColumns = [
    {'id': 'fiscalDateEnding', 'label': 'Fiscal Date Ending'},
    {'id': 'totalAssets', 'label': 'Total Assets'},
    {'id': 'cashAndCashEquivalentsAtCarryingValue', 'label': 'Cash And Cash Equivalents At Carrying Value'},
    {'id': 'totalLiabilities', 'label': 'Total Liabilities'},
    {'id': 'totalShareholderEquity', 'label': 'Total Shareholder Equity'},
    {'id': 'inventory', 'label': 'Inventory'},
    {'id': 'deferredRevenue', 'label': 'Deferred Revenue'},
    {'id': 'totalCurrentLiabilities', 'label': 'Total Current Liabilities'},
    {'id': 'longTermDebt', 'label': 'Long Term Debt'},
    {'id': 'retainedEarnings', 'label': 'Retained Earnings'},
    {'id': 'commonStockSharesOutstanding', 'label': 'Common Stock Shares Outstanding'},
  ];

  const cashflowColumns = [
    {'id': 'fiscalDateEnding', 'label': 'Fiscal Date Ending'},
    {'id': 'operatingCashflow', 'label': 'Operating Cashflow'},
    {'id': 'capitalExpenditures', 'label': 'Capital Expenditures'},
    {'id': 'dividendPayout', 'label': 'Dividend Payout'},
    {'id': 'cashflowFromFinancing', 'label': 'Cashflow From Financing'},
    {'id': 'cashflowFromInvestment', 'label': 'Cashflow From Investment'},
    {'id': 'netIncome', 'label': 'Net Income'},
  ];

  const incomeColumns = [
    {'id': 'fiscalDateEnding', 'label': 'Fiscal Date Ending'},
    {'id': 'totalRevenue', 'label': 'Total Revenue'},
    {'id': 'grossProfit', 'label': 'Gross Profit'},
    {'id': 'operatingIncome', 'label': 'Operating Income'},
    {'id': 'netIncome', 'label': 'Net Income'},
    {'id': 'ebitda', 'label': 'EBITDA'},
    {'id': 'operatingExpenses', 'label': 'Operating Expenses'},
    {'id': 'incomeBeforeTax', 'label': 'Income Before Tax'},
  ];
  const years = determineYears(data);

  return (
    <Box>
        <DYOR />
        <OverviewCard overview={data.overview} />
        {years.map(year =>
        <StockHealth healthData={calculateStockHealth(year, data)} years={year} />
        )}
        <DYOR />
        <DCFCalculator />
        <StockChart data={[...data.revenueTrend.Yearly].reverse()} 
        label={data.ticker + ' Yearly Revenue'}/>
        <StockChart data={[...data.revenueTrend.Quarterly].reverse()} 
        label={data.ticker + ' Quarterly Revenue'}/>
        <DYOR />
        <StockChart data={[...data.cashflowTrend.Yearly].reverse()} 
        label={data.ticker + ' Yearly Cashflow'}/>
        <StockChart data={[...data.cashflowTrend.Quarterly].reverse()} 
        label={data.ticker + ' Quarterly Cashflow'}/>
        <DYOR />
        <StockChart data={[...data.profitMarginTrend.Yearly].reverse()} 
        label={data.ticker + ' Yearly Profit Margin'}/>
        <StockChart data={[...data.profitMarginTrend.Quarterly].reverse()} 
        label={data.ticker + ' Quarterly Profit Margin'}/>
        <DYOR />
        <StockChart data={[...data.operatingMarginTrend.Yearly].reverse()} 
        label={data.ticker + ' Yearly Operating Margin'}/>
        <StockChart data={[...data.operatingMarginTrend.Quarterly].reverse()} 
        label={data.ticker + ' Quarterly Operating Margin'}/>
        <DYOR />
        <StockChart data={[...data.debt2equityRatioTrend.Yearly].reverse()} 
        label={data.ticker + ' Yearly Debt To Equity Ratio'}/>
        <StockChart data={[...data.debt2equityRatioTrend.Quarterly].reverse()} 
        label={data.ticker + ' Quarterly Debt To Equity Ratio'}/>
        <DYOR />
        <StockTable columns={balanceSheetColumns} rows={data.balancesheet.annualReports} label={data.ticker + ' Yearly Balancesheet'}/>
        <StockTable columns={balanceSheetColumns} rows={data.balancesheet.quarterlyReports} label={data.ticker + ' Quarterly Balancesheet'}/>
        <DYOR />
        <StockTable columns={cashflowColumns} rows={data.cashflow.annualReports} label={data.ticker + ' Yearly Cashflow'}/>
        <StockTable columns={cashflowColumns} rows={data.cashflow.quarterlyReports} label={data.ticker + ' Quarterly Cashflow'}/>
        <DYOR />
        <StockTable columns={incomeColumns} rows={data.income.annualReports} label={data.ticker + ' Yearly Income'}/>
        <StockTable columns={incomeColumns} rows={data.income.quarterlyReports} label={data.ticker + ' Quarterly Income'}/>
    </Box>
  );
};

export default StockResult;
