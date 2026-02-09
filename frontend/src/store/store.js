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

const rootReducer = combineReducers({
  auth: authReducer,
  stock: stockReducer,
});

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
