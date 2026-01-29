<script lang="ts">
  import { goto } from '$app/navigation';
  import { topMenus } from '$lib/config/menu';

  let { amountIn = 0, amountOut = 0, totalTransaction = 0 } = $props();

  let menuValues = $derived({
    'Penjualan': amountIn,
    'Pengeluaran': amountOut,
    'Transaksi': totalTransaction
  });

  const cardColors = ['bg-red-100', 'bg-pink-100', 'bg-red-200', 'bg-pink-200', 'bg-red-50'];
</script>

<div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-5 gap-4">
  {#each topMenus as menu, index}
    <button
      on:click={() => goto(menu.route)}
      class={`flex items-center justify-between p-3 rounded-2xl shadow-md 
              hover:shadow-xl hover:scale-105 transition-transform duration-200 ${cardColors[index % cardColors.length]}`}
    >
      <div class="flex items-center gap-3">
        <menu.icon class="w-7 h-7 text-red-500 flex-shrink-0" />
        <div class="flex flex-col">
          <span class="text-sm font-semibold text-gray-700">{menu.label}</span>
          {#if menuValues[menu.label] !== undefined}
            <span class="flex justify-start text-xs font-bold font-medium text-gray-500">
              {menu.label === 'Transaksi' ? menuValues[menu.label] : 'IDR ' + menuValues[menu.label].toLocaleString('id-ID')}
            </span>
          {/if}
        </div>
      </div>
    </button>
  {/each}
</div>
