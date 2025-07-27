import React from 'react';
import StockSearch from './components/StockSearch';
import { Container, Typography } from '@mui/material';
import StockResult from './components/StockResult';

function App() {
  return (
    <Container>
      <Typography variant="h4" gutterBottom align="center">
        GoStock
      </Typography>
      <StockSearch />
      <StockResult />
    </Container>
  );
}

export default App;
