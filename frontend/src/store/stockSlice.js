import { createSlice, createAsyncThunk } from '@reduxjs/toolkit';
import axios from 'axios';

export const fetchStockData = createAsyncThunk(
  'stock/fetchData',
  async (symbol) => {
    const response = await axios.get(`http://192.168.1.70:8080/api/${symbol}`);
    return response.data;
  }
);

export const fetchStockList = createAsyncThunk(
  'stock/fetchList',
  async () => {
    const response = await axios.get('http://192.168.1.70:8080/api/stock/list');
    return response.data;
  }
);

const stockSlice = createSlice({
  name: 'stock',
  initialState: {
    data: null,
    loading: false,
    error: null,
    list: [],
  },
  reducers: {},
  extraReducers: (builder) => {
    builder
      .addCase(fetchStockData.pending, (state) => {
        state.loading = true;
        state.error = null;
      })
      .addCase(fetchStockData.fulfilled, (state, action) => {
        state.loading = false;
        state.data = action.payload;
      })
      .addCase(fetchStockData.rejected, (state, action) => {
        state.loading = false;
        state.error = action.error.message;
      })
      .addCase(fetchStockList.fulfilled, (state, action) => {
        state.listLoading = false;
        state.list = action.payload;
      });
  },
});

export default stockSlice.reducer;
