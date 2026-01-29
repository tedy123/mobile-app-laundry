# Laundry Management System 🧺✨

A professional, multi-platform solution designed for modern laundry businesses. This system integrates a high-performance Go backend with a cutting-edge Svelte 5 mobile interface.

## 📁 System Architecture

The project is architected for scalability and native mobile performance:

- **[Backend](./backend)**: 
    - **Tech**: Go 1.23, PostgreSQL, JWT.
    - **Auth**: Direct Google API verification for server-side security.
    - **Goal**: Provides a lightweight, fast REST API.
  
- **[Frontend](./frontend)**: 
    - **Tech**: SvelteKit (Svelte 5), Capacitor, Tailwind CSS.
    - **Auth**: Native Google login integration.
    - **Goal**: A premium, "App-like" experience on Android and Web.

## 🚀 Getting Started

### 1. Backend Preparation
```bash
cd backend
go mod tidy
go run cmd/api/main.go
```
*Port default: 4000. Ensure PostgreSQL is active.*

### 2. Frontend Mobile Prep
```bash
cd frontend
npm install
npm run build
npx cap sync android
```
*Run the app from Android Studio or via `npx cap run android`.*

## 🔑 Key Workflows

### Authentication
User identity is verified natively on the device. The backend receives a Google ID Token, verifies it server-side, and issues a long-lived JWT.

### Clean Session Logout
Logging out clears:
- Local Storage (Preferences)
- JWT Tokens
- Svelte Stores (Auth & Transactions)
- Local Cache

## 🛠️ Tech Highlights
- **Svelte 5 Runes**: Utilizes `$state`, `$derived`, and `$effect` for ultra-reactive UI.
- **Aggregated Queries**: Backend calculates revenue and quantity sums on-the-fly for real-time dashboard accuracy.
- **Android Networking**: Specifically optimized for Android Emulator connectivity via `10.0.2.2`.

---
🐾 *Ayocuci Laundry - Streamlining Laundry Management.*
