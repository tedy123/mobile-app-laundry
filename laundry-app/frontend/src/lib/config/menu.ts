import IconLayanan from '$lib/assets/Icon_layanan.png?url';
import IconPelanggan from '$lib/assets/Icon_pelanggan.png?url';
import IconKaryawan from '$lib/assets/Icon_karyawan.png?url';
import IconPengeluaran from '$lib/assets/Icon_pengeluaran.png?url';
import IconLaporan from '$lib/assets/Icon_laporan.png?url';
import IconManageOutlet from '$lib/assets/Icon_manageoutlet.png?url';
import {
  Home,
  ClipboardList,
  BarChart3,
  Settings,
  ShoppingCart,
  CreditCard,
  Receipt,
  Users,
  Briefcase,
  Store,
  CirclePlus
} from 'lucide-svelte';


export const topMenus = [
  { label: 'Penjualan', route: '/penjualan', icon: ShoppingCart },
  { label: 'Pengeluaran', route: '/pengeluaran', icon: CreditCard },
  { label: 'Transaksi', route: '/transaksi', icon: Receipt }
];

export const gridMenus = [
  { label: 'Layanan', route: '/layanan', img: IconLayanan },
  { label: 'Pelanggan', route: '/pelanggan', img: IconPelanggan },
  { label: 'Karyawan', route: '/karyawan', img: IconKaryawan },
  { label: 'Pengeluaran', route: '/pengeluaran',img: IconPengeluaran  },
  { label: 'Laporan', route: '/laporan', img: IconLaporan  },
  { label: 'Outlet', route: '/outlet', img: IconManageOutlet  }
];

export const bottomMenus = [
  { label: 'Beranda', route: '/home', icon: Home },
  { label: 'Status', route: '/status', icon: ClipboardList },
  { label: 'Pesanan', route: '/pesanan', icon: CirclePlus },
  { label: 'Laporan', route: '/laporan', icon: BarChart3 },
  { label: 'Setting', route: '/settings', icon: Settings }
];
