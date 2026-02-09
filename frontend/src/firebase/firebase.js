import {initializeApp} from 'firebase/app';
import { getAuth } from 'firebase/auth'

const firebaseConfig = {
  apiKey: "AIzaSyBBfbziEUAGYd9mqbN1xKKgi9UxlblavbQ",
  authDomain: "techpasya-auth.firebaseapp.com",
  projectId: "techpasya-auth",
}

const app = initializeApp(firebaseConfig)
export const auth = getAuth(app)