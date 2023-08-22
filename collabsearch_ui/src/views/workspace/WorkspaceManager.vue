<script setup>
import { ref, computed, onMounted } from 'vue'
import VueDatePicker from '@vuepic/vue-datepicker'
import '@vuepic/vue-datepicker/dist/main.css'
import { storeToRefs } from 'pinia'
import { useWorkspaceStore } from '@/stores/workspace'
import { useAccountStore } from '@/stores/account'
import { workspaceUpdateService } from '@/api/workspace'
import { datetimeToString } from '@/utils/formatter'
import { stringToDatetime } from '@/utils/formatter'
import { accountSearchUsersByEmailService } from '@/api/account'
// Messager
var messager = ref(false)
var timeout = ref(2000)
var updMsg = ref('')
// Dialog
var { uid } = useAccountStore()
var dialog = ref(false)
var loading = ref(false)
var key = ref('')
var users = ref([])
var filteredUsers = computed(() => {
  return users.value.filter((item) => item.uid != uid)
})
// Datetime picker
var dueTime = ref()
// Form
const store = useWorkspaceStore()
var { topic, description, due } = storeToRefs(store)
var { wid, handler, setDue } = store // wid and handler can't be changed
var valid = ref(false)
var disabled = ref(true)

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

async function updateWorkspaceForm() {
  if (!valid.value) return
  // Change update form status.
  // When !disabled after changed, do nothing and return.
  disabled.value = !disabled.value
  if (!disabled.value) {
    return
  }
  try {
    const data = await workspaceUpdateService(
      wid,
      topic.value,
      description.value,
      handler,
      due.value
    )
    messager.value = true
    updMsg.value = data.Msg
  } catch (error) {
    messager.value = true
    updMsg.value = error.Msg
  }
}

function getDatetime() {
  setDue(datetimeToString(dueTime.value))
}

async function searchUserByEmail() {
  if (!key.value) {
    return
  }
  try {
    const data = await accountSearchUsersByEmailService(key.value)
    users.value = data.Users
  } catch (error) {
    updMsg.value = error.Msg
    messager.value = true
  }
}

// TODO: Add user to workspace, user will receive a notification.
async function addUser() {}

onMounted(() => {
  dueTime.value = stringToDatetime(due.value)
})
</script>

<template>
  <v-card class="mx-auto w-100 overflow-visible" density="comfortable">
    <v-card-item>
      <v-card-title class="text-h5 text-center">Update Workspace</v-card-title>
    </v-card-item>
    <v-divider></v-divider>
    <v-sheet class="mx-auto w-75">
      <v-form
        fast-fail
        @submit.prevent
        :disabled="disabled"
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

        <v-container fluid width="300">
          <v-textarea
            counter
            clearable
            clear-icon="mdi-close-circle"
            label="Description"
            v-model="description"
            :rules="descriptionRules"
            :model-value="description"
            class="w-100"
          ></v-textarea>
        </v-container>

        <VueDatePicker
          v-model="dueTime"
          format="dd-MM-yyyy HH:mm:ss"
          @update:model-value="getDatetime"
          :disabled="disabled"
          class="w-75 mx-auto"
        ></VueDatePicker>
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
          v-if="filteredUsers.length === 0"
          title="No users found."
          prepend-icon="mdi-account-question-outline"
        ></v-list-item>
        <v-list-item
          v-for="(user, i) in filteredUsers"
          :key="i"
          :title="user?.email"
          :subtitle="user?.username"
          :prepend-avatar="user?.avatar || '/src/assets/OIP.jpg'"
          @click="console.log(user)"
        ></v-list-item>
      </v-list>
      <v-card-actions>
        <v-btn color="primary" block @click="dialog = false">Close</v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
  <v-snackbar v-model="messager" :timeout="timeout">
    {{ updMsg }}
    <template v-slot:actions>
      <v-btn color="blue" variant="text" @click="messager = false">
        Close
      </v-btn>
    </template>
  </v-snackbar>
</template>
