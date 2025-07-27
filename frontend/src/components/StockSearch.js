import React, { useState } from 'react';
import { TextField, Button, Box } from '@mui/material';
import { useDispatch } from 'react-redux';
import { fetchStock } from '../store/stockSlice';

const StockSearch = () => {
  const [symbol, setSymbol] = useState('');
  const dispatch = useDispatch();

  const handleSubmit = () => {
    if (symbol.trim()) {
      dispatch(fetchStock(symbol.toUpperCase()));
    }
  };

  return (
    <Box display="flex" flexDirection={{ xs: 'column', sm: 'row' }} gap={2} p={2}>
      <TextField
        label="Enter Stock Symbol"
        variant="outlined"
        value={symbol}
        onChange={(e) => setSymbol(e.target.value)}
        fullWidth
      />
      <Button variant="contained" onClick={handleSubmit} fullWidth sx={{ minWidth: '120px' }}>
        Go
      </Button>
    </Box>
  );
};

export default StockSearch;
