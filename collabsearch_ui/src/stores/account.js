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
      setAvatar(defaultAvatarPath)
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
    const defaultAvatarPath = 'src/assets/OIP.jpg'
    const avatar = ref(defaultAvatarPath)
    const setAvatar = (newAvatar) => {
      if (newAvatar) {
        avatar.value = newAvatar
      } else {
        avatar.value = defaultAvatarPath
      }
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
      defaultAvatarPath,
      setAvatar,
      setAll
    }
  },
  { persist: true }
)
