import { createSlice, createAsyncThunk } from "@reduxjs/toolkit"
import axios from "axios"
import { signOut } from "firebase/auth"
import { auth } from "../firebase/firebase"

const API_BASE = process.env.REACT_APP_API_BASE || "";

export const loginWithToken = createAsyncThunk(
  "auth/login",
  async (idToken, {rejectWithValue}) => {
    try {
      const res = await axios.post(
        `${API_BASE}/auth/login`,
        {},
        { headers: { Authorization: `Bearer ${idToken}` } }
      );
      return res.data;
    } catch (err) {
      return rejectWithValue(err.response?.data || "Login failed");
    } 
  }
);

export const logoutUser = createAsyncThunk(
  "auth/logout",
  async (_, { rejectWithValue }) => {
    try {   
      // 1. Tell Firebase to invalidate the client session
      await signOut(auth);   
      // No backend call needed since it is token based login
      return true;
    } catch (err) {
      return rejectWithValue(err.message);
    }
  }
);

const authSlice = createSlice({
  name: "auth",
  initialState: { user: null, status: "idle", error: null },
  reducers: {
    // Optional: local reset if needed without a thunk
    clearAuthState(state) {
      state.user = null;
      state.status = "idle";
    }
  },
  extraReducers(builder) {
    builder
      .addCase(loginWithToken.pending, (state) => {
        state.status = "loading";
        state.error = null;
      })
      .addCase(loginWithToken.fulfilled, (state, action) => {
        state.user = action.payload.user
        state.status = "authenticated";
        state.error = null;
      })
      .addCase(loginWithToken.rejected, (state, action) => {
        state.status = "failed";
        state.error = action.payload;
      })
      .addCase(logoutUser.fulfilled, (state) => {
        state.user = null;
        state.status = "idle";
      });
      
  },
})

export const { clearAuthState } = authSlice.actions
export default authSlice.reducer
