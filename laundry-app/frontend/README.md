# Laundry App Frontend 👕

A modern, high-performance mobile-first application built with **SvelteKit (Svelte 5)** and **Capacitor**. Designed for laundry business owners to manage operations on Android devices.

## 🚀 Tech Stack

- **Framework:** SvelteKit + Svelte 5 (using Runes: `$state`, `$props`, `$derived`, `$effect`)
- **Styling:** Tailwind CSS (Modern, premium aesthetics)
- **Mobile Integration:** [Capacitor](https://capacitorjs.com/) (Native Android)
- **Authentication:** 
    - Firebase Client SDK (Web fallback)
    - Native Google Auth (`@codetrix-studio/capacitor-google-auth`)
- **Icons:** Lucide Svelte
- **Storage:** Capacitor Preferences (Local token persistence)

## 📦 Key Features

- **Svelte 5 Ready**: Fully migrated to Svelte 5 syntax for maximum performance and reactivity.
- **Native Android Experience**:
    - **native Google Login**: Seamless account selection via device dialogs.
    - **Persistence**: Sessions stay logged in even after closing the app.
- **Interactive Dashboard**:
    - Summary of daily revenue (Pendapatan & Penjualan).
    - Responsive grid menu.
    - Floating notification system.
- **Premium Components**:
    - **Logout Modal**: Elegant popup confirmation when clicking the profile photo.
    - **Empty State**: Beautiful 3D illustrations for features currently in development.
    - **Detail Pages**: Clean layout for promo and information details.

## ⚙️ Configuration (.env)

Create a `.env` file in the `frontend/` directory.

> **CRITICAL**: If testing on an **Android Emulator**, use `10.0.2.2` instead of `localhost`.

```env
VITE_API_URL=http://10.0.2.2:4000/api
VITE_FIREBASE_API_KEY=xxx
VITE_FIREBASE_AUTH_DOMAIN=xxx
VITE_FIREBASE_PROJECT_ID=xxx
VITE_FIREBASE_APP_ID=xxx
```

## 📥 Development & Build

### 1. Web Preview
```bash
npm install
npm run dev
```

### 2. Android Build (Recommended for Testing)
Every time you make changes to Svelte files, you MUST run this sequence:
```bash
# 1. Clear old build and create new one
npm run build

# 2. Sync changes to the Android native project
npx cap sync android

# 3. Open Android Studio (Optional)
npx cap open android
```

## ⚠️ Troubleshooting

- **"Loading..." never disappears**: 
    1. Check if `VITE_API_URL` in `.env` is correct.
    2. Ensure your Laptop/Backend is running on the same network.
    3. Check `F12` console (or Logcat) for `Network Error`.
- **Google Login fails**: Ensure `google-services.json` is present in `android/app/` and the SHA-1 key is registered in Firebase Console.

---
🐾 *Built with Svelte 5 for a premium experience.*
