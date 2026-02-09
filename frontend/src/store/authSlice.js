import { createSlice, createAsyncThunk } from "@reduxjs/toolkit"
import axios from "axios"

const API_BASE = process.env.REACT_APP_API_BASE || "";

export const loginWithToken = createAsyncThunk(
  "auth/login",
  async (idToken) => {
    // await axios.get(`${API_BASE}/api/${symbol}`);
    const res = await axios.post(
      `${API_BASE}/auth/login`,
      {},
      { headers: { Authorization: `Bearer ${idToken}` } }
    )
    return res.data
  }
)

const authSlice = createSlice({
  name: "auth",
  initialState: { user: null, status: "idle" },
  reducers: {
    logout(state) {
      state.user = null
    },
  },
  extraReducers(builder) {
    builder
      .addCase(loginWithToken.pending, (state) => {
        state.status = "loading"
      })
      .addCase(loginWithToken.fulfilled, (state, action) => {
        state.user = action.payload.user
        state.status = "authenticated"
      })
  },
})

export const { logout } = authSlice.actions
export default authSlice.reducer
