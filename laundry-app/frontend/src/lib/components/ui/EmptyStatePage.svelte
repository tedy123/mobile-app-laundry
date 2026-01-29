<script lang="ts">
  import { goto } from '$app/navigation';
  import comingSoonImg from '$lib/assets/ComingSoon.png';

  let { title = "Halaman" } = $props();

  function goBack() {
    goto('/home');
  }
</script>

<div class="min-h-screen bg-gray-50 flex flex-col">
  <!-- Header -->
  <div class="w-full flex items-center gap-3 px-4 py-3 bg-red-400 shadow-md z-10">
    <button
      on:click={goBack}
      class="p-2 rounded-full hover:bg-white/20 transition active:scale-90"
    >
      <svg class="w-8 h-8 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
              d="M15 19l-7-7 7-7"/>
      </svg>
    </button>
    <h1 class="text-lg font-semibold text-[22px] text-white">{title}</h1>
  </div>

  <!-- Content -->
  <div class="flex-1 flex flex-col items-center justify-center p-8 space-y-6 text-center">
    <div class="w-64 h-64 animate-bounce-slow">
      <img 
        src={comingSoonImg} 
        alt="Coming Soon" 
        class="w-full h-full object-contain drop-shadow-2xl"
      />
    </div>

    <div class="space-y-2">
      <h2 class="text-2xl font-bold text-gray-800">Ups, Belum Tersedia!</h2>
      <p class="text-gray-500 max-w-xs mx-auto">
        Halaman <span class="font-bold text-red-400">{title}</span> saat ini sedang dalam pengembangan. Tunggu kabar selanjutnya ya!
      </p>
    </div>

    <button
      on:click={goBack}
      class="px-8 py-3 bg-red-400 text-white font-bold rounded-2xl shadow-lg shadow-red-100 hover:bg-red-500 transition active:scale-95"
    >
      Kembali ke Beranda
    </button>
  </div>
</div>

<style>
  .animate-bounce-slow {
    animation: bounce-slow 3s ease-in-out infinite;
  }

  @keyframes bounce-slow {
    0%, 100% {
      transform: translateY(0);
    }
    50% {
      transform: translateY(-20px);
    }
  }
</style>
