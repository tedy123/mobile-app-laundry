<script lang="ts">
  import { goto } from '$app/navigation';
  import { gridMenus } from '$lib/config/menu';
</script>

<div class="p-1  mb-2">
  <div class="bg-white rounded-2xl w-full shadow-lg p-4">
    <div class="grid grid-cols-3 gap-4">
      {#each gridMenus as menu}
        <button
          on:click={() => goto(menu.route)}
          class="flex flex-col items-center gap-2 text-gray-700 transition
                 rounded-lg bg-red-50 hover:bg-red-100 hover:shadow-md p-2 focus:outline-none focus:ring-2 focus:ring-red-400"
        >
          <div class="w-14 h-14 rounded-2xl bg-white flex items-center justify-center transition-transform transform hover:scale-105">
            {#if menu.img}
              <img src={menu.img} alt={menu.label} class="w-13 h-13 object-contain" />
            {:else if menu.component}
              <svelte:component this={menu.component} class="w-12 h-12 text-red-400" />
            {/if}
          </div>
          <span class="text-xs font-medium text-center">{menu.label}</span>
        </button>
      {/each}
    </div>
  </div>
</div>



