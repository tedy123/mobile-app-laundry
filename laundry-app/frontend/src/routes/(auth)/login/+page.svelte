<script lang="ts">
  import { Mail, Lock, Eye, EyeOff } from 'lucide-svelte';
  import { goto } from '$app/navigation';
  import { browser } from '$app/environment';
  import Logo2 from '$lib/assets/Logo_2.png';
  import AssetOne from '$lib/assets/Asset_One.png';
  import { Token } from '$lib/utils/token';
  import { authStore } from '$lib/stores/auth.store'
  import { onMount } from 'svelte';
  import { auth as firebaseAuth } from '$lib/firebase/client';
  import { GoogleAuthProvider, signInWithCredential, signInWithPopup } from 'firebase/auth';
  import { GoogleAuth } from '@codetrix-studio/capacitor-google-auth';
  import { googleLogin } from '$lib/api/auth.api';
  import { Capacitor } from '@capacitor/core';

  let email = '';
  let password = '';
  let remember = false;
  let showPassword = false;
  let loading = false;
  let error: string | null = null;

  const BASE_URL = import.meta.env.VITE_API_URL || 'http://localhost:4000/api';

  async function handleLogin() {
    error = '';
    loading = true;

    try {
      const res = await fetch(`${BASE_URL}/login`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ email, password })
      });

      if (!res.ok) {
        error = await res.text();
        return;
      }

      const data = await res.json();
      await Token.set(data.accessToken, data.refreshToken, data.email);
      await authStore.init();
      goto('/home', { replaceState: true });

    } catch (err) {
      error = 'Gagal koneksi ke server';
      console.error(err);
    } finally {
      loading = false;
    }
  }

  onMount(async () => {
    if (Capacitor.isNativePlatform()) {
      try {
        await GoogleAuth.initialize();
        console.log('GoogleAuth initialized');
      } catch (e) {
        console.error('GoogleAuth init failed:', e);
      }
    }
  });

  async function handleGoogleLogin() {
    try {
      loading = true;
      error = '';

      let idToken: string;

      if (Capacitor.isNativePlatform()) {
        const user = await GoogleAuth.signIn();
        idToken = user.authentication.idToken;
      } else {
        const provider = new GoogleAuthProvider();
        const result = await signInWithPopup(firebaseAuth, provider);
        idToken = await result.user.getIdToken();
      }
      
      const credential = GoogleAuthProvider.credential(idToken);
      await signInWithCredential(firebaseAuth, credential);
      await googleLogin(idToken);
      await authStore.init();
      
      goto('/home', { replaceState: true });
    } catch (err: any) {
      if (err.message?.includes('cancel') || err.message?.includes('12501')) {
        console.log('User cancelled login');
      } else {
        error = 'Login Error: ' + (err.message || JSON.stringify(err));
        console.error('Google Login Error:', err);
      }
    } finally {
      loading = false;
    }
  }
</script>

<div class="min-h-screen flex flex-col justify-center px-6 sm:px-8 bg-gray-50">
  <div class="flex justify-center mb-4">
    <img src={Logo2} alt="Logo" class="w-28 h-14" />
  </div>

  <div class="text-center mb-4">
    <h1 class="text-2xl font-bold text-red-400">Welcome</h1>
    <p class="mt-1 text-gray-500 text-sm">Senang menyambut Anda kembali. Mari lanjutkan aktivitas Anda.</p>
  </div>

  <div class="flex justify-center mb-4">
    <img src={AssetOne} alt="Asset" class="w-40 h-40" />
  </div>

  {#if error}
    <div class="text-red-500 text-sm mb-3 text-center">{error}</div>
  {/if}

  <div class="mb-4 relative">
    <label class="text-sm text-gray-600 block mb-1">Email</label>
    <div class="relative">
      <Mail class="absolute left-3 top-1/2 -translate-y-1/2 w-5 h-5 text-gray-400 pointer-events-none z-10"/>
      <input
        type="email"
        bind:value={email}
        placeholder="Email"
        class="w-full pl-10 pr-4 py-3 rounded-lg shadow-md focus:ring-2 focus:ring-red-400 focus:outline-none relative z-0 bg-white"
      />
    </div>
  </div>

  <div class="mb-3 relative">
    <label class="text-sm text-gray-600 block mb-1">Password</label>
    <div class="relative">
      <Lock class="absolute left-3 top-1/2 -translate-y-1/2 w-5 h-5 text-gray-400 pointer-events-none z-10" />
      <button
        type="button"
        class="absolute right-3 top-1/2 -translate-y-1/2 z-20 text-gray-400"
        on:click={() => (showPassword = !showPassword)}
      >
        {#if showPassword}
          <EyeOff class="w-5 h-5"/>
        {:else}
          <Eye class="w-5 h-5"/>
        {/if}
      </button>
      <input
        type={showPassword ? 'text' : 'password'}
        bind:value={password}
        placeholder="Password"
        class="w-full pl-10 pr-10 py-3 rounded-lg shadow-md focus:ring-2 focus:ring-red-400 focus:outline-none relative z-0 bg-white"
      />
    </div>
  </div>

  <div class="flex justify-between items-center text-sm mb-6">
    <label class="flex items-center gap-2 text-gray-600">
      <input type="checkbox" bind:checked={remember} class="accent-red-400"/>
      Ingat saya
    </label>
    <button class="text-red-400 hover:underline">Lupa password?</button>
  </div>

  <button
    on:click={handleLogin}
    disabled={loading}
    class="w-full py-3 bg-red-400 text-white font-semibold rounded-lg hover:bg-red-500 transition disabled:opacity-50"
  >
    {loading ? 'Masuk...' : 'Login'}
  </button>

  <div class="flex items-center my-4">
    <div class="flex-1 h-px bg-gray-300"></div>
    <span class="px-3 text-sm text-gray-400">atau</span>
    <div class="flex-1 h-px bg-gray-300"></div>
  </div>

  <button
    on:click={handleGoogleLogin}
    class="w-full py-3 border rounded-lg flex items-center justify-center gap-3 text-gray-700 hover:bg-gray-100 transition"
  >
    <svg class="w-5 h-5" viewBox="0 0 533.5 544.3">
      <path fill="#4285F4" d="M533.5 278.4c0-17.4-1.6-34.1-4.7-50.4H272v95.3h146.9c-6.3 34-25.1 62.8-53.5 82v68h86.4c50.5-46.5 81.7-115.1 81.7-195z"/>
      <path fill="#34A853" d="M272 544.3c72.6 0 133.5-24.1 178-65.4l-86.4-68c-24 16.1-54.8 25.6-91.6 25.6-70.5 0-130.2-47.6-151.6-111.3h-89.7v69.9C75.1 482.3 168.7 544.3 272 544.3z"/>
      <path fill="#FBBC05" d="M120.4 325.2c-10.4-30.9-10.4-64.2 0-95.1v-69.9h-89.7c-39 77.9-39 157 0 234.9l89.7-69.9z"/>
      <path fill="#EA4335" d="M272 107.7c39.5-.6 77.5 14.2 106.5 41.3l79.4-79.4C412.5 24.4 343.4-1.3 272 0 168.7 0 75.1 62 30.7 160.2l89.7 69.9C141.8 155.3 201.5 107.7 272 107.7z"/>
    </svg>
    Masuk dengan Google
  </button>

  <div class="flex items-center justify-center my-4 text-sm text-gray-400">
    Belum punya akun? 
    <button class="text-red-400 hover:underline ml-1">Daftar Sekarang?</button>
  </div>
</div>