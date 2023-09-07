<script setup>
import { ref } from 'vue'
import { accountRegisterService } from '@/api/account'
import { success } from '@/utils/resp'
import { useRouter } from 'vue-router'
const router = useRouter()
var email = ref('')
var username = ref('')
var password = ref('')
var confirm = ref('')
var valid = ref(false)
var messager = ref(false)
var timeout = ref(2000)
var regMsg = ref('')
var emailRules = [
  (value) => {
    const reg = /^[a-zA-Z0-9._-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,6}$/
    if (reg.test(value)) return true
    return 'Invalid email address.'
  }
]

var usernameRules = [
  (value) => {
    if (value?.length >= 1 && value?.length <= 16) return true
    return 'Username should be at least 6 characters and no more than 16.'
  }
]

var passwordRules = [
  (value) => {
    const reg = /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[\W_]).{6,16}$/
    if (reg.test(value)) return true
    return 'Password should contain at least One uppercase&lowercase letter, One number, One special english charactor and length between 6 and 16.'
  }
]

var confirmRules = [
  (value) => {
    if (value === password.value) return true
    return 'Confirm password does not match the password you input before.'
  }
]

async function sumbitRegisterForm() {
  if (!valid.value) return
  try {
    const data = await accountRegisterService(
      email.value,
      username.value,
      password.value,
      confirm.value
    )
    regMsg.value = data.Msg
    messager.value = true
    if (success(data)) {
      successToAuth()
    }
  } catch (error) {
    regMsg.value = error.Msg
    messager.value = true
  }
}

function successToAuth() {
  router.push('/login')
}
</script>

<template>
  <v-card width="600">
    <v-card-title class="text-h4 text-center my-2" min-height="100"
      >Register</v-card-title
    >
    <v-divider></v-divider>
    <v-sheet width="400" class="mx-auto my-2">
      <v-form
        fast-fail
        @submit.prevent
        :model-value="valid"
        @update:modelValue="valid = $event"
      >
        <v-text-field
          v-model="email"
          label="Email"
          :rules="emailRules"
        ></v-text-field>
        <v-text-field
          v-model="username"
          label="Username"
          :rules="usernameRules"
        ></v-text-field>
        <v-text-field
          v-model="password"
          label="Password"
          :rules="passwordRules"
          type="password"
        ></v-text-field>
        <v-text-field
          v-model="confirm"
          label="Confirm your password."
          :rules="confirmRules"
          type="password"
        ></v-text-field>
        <v-btn
          type="submit"
          color="primary"
          block
          class="mt-2"
          @click="sumbitRegisterForm"
          >Register</v-btn
        >
      </v-form>
    </v-sheet>
    <br />
  </v-card>
  <v-snackbar v-model="messager" :timeout="timeout">
    {{ regMsg }}
    <template v-slot:actions>
      <v-btn color="blue" variant="text" @click="messager = false">
        Close
      </v-btn>
    </template>
  </v-snackbar>
</template>
