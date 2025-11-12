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
import { formatSmartNumber } from '../utils/numberFormat';

const StockChart = ({data, label}) => {
  const dataWithChange = data.map((item, index, arr) => {
    if (index === 0) return { ...item, change: 0 };
    const prev = arr[index - 1].value;
    const diff = item.value - prev;
    const pct = (diff / Math.abs(prev)) * 100;
    return { ...item, change: pct };
  });
  return (
    <Box p={2}>
      <Typography variant="h5" gutterBottom>{label}</Typography>
      {/* Responsive Chart */}
        <ResponsiveContainer width="100%" height={300}>
          <LineChart data={dataWithChange}>
            <CartesianGrid />
            <XAxis dataKey="key" />
            <YAxis tickFormatter={formatSmartNumber} />
            <Tooltip formatter={(value, name, props) => 
            [`${formatSmartNumber(value)} (${props.payload.change.toFixed(1)}%)`,
              'Value',
            ]}/>
            <Line type="monotone" dataKey="value" stroke="#3f51b5"
            dot={({ cx, cy, payload }) => (
              <circle
                cx={cx}
                cy={cy}
                r={4}
                fill={payload.change >= 0 ? 'green' : 'red'}
              />
            )}
            />
          </LineChart>
        </ResponsiveContainer>
    </Box>
  );
};

export default StockChart;
