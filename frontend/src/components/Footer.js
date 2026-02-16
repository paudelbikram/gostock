import { Typography, Box } from '@mui/material';

const Footer = () => {
  return (
    <Box sx={{ display: 'flex', justifyContent: 'center', m: 2 }}>
        <Typography variant="body2" color="text.secondary">
            © {new Date().getFullYear()} TechPasya <br/>
            <b>Disclaimer</b>: This content is for informational purposes only and does not constitute financial, investment, or legal advice. 
            Use at your own risk.
        </Typography>
    </Box>
  );
};

export default Footer;
