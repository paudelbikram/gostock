import React from 'react';
import {
  Box,
  Typography,
} from '@mui/material';
import {
  LineChart,
  Line,
  XAxis,
  YAxis,
  CartesianGrid,
  Tooltip,
  ResponsiveContainer,
} from 'recharts';
import { formatSmartMoney, formatSmartNumber } from '../utils/numberFormat';

const StockChart = ({data, label}) => {
  return (
    <Box p={2}>
      <Typography variant="h5" gutterBottom>{label}</Typography>
      {/* Responsive Chart */}
        <ResponsiveContainer width="100%" height={300}>
          <LineChart data={data}>
            <CartesianGrid />
            <XAxis dataKey="key" />
            <YAxis tickFormatter={formatSmartNumber} />
            <Tooltip formatter={formatSmartNumber} />
            <Line type="monotone" dataKey="value" stroke="#3f51b5" />
          </LineChart>
        </ResponsiveContainer>
    </Box>
  );
};

export default StockChart;
