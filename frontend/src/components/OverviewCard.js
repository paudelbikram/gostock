import React from 'react';
import {
  Box, Typography, Grid, Paper, Divider, Link
} from '@mui/material';

const fields = [
  ['Name', 'OfficialSite', 'Exchange', 'Description'],
  ['Country', 'Industry', 'AssetType', 'Sector'],
  ['MarketCapitalization', 'EPS', 'BookValue', 'PriceToBookRatio'],
  ['PERatio', 'ForwardPE', 'TrailingPE', 'PEGRatio'],
  ['DividendYield', 'DividendDate', 'ExDividendDate', 'DividendPerShare'],
  ['EBITDA', 'EVToEBITDA', 'EVToRevenue', 'Currency'],
  ['Beta' , 'RevenuePerShareTTM', 'SharesOutstanding' , 'ReturnOnEquityTTM', 'ReturnOnAssetsTTM'],
  ['ProfitMargin', 'FiscalYearEnd', 'LatestQuarter', 'QuarterlyEarningsGrowthYOY', 'QuarterlyRevenueGrowthYOY'],
  ['RevenueTTM', 'GrossProfitTTM', 'DilutedEPSTTM', 'PriceToSalesRatioTTM', 'OperatingMarginTTM'],
  ['52WeekHigh', '52WeekLow', '50DayMovingAverage', '200DayMovingAverage', 'AnalystTargetPrice'],
  ['AnalystRatingStrongBuy', 'AnalystRatingBuy', 'AnalystRatingHold', 'AnalystRatingSell', 'AnalystRatingStrongSell'],
];

const formatNumber = (val) => {
  const num = Number(val);
  return isNaN(num) ? val : Number(val).toLocaleString();
};

const OverviewCard = ({overview}) => {
  return (
    <Box p={2}>
      <Typography variant="h5" gutterBottom>{overview.Name} Overview</Typography>
      <Grid container spacing={2}>
        {fields.map((group, i) => (
          <Grid item xs={12} md={6} key={i}>
            <Paper elevation={2} sx={{ p: 2 }}>
              {group.map(key => (
                <Box key={key} mb={1}>
                  <Typography variant="body2" fontWeight="bold">{key}</Typography>
                  {key === 'OfficialSite' ? (
                    <Link href={overview[key]} target="_blank" rel="noopener">
                      {overview[key]}
                    </Link>
                  ) : (
                    <Typography variant="body2">
                      {formatNumber(overview[key])}
                    </Typography>
                  )}
                  <Divider sx={{ my: 1 }} />
                </Box>
              ))}
            </Paper>
          </Grid>
        ))}
      </Grid>
    </Box>
  );
};

export default OverviewCard;