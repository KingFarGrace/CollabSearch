<script setup>
import { ref } from 'vue'
import { storeToRefs } from 'pinia'
import { useAccountStore } from '@/stores/account'
import { convertBinaryToImageUrl } from '@/utils/file'
import { accountUpdateService } from '@/api/account'
import { useRouter } from 'vue-router'
const store = useAccountStore()
// Form
var { uid, username, email, profile, avatar } = storeToRefs(store)
var { setAvatar } = store
var logout = store.logout
var router = useRouter()
var disabled = ref(true)
var valid = ref(false)
var usernameRules = [
  (value) => {
    if (value?.length >= 1 && value?.length <= 16) return true
    return 'Username should be at least 6 characters and no more than 16.'
  }
]
var profileRules = [
  (value) => {
    if (value?.length > 0 && value?.length <= 255) return true
    return 'Profile cannot be empty or longer than 255 characters.'
  }
]
// Messager
var messager = ref(false)
var timeout = ref(2000)
var updMsg = ref('')
// Avatar upload dialog
var dialog = ref(false)
var loading = ref(false)
var rules = [
  (value) => {
    return (
      !value ||
      !value.length ||
      value[0].size < 2000000 ||
      'Avatar size should be less than 2 MB!'
    )
  }
]
async function updateInfo() {
  if (!valid.value) return
  // Change update form status.
  // When !disabled after changed, do nothing and return.
  disabled.value = !disabled.value
  if (!disabled.value) {
    return
  }
  try {
    const data = await accountUpdateService(
      uid.value,
      '',
      username.value,
      profile.value,
      ''
    )
    updMsg.value = data.Msg
    messager.value = true
  } catch (error) {
    updMsg.value = error.Msg
    messager.value = true
  }
}
function logoutToAuthPage() {
  logout()
  router.push('/login')
}
function readAvatarFile(file) {
  return new Promise((resolve, reject) => {
    const reader = new FileReader()

    reader.onload = (event) => {
      resolve(event.target.result)
    }

    reader.onerror = (error) => {
      reject(error)
    }

    reader.readAsArrayBuffer(file)
  })
}
// TODO: Optimse - store image in binary type, display it by url.
// Now: store image and display it all by url.
async function updateAvatar(files) {
  const content = await readAvatarFile(files[0])
  setAvatar(convertBinaryToImageUrl(content))
  try {
    loading.value = true
    const data = await accountUpdateService(uid.value, '', '', '', avatar.value)
    loading.value = false
    updMsg.value = data.Msg
    messager.value = true
  } catch (error) {
    loading.value = false
    updMsg.value = error.Msg
    messager.value = true
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <v-card width="500" height="auto"
    ><v-sheet width="300" class="mx-auto">
      <v-form
        fast-fail
        @submit.prevent
        :model-value="valid"
        @update:modelValue="valid = $event"
      >
        <v-avatar size="90" @click="dialog = true">
          <v-img alt="Avatar" :src="avatar" elevation="3">
            <template v-slot:error>
              <v-img
                class="mx-auto"
                height="300"
                max-width="500"
                src="/src/assets/OIP.jpg"
              ></v-img>
            </template>
          </v-img>
        </v-avatar>
        <br />

        <v-text-field
          v-model="uid"
          label="User ID"
          disabled="true"
        ></v-text-field>
        <v-text-field
          v-model="email"
          label="Email"
          disabled="true"
        ></v-text-field>
        <v-text-field
          v-model="username"
          label="Username"
          :disabled="disabled"
          :rules="usernameRules"
        ></v-text-field>
        <v-text-field
          v-model="profile"
          label="Profile"
          :disabled="disabled"
          :rules="profileRules"
        ></v-text-field>
        <v-btn class="mt-2 mx-auto w-50" @click="updateInfo" v-if="disabled">
          Update
        </v-btn>
        <v-btn class="mt-2 mx-auto w-50" @click="updateInfo" v-if="!disabled">
          Finish
        </v-btn>
      </v-form>
      <br />
      <v-divider></v-divider>
      <v-btn color="error" class="mt-2 mx-auto w-50" @click="logoutToAuthPage"
        >Log out</v-btn
      > </v-sheet
    ><br
  /></v-card>
  <v-dialog v-model="dialog" width="1000" height="500">
    <v-card>
      <span class="mx-auto w-75"
        ><v-file-input
          :rules="rules"
          :loading="loading"
          show-size
          accept="image/jpeg"
          placeholder="Pick an avatar"
          prepend-icon="mdi-camera"
          label="Avatar"
          @update:modelValue="updateAvatar"
        ></v-file-input
      ></span>

      <v-card-actions>
        <v-btn color="primary" block @click="dialog = false"
          >Save and Close</v-btn
        >
      </v-card-actions>
    </v-card>
  </v-dialog>
  <v-snackbar v-model="messager" :timeout="timeout">
    {{ updMsg }}
    <template v-slot:actions>
      <v-btn color="blue" variant="text" @click="messager = false">
        Close
      </v-btn>
    </template>
  </v-snackbar>
</template>
