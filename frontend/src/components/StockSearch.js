import React, { useState, useEffect } from 'react';
import { TextField, Button, Box, Autocomplete } from '@mui/material';
import { useDispatch, useSelector } from 'react-redux';
import { fetchStockData,  fetchStockList} from '../store/stockSlice';

const StockSearch = () => {
  const [symbol, setSymbol] = useState('');
  const dispatch = useDispatch();
  
  const stockList = useSelector((state) => state.stock.list || []);

  useEffect(() => {
    dispatch(fetchStockList());
  }, [dispatch]);

  const handleSubmit = () => {
    if (symbol.trim()) {
      dispatch(fetchStockData(symbol.toUpperCase()));
    }
  };

  return (
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
  );
};

export default StockSearch;
