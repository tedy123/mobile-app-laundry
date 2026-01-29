<script lang="ts">
  import { onMount } from 'svelte';
  import { goto } from '$app/navigation';
  import { authStore } from '$lib/stores/auth.store';

  onMount(async () => {
    // Jalankan init untuk cek token lama
    await authStore.init();
    
    setTimeout(() => {
      if ($authStore.loggedIn) {
        goto('/home', { replaceState: true });
      } else {
        goto('/onboarding', { replaceState: true });
      }
    }, 2000);
  });
</script>

<div class="relative w-full h-screen overflow-hidden shadow-2xl">
  <img
    src="/images/Logo.png"
    alt="Laundry"
    class="absolute inset-0 w-full h-full object-cover"
    decoding="async"
    loading="eager"
    draggable="false"
  />
  <div class="absolute inset-0 bg-gradient-to-t from-black/70 via-black/40 to-transparent"></div>
</div>
