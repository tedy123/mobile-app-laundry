<script lang="ts">
  import { goto } from '$app/navigation';
  import { page } from '$app/stores';
  import { bottomMenus } from '$lib/config/menu';
</script>

<nav
  class="fixed bottom-0 left-0 right-0
         bg-white border-t shadow
         flex justify-around py-3"
>
  {#each bottomMenus as menu}
    <button
      on:click={() => goto(menu.route)}
      class="flex flex-col items-center text-xs"
    >
      <menu.icon
        class="w-6 h-6
        {$page.url.pathname.startsWith(menu.route)
          ? 'text-red-400'
          : 'text-gray-400'}"
      />
      <span
        class="mt-1
        {$page.url.pathname.startsWith(menu.route)
          ? 'text-red-400 font-semibold'
          : 'text-gray-400'}"
      >
        {menu.label}
      </span>
    </button>
  {/each}
</nav>
