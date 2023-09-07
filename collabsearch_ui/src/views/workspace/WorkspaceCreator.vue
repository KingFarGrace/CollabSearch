<script setup>
import { ref } from 'vue'
import VueDatePicker from '@vuepic/vue-datepicker'
import '@vuepic/vue-datepicker/dist/main.css'
import { workspaceCreateService } from '@/api/workspace'
import { useAccountStore } from '@/stores/account'
import { datetimeToString } from '@/utils/formatter'
import { storeToRefs } from 'pinia'
// Messager
var messager = ref(false)
var timeout = ref(2000)
var crtMsg = ref('')
// Form
var topic = ref('')
var description = ref('')
var dueTime = ref()
var due = ref('')
var valid = ref(false)
var titleRules = [
  (value) => {
    if (value?.length > 0 && value?.length <= 255) return true
    return 'Title cannot be empty or longer than 255 characters.'
  }
]

var descriptionRules = [
  (value) => {
    if (value?.length > 0) return true
    return 'Description cannot be empty.'
  }
]

var store = useAccountStore()
var { uid } = storeToRefs(store)

async function submitWorkspaceForm() {
  if (!valid.value) return
  try {
    console.log(description.value)
    const data = await workspaceCreateService(
      topic.value,
      description.value,
      uid.value,
      due.value
    )
    // TODO: Message color.
    crtMsg.value = data.Msg
    messager.value = true
  } catch (error) {
    crtMsg.value = error.Msg
    messager.value = true
  }
}

function getDatetime() {
  due.value = datetimeToString(dueTime.value)
}
</script>

<template>
  <v-card class="mx-auto w-100 overflow-visible" density="comfortable">
    <v-card-title class="text-h5 text-center">Create Workspace</v-card-title>
    <v-divider></v-divider>
    <br />
    <v-sheet class="mx-auto w-75">
      <v-form
        fast-fail
        @submit.prevent
        :model-value="valid"
        @update:modelValue="valid = $event"
      >
        <v-text-field
          counter
          v-model="topic"
          label="Topic"
          :rules="titleRules"
          class="mx-auto w-75"
        ></v-text-field>

        <v-container class="w-100"
          ><v-textarea
            counter
            clearable
            clear-icon="mdi-close-circle"
            v-model="description"
            label="Description"
            :rules="descriptionRules"
            :model-value="description"
            class="w-100 mx-auto"
          ></v-textarea
        ></v-container>

        <input
          class="float-middle"
          type="datetime-local"
          id="datetime"
          name="datetime"
          v-if="false"
        />
        <v-row
          ><VueDatePicker
            v-model="dueTime"
            format="dd-MM-yyyy HH:mm:ss"
            @update:model-value="getDatetime"
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
  <v-snackbar v-model="messager" :timeout="timeout">
    {{ crtMsg }}
    <template v-slot:actions>
      <v-btn color="blue" variant="text" @click="messager = false">
        Close
      </v-btn>
    </template>
  </v-snackbar>
</template>
