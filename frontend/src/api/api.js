// src/api/api.js
import axios from 'axios';
import { logoutUser } from "../store/authSlice"
import { useNavigate } from 'react-router-dom';


const api = axios.create({ 
  baseURL: process.env.REACT_APP_API_BASE 
});

let store;
export const injectStore = (_store) => {
  store = _store;
};

api.interceptors.response.use(
  (response) => response,
  (error) => {
    if (error.response && error.response.status === 401) {
      if (store) {
        store.dispatch(logoutUser());
      }
      const navigate = useNavigate();
      navigate('/login');
    }
    return Promise.reject(error);
  }
);

export default api;
