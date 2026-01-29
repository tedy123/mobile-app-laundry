<script lang="ts">
  import { onMount } from 'svelte'
  import { authStore } from '$lib/stores/auth.store'
  import { transactionStore } from '$lib/stores/transaction.store'
  import DashboardHeader from '$lib/components/dashboard/DashboardHeader.svelte'
  import DashboardTopMenu from '$lib/components/dashboard/DashboardTopMenu.svelte'
  import DashboardGridMenu from '$lib/components/dashboard/DashboardGridMenu.svelte'
  import DashboardTopNotif from '$lib/components/dashboard/DashboardTopNotif.svelte'
  import DashboardSlideShow from '$lib/components/dashboard/DashboardSlideShow.svelte'
  import { Token } from '$lib/utils/token';
  import { goto } from '$app/navigation';

  onMount(async () => {
    await authStore.init();
  });

  function getLocalDateString() {
    const now = new Date();
    const year = now.getFullYear();
    const month = String(now.getMonth() + 1).padStart(2, '0');
    const day = String(now.getDate()).padStart(2, '0');
    return `${year}-${month}-${day}`;
  }

  $effect(() => {
    if (!$authStore.loading && $authStore.loggedIn && $authStore.user) {
      const today = getLocalDateString();
      console.log(`[Home] Loading transactions for ${$authStore.user.email} on ${today}`);
      transactionStore.load($authStore.user.email, today);
    }
  });
</script>

<div
  class="min-h-screen px-4 pt-5 space-y-2
         bg-gradient-to-b from-red-500 via-pink-200 to-white"
>
  <DashboardTopNotif />

  {#if $transactionStore.loading}
    <p class="text-center text-white">Loading transaksi...</p>
  {/if}

  {#if $transactionStore.error}
    <p class="text-center text-red-600 font-bold bg-white/80 p-2 rounded">
      {$transactionStore.error}
    </p>
  {/if}

  {#if $transactionStore.data}
    <DashboardHeader
      pendapatanHariIni={
        ($transactionStore.data.amountIn ?? 0) -
        ($transactionStore.data.amountOut ?? 0)
      }
    />

    <DashboardTopMenu
      amountIn={$transactionStore.data.amountIn ?? 0}
      amountOut={$transactionStore.data.amountOut ?? 0}
      totalTransaction={
        ($transactionStore.data.qtyIn ?? 0) +
        ($transactionStore.data.qtyOut ?? 0)
      }
    />
  {/if}

  <DashboardGridMenu />
  <DashboardSlideShow />
</div>
