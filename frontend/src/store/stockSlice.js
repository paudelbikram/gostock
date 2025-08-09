import { createSlice, createAsyncThunk } from '@reduxjs/toolkit';
import axios from 'axios';

const API_BASE = process.env.REACT_APP_API_BASE || "";

export const fetchStockData = createAsyncThunk(
  'stock/fetchData',
  async (symbol, { rejectWithValue }) => {
    try {
      const { data } = await axios.get(`${API_BASE}/api/${symbol}`);
      return data;
    } catch (err) {
      if (err.response?.data) {
        return rejectWithValue(
          err.response.data.error || err.response.data.message || "Internal Server Error"
        );
      }
      return rejectWithValue(err.message || "Network error");
    }
  }
);

export const fetchStockList = createAsyncThunk(
  'stock/fetchList',
  async (_, { rejectWithValue }) => {
    try {
      const { data } = await axios.get(`${API_BASE}/api/stock/list`);
      return data;
    } catch (err) {
      if (err.response?.data) {
        return rejectWithValue(
          err.response.data.error || err.response.data.message || "Internal Server Error"
        );
      }
      return rejectWithValue(err.message || "Network error");
    }
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
        state.error = action.payload || action.error.message;
      })
      .addCase(fetchStockList.fulfilled, (state, action) => {
        state.listLoading = false;
        state.list = action.payload;
      });
  },
});

export default stockSlice.reducer;
