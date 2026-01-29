import { writable } from "svelte/store";

export const session = writable<{ token: string | null; role: string | null }>({ token: null, role: null });

export function setSession(token: string | null, role: string | null = null) {
  session.set({ token, role });
  if (token) localStorage.setItem("token", token);
  else localStorage.removeItem("token");
  if (role) localStorage.setItem("role", role);
  else localStorage.removeItem("role");
}

export function loadSession() {
  const token = localStorage.getItem("token");
  const role = localStorage.getItem("role");
  if (token) session.set({ token, role });
}
