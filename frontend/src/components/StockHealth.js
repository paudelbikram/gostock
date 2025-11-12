import React from 'react';
import { Box, Typography } from '@mui/material';
import ArrowUpwardIcon from '@mui/icons-material/ArrowUpward';
import ArrowDownwardIcon from '@mui/icons-material/ArrowDownward';

const StockHealth = ({healthData, years}) => {
  if (!healthData || healthData.length === 0) return null;

  return (
    <Box mb={5} border="1px #ccc" borderRadius={2}>
      <Typography variant="h5" gutterBottom>{years}Y Health Check</Typography>
      {healthData.map((data) => {
        // Determine color based on check
        const color = data.check ? 'green' : 'red';

        // Determine arrow for growth-related pillars
        let arrow = null;
        if (data.name.includes('Growth') || data.name.includes('ROIC')) {
          arrow = data.value && parseFloat(data.value) >= 0 
            ? <ArrowUpwardIcon fontSize="small" sx={{ verticalAlign: 'middle' }}/> 
            : <ArrowDownwardIcon fontSize="small" sx={{ verticalAlign: 'middle' }}/>;
        }
        return (
          <Box 
            key={data.name} 
            display="flex" 
            justifyContent="space-between" 
            alignItems="center" 
            mb={1} 
            p={1} 
            sx={{ backgroundColor: '#f9f9f9', borderRadius: 1 }}
          >
            <Box>
              <Typography variant="subtitle1" fontWeight={600}>
                {data.name}
              </Typography>
              <Typography variant="body2" color="text.secondary">
                {data.description}
              </Typography>
            </Box>
            <Box display="flex" alignItems="center" gap={1}>
              <Typography sx={{ color, fontWeight: 600 }}>{data.value}</Typography>
              {arrow}
              <Typography sx={{ color }}>{data.check ? '✅' : '❌'}</Typography>
            </Box>
          </Box>
        );
      })}
    </Box>
  );
};

export default StockHealth;
