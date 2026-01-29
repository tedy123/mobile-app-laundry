<script>
  import { goto } from '$app/navigation';
  import { slides } from '$lib/features/onboarding/slides';

  let current = 0;
  const total = slides.length;

  let startX = 0;
  let endX = 0;

  function next() {
    if (current < total - 1) {
      current += 1;
    } else {
      goto('/login');
    }
  }

  function prev() {
    if (current > 0) {
      current -= 1;
    }
  }

  function handleTouchStart(e) {
    startX = e.touches[0].clientX;
  }

  function handleTouchEnd(e) {
    endX = e.changedTouches[0].clientX;
    const diff = startX - endX;

    if (Math.abs(diff) < 50) return;
    diff > 0 ? next() : prev();
  }
</script>

<div
  class="min-h-screen flex items-center justify-center bg-white px-4"
  on:touchstart={handleTouchStart}
  on:touchend={handleTouchEnd}
>
  <div class="w-full max-w-md h-[85vh] rounded-3xl shadow-xl flex flex-col">
    <div class="flex-1 flex flex-col px-6 pt-10">
      <h1 class="text-2xl font-bold text-gray-900 text-center">
        {slides[current].title}
      </h1>

      <div class="mt-30 flex justify-center">
        <img
          src={slides[current].image}
          alt={slides[current].title}
          class="w-64 h-64 object-contain transition-all duration-300"
        />
      </div>

      <h2 class="mt-6 text-lg font-semibold text-gray-800 text-center">
        {slides[current].descriptionHeader}
      </h2>

      <p class="mt-2 text-center text-gray-600 leading-relaxed">
        {slides[current].descriptionDetail}
      </p>
    </div>

<div class="relative px-6 pb-8 flex items-center">
  <div class="absolute left-1/2 -translate-x-1/2 flex gap-2">
    {#each slides as _, i}
      <span
        class="h-2 rounded-full transition-all
          {i === current
            ? 'bg-red-500 w-2'
            : 'bg-gray-300 w-2'}"
      />
    {/each}
  </div>

  <button
    on:click={next}
    class="
      ml-auto px-5 py-2.5 rounded-full
      bg-white
      text-gray-900 font-medium
      shadow-sm
      active:scale-95
      transition-all
    "
  >
    {current === total - 1 ? 'Mulai' : 'Lanjut'}
  </button>
</div>

  </div>
</div>
