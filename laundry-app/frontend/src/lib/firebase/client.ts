import { initializeApp } from 'firebase/app';
import { getAuth } from 'firebase/auth';

const firebaseConfig = {
  apiKey: import.meta.env.VITE_FIREBASE_API_KEY || 'XXX',
  authDomain: import.meta.env.VITE_FIREBASE_AUTH_DOMAIN || 'XXX',
  projectId: import.meta.env.VITE_FIREBASE_PROJECT_ID || 'XXX',
  appId: import.meta.env.VITE_FIREBASE_APP_ID || 'XXX'
};

const app = initializeApp(firebaseConfig);
export const auth = getAuth(app);
