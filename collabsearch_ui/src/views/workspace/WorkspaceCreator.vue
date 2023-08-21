<script setup>
import { ref } from 'vue'
import VueDatePicker from '@vuepic/vue-datepicker'
import '@vuepic/vue-datepicker/dist/main.css'
var topic = ref('')
var description = ref('')
var due = ref()
var titleRules = [
  (value) => {
    if (value?.length > 0 && value?.length <= 255) return true
    return 'Title cannot be empty or longer than 255 charactors.'
  }
]

var descriptionRules = [
  (value) => {
    if (value?.length > 0) return true
    return 'Description cannot be empty.'
  }
]

function submitWorkspaceForm() {
  alert('Create Workspace!')
}

function printDatetime() {
  const date = new Date(due.value)

  const day = String(date.getDate()).padStart(2, '0')
  const month = String(date.getMonth() + 1).padStart(2, '0') // +1 because months are 0-indexed
  const year = date.getFullYear()

  const hours = String(date.getHours()).padStart(2, '0')
  const minutes = String(date.getMinutes()).padStart(2, '0')
  const seconds = String(date.getSeconds()).padStart(2, '0')

  console.log(`${day}-${month}-${year} ${hours}:${minutes}:${seconds}`)
}
</script>

<template>
  <v-card class="mx-auto w-100 overflow-visible" density="comfortable">
    <v-card-title class="text-h5 text-center">Create Workspace</v-card-title>
    <v-divider></v-divider>
    <br />
    <v-sheet class="mx-auto w-75 h-auto">
      <v-form fast-fail @submit.prevent>
        <v-text-field
          counter
          v-model="topic"
          label="Topic"
          :rules="titleRules"
          class="mx-auto w-75"
        ></v-text-field>

        <v-container fluid width="300">
          <v-textarea
            counter
            clearable
            clear-icon="mdi-close-circle"
            label="Description"
            :rules="descriptionRules"
            :model-value="description"
            class="w-100"
          ></v-textarea>
        </v-container>
        <input
          class="float-middle"
          type="datetime-local"
          id="datetime"
          name="datetime"
          v-if="false"
        />
        <v-row
          ><VueDatePicker
            v-model="due"
            format="dd/MM/yyyy HH:mm:ss"
            @update:model-value="printDatetime"
            class="w-75 mx-auto"
          ></VueDatePicker>
          <v-btn type="submit" class="w-25 mx-auto" @click="submitWorkspaceForm"
            >Create</v-btn
          ></v-row
        >
        <br />
      </v-form>
    </v-sheet>
  </v-card>
</template>

<style>
.mh-form {
  height: 100%;
}
</style>
