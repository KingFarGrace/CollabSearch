<script setup>
import { ref } from 'vue'
// import { storeToRef } from 'pinia'
// import { useUWStore } from '@/stores/userWorkspace'
// var { getCreatedWorkspaceByWid } = useUWStore()
// var workspace = getCreatedWorkspaceByWid()
var title = ref('')
var description = ref('')
var due = ref('')
var disabled = ref(true)
var dialog = ref(false)
var loading = ref(false)
var key = ref('')
var users = ref([])
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

var dueRules = [
  (value) => {
    const reg =
      /^\d{4}-((0[1-9]|1[0-2])-(0[1-9]|1\d|2[0-8])|(0[13-9]|1[0-2])-(29|30)|(0[13578]|1[02])-31) (00|0\d|1\d|2[0-3]):[0-5]\d:[0-5]\d$/

    if (reg.test(value)) return true
    return "Datetime format must be valid and its format should be like '2000-12-31 00:00:00'."
  }
]

function updateWorkspaceForm() {
  disabled.value = !disabled.value
}

function searchUserByEmail() {
  if (key.value === '' || key.value === null) {
    return
  }
  users.value = [{ email: 'kingfargrace@gmail.com', username: 'king' }]
}
</script>

<template>
  <v-card class="mx-auto w-100">
    <v-card-item>
      <v-card-title class="text-h5 text-center"
        >Create a workspace</v-card-title
      >
    </v-card-item>
    <v-divider></v-divider>
    <v-sheet class="mx-auto w-75">
      <v-form fast-fail @submit.prevent :disabled="disabled">
        <v-text-field
          counter
          v-model="title"
          label="Title"
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

        <v-text-field
          v-model="due"
          label="Due"
          :rules="dueRules"
          class="mx-auto w-50"
        ></v-text-field>
        <v-btn
          type="submit"
          class="mt-2 mx-auto w-50"
          v-if="disabled"
          @click="updateWorkspaceForm"
          >Update</v-btn
        >
        <v-btn
          type="submit"
          class="mt-2 mx-auto w-50"
          v-if="!disabled"
          @click="updateWorkspaceForm"
          >Finish</v-btn
        >
        <v-btn type="submit" class="mt-2 mx-auto w-50" @click="dialog = true"
          >Add Collaborators</v-btn
        >
      </v-form>
    </v-sheet>
    <br />
  </v-card>
  <v-dialog v-model="dialog" class="w-auto">
    <v-card class="w-75 mx-auto">
      <v-card-text>
        <v-text-field
          v-model="key"
          :loading="loading"
          density="compact"
          variant="solo"
          label="Search user emails."
          append-inner-icon="mdi-magnify"
          single-line
          hide-details
          @click:append-inner="searchUserByEmail"
        ></v-text-field>
      </v-card-text>

      <v-list>
        <v-list-item
          v-for="(user, i) in users"
          :key="i"
          :title="user.email"
          :subtitle="user.username"
        ></v-list-item>
      </v-list>
      <v-card-actions>
        <v-btn color="primary" block @click="dialog = false">Close</v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>
