<script setup>
import { ref } from 'vue'
var email = ref('')
var password = ref('')
var emailRules = [
  (value) => {
    const reg = /^[a-zA-Z0-9._-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,6}$/
    if (reg.test(value)) return true
    return 'Invalid email address.'
  }
]

var passwordRules = [
  (value) => {
    const reg = /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[\W_]).{6,16}$/
    if (reg.test(value)) return true
    return 'Password should contain at least One uppercase&lowercase letter, One number, One special english charactor and length between 6 and 16.'
  }
]

function submitLoginForm() {
  alert('Login!')
}
</script>

<template>
  <v-card width="400">
    <v-card-item>
      <v-card-title>Login</v-card-title>
    </v-card-item>
    <v-sheet width="300" class="mx-auto">
      <v-form fast-fail @submit.prevent>
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

        <v-btn type="submit" block class="mt-2" @click="submitLoginForm"
          >Submit</v-btn
        >
      </v-form>
    </v-sheet>
    <br />
    <v-divider inset></v-divider>
    New user? <router-link to="/reg">Click to create your accout.</router-link>
    <v-card-text></v-card-text>
  </v-card>
</template>
