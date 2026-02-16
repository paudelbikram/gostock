import { useState } from 'react';
import { useSelector, useDispatch } from 'react-redux';
import { Link } from 'react-router-dom';
import { Box, Typography, Button, Menu, Container } from '@mui/material';
import KeyboardArrowDownIcon from '@mui/icons-material/KeyboardArrowDown';
import LogoutIcon from '@mui/icons-material/Logout';
import Logo from '../logo.svg';

import { useNavigate } from 'react-router-dom';
import { logoutUser } from "../store/authSlice"

const Header = () => {
  const dispatch = useDispatch();
  const navigate = useNavigate();
  const { user } = useSelector((state) => state.auth);
  
  const [anchorEl, setAnchorEl] = useState(null);
  const open = Boolean(anchorEl);

  const handleOpen = (event) => setAnchorEl(event.currentTarget);
  const handleClose = () => setAnchorEl(null);

  const handleLogout = () => {
    handleClose();
    dispatch(logoutUser());
    navigate('/login');
  };

  return (
    <Container maxWidth="lg">
  <Box
    sx={{
      display: "flex",
      justifyContent: "space-between",
      alignItems: "center",
      py: { xs: 1.5, sm: 2 },
    }}
  >
    {/* Logo */}
    <Box>
      <Typography component={Link} to="/">
        <img
          src={Logo}
          alt="GoStock"
          style={{
            display: "block",
            width: "100%",
            maxWidth: 120,
          }}
        />
      </Typography>
    </Box>

    {/* Right Side */}
    <Box
      sx={{
        display: "flex",
        alignItems: "center",
        gap: { xs: 1, sm: 2 },
      }}
    >
      {user && (
        <>
          {/* Hide email on mobile */}
          <Typography
            variant="body2"
            sx={{
              display: { xs: "none", sm: "block" },
              fontWeight: 500,
              color: "text.secondary",
              maxWidth: 180,
              whiteSpace: "nowrap",
              overflow: "hidden",
              textOverflow: "ellipsis",
            }}
          >
            {user.email}
          </Typography>

          {/* Account Button */}
          <Button
            size="small"
            onClick={handleOpen}
            endIcon={<KeyboardArrowDownIcon />}
            variant="contained"
            sx={{
              textTransform: "none",
              fontWeight: 600,
              borderRadius: "8px",
              px: { xs: 1.5, sm: 2 },
              minWidth: { xs: 40, sm: "auto" },
            }}
          >
            {/* Hide text on mobile */}
            <Box sx={{ display: { xs: "none", sm: "inline" } }}>
              Account
            </Box>
          </Button>

          <Menu
            anchorEl={anchorEl}
            open={open}
            onClose={handleClose}
            transformOrigin={{ horizontal: "right", vertical: "top" }}
            anchorOrigin={{ horizontal: "right", vertical: "bottom" }}
            PaperProps={{ sx: { mt: 1, minWidth: 150 } }}
          >
            <Box sx={{ px: 2, py: 1 }}>
              <Button
                fullWidth
                variant="contained"
                onClick={handleLogout}
                startIcon={<LogoutIcon fontSize="small" />}
                sx={{
                  textTransform: "none",
                  fontWeight: 600,
                  borderRadius: "8px",
                }}
              >
                Logout
              </Button>
            </Box>
          </Menu>
        </>
      )}
    </Box>
  </Box>
</Container>

  );
};

export default Header;
