import React from 'react';
import { useSelector } from 'react-redux';
import {
  Box,
  Typography,
  CircularProgress,
  Container,
} from '@mui/material';
import OverviewCard from './OverviewCard';
import StockChart from './StockChart';
import StockTable from './StockTable';

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

  return (
    <Container>
        <OverviewCard overview={data.overview} />

        <StockChart data={[...data.revenueTrend.Yearly].reverse()} 
        label={data.ticker + ' Yearly Revenue'}/>
        <StockChart data={[...data.revenueTrend.Quarterly].reverse()} 
        label={data.ticker + ' Quarterly Revenue'}/>

        <StockChart data={[...data.cashflowTrend.Yearly].reverse()} 
        label={data.ticker + ' Yearly Cashflow'}/>
        <StockChart data={[...data.cashflowTrend.Quarterly].reverse()} 
        label={data.ticker + ' Quarterly Cashflow'}/>

        <StockChart data={[...data.profitMarginTrend.Yearly].reverse()} 
        label={data.ticker + ' Yearly Profit Margin'}/>
        <StockChart data={[...data.profitMarginTrend.Quarterly].reverse()} 
        label={data.ticker + ' Quarterly Profit Margin'}/>

        <StockChart data={[...data.operatingMarginTrend.Yearly].reverse()} 
        label={data.ticker + ' Yearly Operating Margin'}/>
        <StockChart data={[...data.operatingMarginTrend.Quarterly].reverse()} 
        label={data.ticker + ' Quarterly Operating Margin'}/>

        <StockChart data={[...data.debt2equityRatioTrend.Yearly].reverse()} 
        label={data.ticker + ' Yearly Debt To Equity Ratio'}/>
        <StockChart data={[...data.debt2equityRatioTrend.Quarterly].reverse()} 
        label={data.ticker + ' Quarterly Debt To Equity Ratio'}/>

        <StockTable columns={balanceSheetColumns} rows={data.balancesheet.annualReports} label={data.ticker + ' Yearly Balancesheet'}/>
        <StockTable columns={balanceSheetColumns} rows={data.balancesheet.quarterlyReports} label={data.ticker + ' Quarterly Balancesheet'}/>
        
        <StockTable columns={cashflowColumns} rows={data.cashflow.annualReports} label={data.ticker + ' Yearly Cashflow'}/>
        <StockTable columns={cashflowColumns} rows={data.cashflow.quarterlyReports} label={data.ticker + ' Quarterly Cashflow'}/>
        
        <StockTable columns={incomeColumns} rows={data.income.annualReports} label={data.ticker + ' Yearly Income'}/>
        <StockTable columns={incomeColumns} rows={data.income.quarterlyReports} label={data.ticker + ' Quarterly Income'}/>
    </Container>
  );
};

export default StockResult;
