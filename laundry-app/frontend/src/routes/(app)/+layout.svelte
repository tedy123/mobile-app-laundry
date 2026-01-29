<script lang="ts">
  import DashboardBottomNav from '$lib/components/dashboard/DashboardBottomNav.svelte';
  import { onMount } from 'svelte';
  import { goto } from '$app/navigation';
  import { authStore } from '$lib/stores/auth.store';

  let { children } = $props();

  onMount(async () => {
    await authStore.init();
  });

  $effect(() => {
    if (!$authStore.loading && !$authStore.loggedIn) {
      goto('/login', { replaceState: true });
    }
  });
</script>

<div class="min-h-screen bg-gray-50 pb-20">
  {@render children()}
</div>

<DashboardBottomNav />
