<script lang="ts">
  import { Bell, LogOut, X } from 'lucide-svelte';
  import iconOutlet from '$lib/assets/Icon_manageoutlet.png';
  import { goto } from '$app/navigation';
  import { authStore } from '$lib/stores/auth.store';

  let showModal = $state(false);

  async function handleLogout() {
    await authStore.logout();
    goto('/login');
  }
</script>

<div class="w-full flex items-center justify-between px-1 py-2">
  <div class="flex items-center gap-3">
    <button 
      on:click={() => showModal = true}
      class="focus:outline-none active:scale-95 transition-transform"
    >
      <img 
        src={iconOutlet} 
        alt="Logo" 
        class="w-10 h-10 rounded-full object-contain border-2 border-white shadow-sm"
      />
    </button>

    <div class="flex flex-col leading-tight">
      <span class="text-sm font-bold text-white">Ayocuci Laundry</span>
      <span class="text-xs text-white opacity-90">Jakarta, Indonesia</span>
    </div>
  </div>

  <button class="flex items-center justify-center w-10 h-10 rounded-full bg-red-100 hover:bg-red-200 transition">
    <Bell class="w-5 h-5 text-red-600"/>
  </button>
</div>

{#if showModal}
  <div class="fixed inset-0 z-50 flex items-center justify-center p-6">
    <div 
      class="absolute inset-0 bg-black/40 backdrop-blur-sm"
      on:click={() => showModal = false}
    ></div>

    <div class="relative w-full max-w-xs bg-white rounded-3xl p-6 shadow-2xl scale-in-center">
      <button 
        on:click={() => showModal = false}
        class="absolute right-4 top-4 text-gray-400 hover:text-gray-600"
      >
        <X size={20} />
      </button>

      <div class="flex flex-col items-center text-center space-y-4">
        <div class="w-20 h-20 bg-red-50 rounded-full flex items-center justify-center">
          <img src={iconOutlet} alt="Logo" class="w-14 h-14 object-contain" />
        </div>
        
        <div>
          <h3 class="text-xl font-bold text-gray-800">Ayocuci Laundry</h3>
          <p class="text-sm text-gray-500">Apakah Anda ingin keluar dari akun ini?</p>
        </div>

        <button
          on:click={handleLogout}
          class="w-full py-3 bg-red-500 text-white font-bold rounded-2xl flex items-center justify-center gap-2 hover:bg-red-600 transition active:scale-95 shadow-lg shadow-red-200"
        >
          <LogOut size={18} />
          Keluar
        </button>

        <button 
          on:click={() => showModal = false}
          class="text-sm font-medium text-gray-400 hover:text-gray-600"
        >
          Batal
        </button>
      </div>
    </div>
  </div>
{/if}

<style>
  .scale-in-center {
    animation: scale-in-center 0.3s cubic-bezier(0.250, 0.460, 0.450, 0.940) both;
  }

  @keyframes scale-in-center {
    0% {
      transform: scale(0);
      opacity: 1;
    }
    100% {
      transform: scale(1);
      opacity: 1;
    }
  }
</style>
