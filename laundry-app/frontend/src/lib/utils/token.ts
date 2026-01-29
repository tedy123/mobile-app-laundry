import { Preferences } from '@capacitor/preferences'

export class Token {
  static async getAccessToken(): Promise<string | null> {
    const { value } = await Preferences.get({ key: 'accessToken' })
    return value
  }

  static async getRefreshToken(): Promise<string | null> {
    const { value } = await Preferences.get({ key: 'refreshToken' })
    return value
  }

  static async getEmail(): Promise<string | null> {
    const { value } = await Preferences.get({ key: 'email' })
    return value
  }

  static async set(
    accessToken: string,
    refreshToken: string,
    email: string | null
  ) {
    await Preferences.set({ key: 'accessToken', value: accessToken })
    await Preferences.set({ key: 'refreshToken', value: refreshToken })

    if (email) {
      await Preferences.set({ key: 'email', value: email })
    }
  }

  static async clear() {
    await Preferences.remove({ key: 'accessToken' })
    await Preferences.remove({ key: 'refreshToken' })
    await Preferences.remove({ key: 'email' })
  }
}
