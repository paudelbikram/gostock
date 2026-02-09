import { Container } from '@mui/material';
import StockSearch from './StockSearch';
import StockResult from './StockResult';
import Header from './Header';
import Footer from './Footer';

const Home = () => {
  return (
    <Container>
      <Header />
      <StockSearch />
      <StockResult />
      <Footer />
    </Container>
  );
};

export default Home;
