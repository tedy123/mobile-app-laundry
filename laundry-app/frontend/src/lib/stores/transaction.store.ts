import { writable } from 'svelte/store'
import type { TransactionSummary } from '$lib/api/transaction.api'
import { fetchTransactionByDate } from '$lib/api/transaction.api'

type TransactionState = {
  data: TransactionSummary | null
  loading: boolean
  error: string | null
  lastKey: string | null
}

function createTransactionStore() {
  const { subscribe, set, update } = writable<TransactionState>({
    data: null,
    loading: false,
    error: null,
    lastKey: null
  })

  return {
    subscribe,

    async load(email: string, tanggal: string) {
      const key = `${email}_${tanggal}`
      update(s => {
        if (s.loading || s.lastKey === key) return s
        return { ...s, loading: true, error: null, lastKey: key }
      })

      try {
        const data = await fetchTransactionByDate(email, tanggal)
        set({ data, loading: false, error: null, lastKey: key })
      } catch (err: any) {
        set({ data: null, loading: false, error: err.message ?? 'Failed to load transaction', lastKey: null })
      }
    },

    reset() {
      set({ data: null, loading: false, error: null, lastKey: null })
    }
  }
}

export const transactionStore = createTransactionStore()
