import { Button, Card, Container, Stack} from "@mui/material"
import {
  GoogleAuthProvider,
  signInWithPopup,
} from "firebase/auth"
import { useNavigate } from 'react-router-dom';
import { auth } from "../firebase/firebase"
import { useDispatch } from "react-redux"
import { loginWithToken } from "../store/authSlice"
import Header from "./Header"
import Footer from "./Footer"

export default function Login() {
  const dispatch = useDispatch();
  const navigate = useNavigate();
  const login = async (provider) => {
    const result = await signInWithPopup(auth, provider)
    const token = await result.user.getIdToken();
    try {
        await dispatch(loginWithToken(token)).unwrap();    
        navigate("/");
    } catch (err) {
        console.error("Failed to login:", err);
    }
  }

  return (
    <Container maxWidth="sm" sx={{ mt: 10 }}>
      <Header />
      <Card sx={{ p: 4 }}>
        <Stack spacing={2} mt={3}>
          <Button variant="contained" onClick={() => login(new GoogleAuthProvider())}>
            Continue with Google
          </Button>
        </Stack>
      </Card>
      <Footer />
    </Container>
  )
}
