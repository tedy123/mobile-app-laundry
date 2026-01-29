<script lang="ts">
  import { goto } from '$app/navigation';

  function goBack() {
    goto('/home'); // kembali ke halaman home
  }

  // Contoh data detail
  const detail = {
    image: '/images/Promo_tWO.png', // ganti dengan path foto
    title: 'Satu Aplikasi Untuk Semua Jenis Laundry',
    description: 'Ada pemilik laundry kiloan yang setiap hari sibuk mencatat berat cucian, mengecek status cucian, mengecek status pesanan dan menghitung omset diakhir hari, dan seteusnya......'
  };
</script>

<!-- Header -->
<div class="w-full flex items-center gap-3 px-4 py-3 bg-red-400 shadow-md">
  <button
    on:click={goBack}
    class="p-2 rounded-full hover:bg-gray-100 transition"
  >
    <svg class="w-8 h-8 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
            d="M15 19l-7-7 7-7"/>
    </svg>
  </button>

  <h1 class="text-lg font-semibold text-[22px] text-white">Detail</h1>
</div>

<!-- Konten -->
<div class="">
  <!-- Foto full-width -->
  <div class="w-full  overflow-hidden shadow-md">
    <img
      src={detail.image}
      alt={detail.title}
      class="w-full h-auto object-cover"
    />
  </div>

  <!-- Judul -->
  <h2 class="text-xl mt-8 px-4 mb-2 font-bold text-gray-800">{detail.title}</h2>

  <!-- Deskripsi -->
  <p class="text-gray-700 px-4 text-sm leading-relaxed">{detail.description}</p>
</div>



