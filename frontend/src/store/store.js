import { configureStore, combineReducers } from '@reduxjs/toolkit';
import stockReducer from './stockSlice';
import authReducer from './authSlice';
import storage from 'redux-persist/lib/storage'; // defaults to localStorage
import { persistReducer, persistStore } from 'redux-persist';
import { FLUSH, REHYDRATE, PAUSE, PERSIST, PURGE, REGISTER } from 'redux-persist';

const persistConfig = {
  key: 'root',
  storage,
  whitelist: ['auth'], // Only auth will be persisted
};

// 1. Define the combined reducer first
const appReducer = combineReducers({
  auth: authReducer,
  stock: stockReducer,
});

// 2. Create the root reducer that handles the reset
const rootReducer = (state, action) => {
  if (action.type === 'auth/logout/fulfilled') {
    // This clears the state memory
    state = undefined;
    // Optional: Manually clear local storage if you want to be 100% sure
    localStorage.removeItem('persist:root'); 
  }  
  return appReducer(state, action);
};


const persistedReducer = persistReducer(persistConfig, rootReducer);

export const store = configureStore({
  reducer: persistedReducer,
  middleware: (getDefaultMiddleware) =>
    getDefaultMiddleware({
      serializableCheck: {
        // Ignore redux-persist actions to avoid console warnings
        ignoredActions: [FLUSH, REHYDRATE, PAUSE, PERSIST, PURGE, REGISTER],
      },
    }),
});

export const persistor = persistStore(store);
