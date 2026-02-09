import React from 'react';
import { Typography } from '@mui/material';
import WarningIcon from '@mui/icons-material/Warning';


const DYOR = () => {
  return (
    <Typography color="text.secondary" textAlign="center">
        <WarningIcon fontSize="medium" color="error" sx={{ verticalAlign: 'middle' }}/> 
        <b>Disclaimer</b>: This content is for informational purposes only and does not constitute financial, investment, or legal advice. 
        Use at your own risk.
    </Typography>
  );
};

export default DYOR;
