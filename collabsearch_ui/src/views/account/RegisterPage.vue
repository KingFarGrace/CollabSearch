<script setup>
import { ref } from 'vue'
var email = ref('')
var username = ref('')
var password = ref('')
var confirm = ref('')
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

function sumbitRegisterForm() {
  alert('Register!')
}
</script>

<template>
  <v-card width="600">
    <v-card-title class="text-h4 text-center" min-height="100"
      >Register</v-card-title
    >
    <br />
    <v-divider></v-divider>
    <br />
    <v-sheet width="400" class="mx-auto">
      <v-form fast-fail @submit.prevent>
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
        <v-btn type="submit" block class="mt-2" @click="sumbitRegisterForm"
          >Register</v-btn
        >
      </v-form>
    </v-sheet>
    <br />
  </v-card>
</template>
