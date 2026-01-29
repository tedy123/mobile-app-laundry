<script lang="ts">
  let { pendapatanHariIni = 0 } = $props();
</script>

<div class="bg-gradient-to-r from-red-400 via-pink-200 to-white
            rounded-xl shadow-md px-4 py-2 w-60 mx-auto
            flex items-center justify-between">
  <div class="w-10 h-10 rounded-full bg-red-100 flex items-center justify-center shadow-inner flex-shrink-0">
    <svg class="w-5 h-5 text-red-500" fill="currentColor" viewBox="0 0 24 24">
      <path d="M12 1v22m11-11H1" stroke="currentColor" stroke-width="2" stroke-linecap="round"/>
    </svg>
  </div>

  <div class="flex flex-col items-end justify-center ml-3">
    <p class="text-xs text-gray-700 opacity-90">Pendapatan hari ini</p>
    <h1 class="text-lg font-bold text-red-600">
      IDR {pendapatanHariIni.toLocaleString('id-ID')}
    </h1>
  </div>
</div>
