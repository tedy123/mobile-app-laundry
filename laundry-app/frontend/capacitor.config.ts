import type { CapacitorConfig } from '@capacitor/cli';

const config: CapacitorConfig = {
  appId: 'com.example.app',
  appName: 'laundry-app',
  webDir: 'build',
  server: {
    androidScheme: 'http'
  },
  plugins: {
    GoogleAuth: {
      scopes: ['profile', 'email'],
      serverClientId: '971457850261-v3mf97mkgknhciestbrta83tbrh9tm3d.apps.googleusercontent.com',
      forceCodeForRefreshToken: true,
    },
  },
};

export default config;
