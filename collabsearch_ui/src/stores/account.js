import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useAccountStore = defineStore(
  'account-store',
  () => {
    const token = ref('')
    const setToken = (newToken) => {
      token.value = newToken
    }
    const isLogin = ref(false)
    const login = () => {
      isLogin.value = true
    }
    const logout = () => {
      isLogin.value = false
    }
    const uid = ref(0)
    const setUid = (newUid) => {
      uid.value = newUid
    }
    const email = ref('')
    const setEmail = (newEmail) => {
      email.value = newEmail
    }
    const username = ref('')
    const setUsername = (newName) => {
      username.value = newName
    }
    const profile = ref('')
    const setProfile = (newProfile) => {
      profile.value = newProfile
    }
    // Default avatar path
    const avatar = ref('src/assets/OIP.jpg')
    const setAvatar = (newAvatar) => {
      avatar.value = newAvatar
    }
    const getUserInfo = async () => {}

    return {
      token,
      setToken,
      isLogin,
      login,
      logout,
      uid,
      setUid,
      email,
      setEmail,
      username,
      setUsername,
      profile,
      setProfile,
      avatar,
      setAvatar,
      getUserInfo
    }
  },
  { persist: true }
)
