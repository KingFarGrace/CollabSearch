<script setup>
import { ref } from 'vue'
var title = ref('')
var description = ref('')
var due = ref('')
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

function submitWorkspaceForm() {
  alert('Create Workspace!')
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
      <v-form fast-fail @submit.prevent>
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
        <input
          class="float-middle"
          type="datetime-local"
          id="datetime"
          name="datetime"
          v-if="false"
        />
        <!-- TODO: Datetime picker here. -->
        <v-text-field
          v-model="due"
          label="Due"
          :rules="dueRules"
          class="mx-auto w-50"
          v-if="true"
        ></v-text-field>
        <v-btn type="submit" block class="mt-2" @click="submitWorkspaceForm"
          >Submit</v-btn
        >
      </v-form>
    </v-sheet>
    <br />
  </v-card>
</template>
