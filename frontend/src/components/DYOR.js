import React from 'react';
import { Typography } from '@mui/material';
import WarningIcon from '@mui/icons-material/Warning';


const DYOR = () => {
  return (
    <Typography color="text.secondary" textAlign="center">
        <WarningIcon fontSize="medium" color="error" sx={{ verticalAlign: 'middle' }}/> 
        This analysis is for educational purposes only and should not be considered financial advice.
        Always do your own research.
    </Typography>
  );
};

export default DYOR;
