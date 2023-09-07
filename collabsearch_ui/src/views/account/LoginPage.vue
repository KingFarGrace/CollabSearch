<script setup>
import { ref } from 'vue'
import { useAccountStore } from '@/stores/account'
import { accountCheckStatusService, accountLoginService } from '@/api/account'
import { useRouter } from 'vue-router'
const router = useRouter()
const store = useAccountStore()
const { setAll } = store
var email = ref('')
var password = ref('')
var valid = ref(false)
var messager = ref(false)
var timeout = ref(2000)
var loginMsg = ref('')
var emailRules = [
  (value) => {
    const reg = /^[a-zA-Z0-9._-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,6}$/
    if (reg.test(value)) {
      return true
    }
    return 'Invalid email address.'
  }
]

var passwordRules = [
  (value) => {
    const reg = /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[\W_]).{6,16}$/
    if (reg.test(value)) {
      return true
    }
    return 'Password should contain at least One uppercase&lowercase letter, One number, One special english charactor and length between 6 and 16.'
  }
]

async function submitLoginForm() {
  if (!valid.value) return
  try {
    const data = await accountCheckStatusService()
    const user = data.users[0]
    setAll(user)
    loginMsg.value = "You're already online."
    messager.value = true
    successToAccount()
    return
  } catch (error) {
    // Intentionally empty. We will try to login next.
    console.error(error)
  }
  try {
    const data = await accountLoginService(email.value, password.value)
    const user = data.Users[0]
    setAll(user)
    loginMsg.value = 'Login successfully.'
    messager.value = true
    successToAccount()
  } catch (error) {
    console.error(error)
    loginMsg.value = error.Msg
    messager.value = true
  }
}

function successToAccount() {
  router.push('/user')
}
</script>

<template>
  <v-card width="600">
    <v-card-title class="text-h4 text-center my-2" min-height="100"
      >Login</v-card-title
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
          v-model="password"
          label="Password"
          :rules="passwordRules"
          type="password"
        ></v-text-field>

        <v-btn
          type="submit"
          color="info"
          block
          class="my-2"
          @click="submitLoginForm"
          >Login</v-btn
        >
      </v-form>
    </v-sheet>
    <br />
    <v-divider></v-divider>
    <br />
    <span class="float-right mr-16"
      >New user?
      <router-link to="/reg">Click to create your accout.</router-link></span
    >
    <br />
    <br />
  </v-card>
  <v-snackbar v-model="messager" :timeout="timeout">
    {{ loginMsg }}
    <template v-slot:actions>
      <v-btn color="blue" variant="text" @click="messager = false">
        Close
      </v-btn>
    </template>
  </v-snackbar>
</template>
