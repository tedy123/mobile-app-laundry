import { Token } from '$lib/utils/token'

const BASE_URL = import.meta.env.VITE_API_URL || 'http://localhost:4000/api'
const API_ROOT = BASE_URL.replace(/\/api$/, '')

export async function apiFetch<T>(path: string, options: RequestInit = {}): Promise<T> {
  let accessToken = await Token.getAccessToken()

  const headers = new Headers(options.headers)
  headers.set('Content-Type', 'application/json')
  if (accessToken) headers.set('Authorization', `Bearer ${accessToken}`)

  const url = `${BASE_URL}${path}`;
  console.log(`[API] Fetching: ${url}`);
  let res: Response;

  try {
    res = await fetch(url, { ...options, headers });
    console.log(`[API] Response: ${res.status} ${res.statusText}`);
  } catch (err: any) {
    console.error(`[API] Network Error: ${err.message}`);
    throw err;
  }

  if (res.status === 401) {
    const refreshToken = await Token.getRefreshToken()
    if (!refreshToken) {
      await Token.clear()
      throw new Error('Session expired')
    }

    const refreshRes = await fetch(`${API_ROOT}/refresh`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ refreshToken })
    })

    if (!refreshRes.ok) {
      await Token.clear()
      throw new Error('Session expired')
    }

    const data: { accessToken: string } = await refreshRes.json()
    accessToken = data.accessToken

    const email = await Token.getEmail()
    await Token.set(accessToken, refreshToken, email ?? '')

    // Retry
    headers.set('Authorization', `Bearer ${accessToken}`)
    res = await fetch(`${BASE_URL}${path}`, { ...options, headers })
  }

  if (!res.ok) {
    throw new Error(await res.text())
  }

  return res.json() as Promise<T>
}
