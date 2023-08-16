<script setup>
import { ref } from 'vue'
import { storeToRefs } from 'pinia'
import { useAccountStore } from '@/stores/account'
const store = useAccountStore()
var { uid, username, email, profile, avatar } = storeToRefs(store)
var logout = store.logout
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
</script>

<template>
  <v-sheet width="300" class="mx-auto">
    <v-form>
      <div>
        <v-avatar size="90" @click="dialog = true">
          <v-img alt="Avatar" :src="avatar"></v-img>
        </v-avatar>
        <v-dialog v-model="dialog" width="480px">
          <v-card>
            <v-file-input
              :rules="rules"
              show-size
              accept="image/png, image/jpeg, image/bmp"
              placeholder="Pick an avatar"
              prepend-icon="mdi-camera"
              label="Avatar"
            ></v-file-input>

            <v-card-actions>
              <v-btn color="primary" block @click="updateAvatar"
                >Save and Close</v-btn
              >
            </v-card-actions>
          </v-card>
        </v-dialog>
      </div>

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
      <v-btn @click="updateInfo"> Update account </v-btn>
    </v-form>
    <br />
    <v-btn color="error" @click="logout">Log out</v-btn>
  </v-sheet>
</template>
