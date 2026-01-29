<script>
  import { page } from '$app/stores';
</script>

<div class="min-h-screen flex flex-col items-center justify-center p-4 text-center">
  <h1 class="text-4xl font-bold text-red-500">{$page.status}</h1>
  <p class="mt-2 text-xl text-gray-600">{$page.error?.message || 'Page not found'}</p>
  <a href="/" class="mt-6 px-4 py-2 bg-red-400 text-white rounded-lg">Kembali ke Beranda</a>
</div>
