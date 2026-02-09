import { Typography, Box } from '@mui/material';
import Logo from '../logo.svg';

const Header = () => {
  return (
    <Box>
        <Typography variant="h4" gutterBottom align="center">
            <img src={Logo} alt="GoStock" width={120} />
        </Typography>
    </Box>
  );
};

export default Header;
