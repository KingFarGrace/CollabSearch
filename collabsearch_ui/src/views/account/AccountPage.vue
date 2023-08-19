<script setup>
import { ref } from 'vue'
import { storeToRefs } from 'pinia'
import { useAccountStore } from '@/stores/account'
import { useRouter } from 'vue-router'
const store = useAccountStore()
var { uid, username, email, profile, avatar } = storeToRefs(store)
var logout = store.logout
var router = useRouter()
var disabled = ref(true)
var dialog = ref(false)
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
function updateInfo() {
  // TODO: PUT user info
  disabled.value = !disabled.value
}
function updateAvatar() {
  dialog.value = false
}
function logoutToAuthPage() {
  logout()
  router.push('/login')
}
</script>

<template>
  <v-card width="500" height="auto"
    ><v-sheet width="300" class="mx-auto">
      <v-form>
        <v-avatar size="90" @click="dialog = true">
          <v-img alt="Avatar" :src="avatar"></v-img>
        </v-avatar>
        <br />
        <v-dialog v-model="dialog" width="1000" height="500">
          <v-card>
            <span class="mx-auto w-75"
              ><v-file-input
                :rules="rules"
                show-size
                accept="image/png, image/jpeg, image/bmp"
                placeholder="Pick an avatar"
                prepend-icon="mdi-camera"
                label="Avatar"
              ></v-file-input
            ></span>

            <v-card-actions>
              <v-btn color="primary" block @click="updateAvatar"
                >Save and Close</v-btn
              >
            </v-card-actions>
          </v-card>
        </v-dialog>

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
        ></v-text-field>
        <v-text-field
          v-model="profile"
          label="Profile"
          :disabled="disabled"
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
</template>
