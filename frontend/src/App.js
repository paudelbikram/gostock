import { BrowserRouter, Routes, Route } from "react-router-dom";
import Home from './components/Home';
import Login from './components/Login';
import RequireAuth from "./auth/RequireAuth"
import PublicRoute from "./auth/PublicRoute";


function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/login" 
          element={
            <PublicRoute>
              <Login />
              </PublicRoute>} 
        />
        <Route
          path="/*"
          element={
            <RequireAuth>
              <Home />
            </RequireAuth>
          }
        />
      </Routes>
    </BrowserRouter>
  );
}

export default App;
