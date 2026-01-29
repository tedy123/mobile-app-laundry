<script lang="ts">
  import { slides } from '$lib/config/slides';
  import { goto } from '$app/navigation';
  import { onMount } from 'svelte';

  let current = 0;
  const total = slides.length;

  let startX = 0;
  let endX = 0;

  function next() {
    current = Math.min(current + 1, total - 1);
  }

  function prev() {
    current = Math.max(current - 1, 0);
  }

  function handleTouchStart(e: TouchEvent) {
    startX = e.touches[0].clientX;
  }

  function handleTouchMove(e: TouchEvent) {
  }

  function handleTouchEnd(e: TouchEvent) {
    endX = e.changedTouches[0].clientX;
    const diff = startX - endX;
    if (diff > 50) next();
    if (diff < -50) prev();
  }

  function handleClick(slide) {
    goto(slide.route);
  }

  let interval: number;
  onMount(() => {
    interval = setInterval(() => {
      current = current < total - 1 ? current + 1 : 0;
    }, 5000);

    return () => clearInterval(interval);
  });
</script>

<div class="pl-1 mb-1"><span class="text-xs font-medium text-center">Promo</span></div>
<div 
  class="w-full overflow-hidden relative rounded-xl"
  on:touchstart={handleTouchStart}
  on:touchmove={handleTouchMove}
  on:touchend={handleTouchEnd}
>
  <div
    class="flex transition-transform duration-500 ease-in-out"
    style="transform: translateX(-{current * 70}%);"
  >
    {#each slides as slide}
      <div class="flex-shrink-0 w-[90%] pr-2 cursor-pointer">
        <div class="rounded-xl overflow-hidden shadow-lg hover:shadow-xl transition-shadow">
          <img
            src={slide.image}
            alt="slide"
            class="w-full h-[150px] object-cover"
            on:click={() => handleClick(slide)}
          />
        </div>
      </div>
    {/each}
  </div>

  <div class="flex justify-center mt-2 gap-2">
    {#each slides as _, i}
      <span
        class="w-2 h-2 rounded-full transition-colors"
        class:bg-red-500={i === current}
        class:bg-gray-300={i !== current}
      ></span>
    {/each}
  </div>
</div>
