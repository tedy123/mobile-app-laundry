import { apiFetch } from '$lib/api/client'

// Ini sesuai dengan JSON dari Go (lowercase/snake_case)
type TransactionSummaryResponse = {
  email: string
  tanggal: string
  qty_in: number
  qty_out: number
  amount_in: number
  amount_out: number
}

// Ini yang digunakan di UI Svelte (camelCase)
export type TransactionSummary = {
  email: string
  tanggal: string
  qtyIn: number
  qtyOut: number
  amountIn: number
  amountOut: number
}

export async function fetchTransactionByDate(
  email: string,
  tanggal: string
): Promise<TransactionSummary> {

  const params = new URLSearchParams({ email, tanggal })

  const res = await apiFetch<TransactionSummaryResponse>(
    `/transactions?${params.toString()}`,
    { method: 'GET' }
  )

  return {
    email: res.email,
    tanggal: typeof res.tanggal === 'string' ? res.tanggal.slice(0, 10) : String(res.tanggal),
    qtyIn: res.qty_in || 0,
    qtyOut: res.qty_out || 0,
    amountIn: res.amount_in || 0,
    amountOut: res.amount_out || 0
  }
}
