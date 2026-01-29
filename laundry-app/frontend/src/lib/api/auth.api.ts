import { Token } from '$lib/utils/token'

const BASE_URL = import.meta.env.VITE_API_URL || 'https://localhost:4000'

export async function login(email: string, password: string) {
  const res = await fetch(`${BASE_URL}/login`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ email, password })
  })

  if (!res.ok) throw new Error(await res.text())
  const data = await res.json()
  await Token.set(data.accessToken, data.refreshToken, data.email)
  return data
}

export async function googleLogin(idToken: string) {
  const res = await fetch(`${BASE_URL}/google-login`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ idToken })
  })

  if (!res.ok) throw new Error(await res.text())
  const data = await res.json()
  await Token.set(data.accessToken, data.refreshToken, data.email)
  return data
}

export async function logout() {
  await Token.clear()
}
