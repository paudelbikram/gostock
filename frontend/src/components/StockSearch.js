import React, { useState, useEffect } from 'react';
import { TextField, Button, Box, Autocomplete, Typography } from '@mui/material';
import { useDispatch, useSelector } from 'react-redux';
import { fetchStockData,  fetchStockList} from '../store/stockSlice';

const StockSearch = () => {
  const [symbol, setSymbol] = useState('');
  const dispatch = useDispatch();
  
  const stockList = useSelector((state) => state.stock.list || []);

  useEffect(() => {
    dispatch(fetchStockList());
  }, [dispatch]);

  // Automatically run example fetch on component load
  useEffect(() => {
    const exampleSymbol = 'NVDA';
    setSymbol(exampleSymbol);
    dispatch(fetchStockData(exampleSymbol));
  }, [dispatch]);

  const handleSubmit = () => {
    if (symbol.trim()) {
      dispatch(fetchStockData(symbol.toUpperCase()));
    }
  };

  return (
    <Box> 
    <Box display="flex" flexDirection={{ xs: 'column', sm: 'row' }} gap={2} p={2}>
      <Autocomplete
        freeSolo
        options={stockList}
        inputValue={symbol}
        onInputChange={(_, newInputValue) => setSymbol(newInputValue)}
        renderInput={(params) => (
          <TextField {...params} label="Enter Stock Symbol" variant="outlined" fullWidth />
        )}
        fullWidth
      />
      <Button variant="contained" onClick={handleSubmit} fullWidth sx={{ minWidth: '120px' }}>
        Go
      </Button>
    </Box>
    <Typography color="text.secondary" mb={1} textAlign="center">
      Explore key financial data about companies in a clear and easy-to-understand format. Try typing a stock symbol like <b>AAPL</b>, <b>MSFT</b>, or <b>NVDA</b>.
    </Typography>
    </Box>
  );
};

export default StockSearch;
