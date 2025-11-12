import React from 'react';
import StockSearch from './components/StockSearch';
import { Container, Typography, Box } from '@mui/material';
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
      <Box sx={{ display: 'flex', justifyContent: 'center', m: 2 }}>
        <Typography variant="body2" color="text.secondary">
          © 2025 TechPasya — For informational purposes only. Not financial advice.
        </Typography>
      </Box>
    </Container>
  );
}

export default App;
