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
      setToken('')
      setUid(0)
      setEmail('')
      setUsername('')
      setProfile('')
      setAvatar('')
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
    const avatar = ref('')
    const setAvatar = (content) => {
      avatar.value = content
    }
    const setAll = ({
      uid: newUid,
      email: newEmail,
      username: newUsername,
      profile: newProfile,
      avatar: newAvatar
    }) => {
      setUid(newUid)
      setEmail(newEmail)
      setUsername(newUsername)
      setProfile(newProfile)
      setAvatar(newAvatar)
      login()
    }

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
      setAll
    }
  },
  { persist: true }
)
