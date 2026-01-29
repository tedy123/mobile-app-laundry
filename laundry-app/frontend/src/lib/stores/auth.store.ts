import { writable } from 'svelte/store'
import { login as loginAPI, logout as logoutAPI } from '$lib/api/auth.api'
import { Token } from '$lib/utils/token'
import { transactionStore } from '$lib/stores/transaction.store'

type AuthState = {
  user: { email: string } | null
  loggedIn: boolean
  loading: boolean
}

function createAuthStore() {
  const { subscribe, set, update } = writable<AuthState>({
    user: null,
    loggedIn: false,
    loading: true
  })

  return {
    subscribe,

    async login(email: string, password: string) {
      update(s => ({ ...s, loading: true }))
      const data = await loginAPI(email, password)
      set({ user: { email: data.email }, loggedIn: true, loading: false })
    },

    async logout() {
      await logoutAPI()
      transactionStore.reset()
      set({ user: null, loggedIn: false, loading: false })
    },

    async init() {
      const token = await Token.getAccessToken()
      const email = await Token.getEmail()
      set({
        user: token && email ? { email } : null,
        loggedIn: !!token,
        loading: false
      })
    }
  }
}

export const authStore = createAuthStore()
