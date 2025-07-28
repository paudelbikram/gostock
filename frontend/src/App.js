import React from 'react';
import StockSearch from './components/StockSearch';
import { Container, Typography } from '@mui/material';
import StockResult from './components/StockResult';
import Logo from './logo.svg';

function App() {
  return (
    <Container>
      <Typography variant="h4" gutterBottom align="center">
        <img src={Logo} alt="GoStock" width={120} />
      </Typography>
      <StockSearch />
      <StockResult />
    </Container>
  );
}

export default App;
