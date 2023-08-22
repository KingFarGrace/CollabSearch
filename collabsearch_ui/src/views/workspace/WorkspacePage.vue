<script setup>
import { computed, onMounted, ref } from 'vue'
import { useRouter } from 'vue-router'
import { storeToRefs } from 'pinia'
import { useWorkspaceStore } from '@/stores/workspace'
import { searchRequestService } from '@/api/search'
import { workspaceGetParticipantsService } from '@/api/workspace'
// Pagination
var page = ref(1)
var pageNum = computed(() => {
  return page.value - 1
})
var maxPage = ref(0)
var MAX_CAP = 10
// Nav
var drawer = ref(true)
var rail = ref(true)
// Info Diglog
const store = useWorkspaceStore()
var { wid, topic, description, due } = storeToRefs(store)
var infoDialog = ref(false)
// Knowledge Graph Dialog
var kgDialog = ref(false)
// Participants Dialog
var partDialog = ref(false)
var participants = ref([])
var handler = computed(() => {
  return participants.value[0]
})
var collaborators = computed(() => {
  return participants.value.slice(1)
})
// Search
var loading = ref(false)
var phrase = ref('')
var searchResults = ref([])
async function search() {
  if (!phrase.value) {
    return
  }
  loading.value = true
  try {
    const resp = await searchRequestService(
      phrase.value,
      pageNum.value,
      MAX_CAP
    )
    maxPage.value = Math.ceil(resp.data.result.count / MAX_CAP)
    searchResults.value = resp.data.result.results
    console.log(searchResults.value)
  } catch (error) {
    console.log(error)
  } finally {
    loading.value = false
  }
}

async function getParticipants() {
  try {
    const data = await workspaceGetParticipantsService(wid.value)
    participants.value = data.Users
  } catch (error) {
    console.log(error)
  }
}

const router = useRouter()
function goDetailPage() {
  router.push('/workspace/detail')
}

onMounted(() => {
  getParticipants()
})
</script>
<template>
  <v-container class="mx-auto">
    <v-navigation-drawer
      v-model="drawer"
      :rail="rail"
      permanent
      @click="rail = false"
      location="top"
    >
      <v-list-item prepend-icon="mdi-file-cabinet" title="Workspace Detail" nav>
        <template v-slot:append>
          <v-btn
            variant="text"
            icon="mdi-chevron-down"
            @click.stop="rail = !rail"
          ></v-btn>
        </template>
      </v-list-item>

      <v-divider></v-divider>

      <v-list density="compact" nav>
        <v-list-item
          prepend-icon="mdi-information-box"
          title="Information"
          @click="infoDialog = true"
        ></v-list-item>
        <v-list-item
          prepend-icon="mdi-graph-outline"
          title="Knowledge Graph"
          @click="kgDialog = true"
        ></v-list-item>
        <v-list-item
          prepend-icon="mdi-account-group-outline"
          title="Collaborators"
          @click="partDialog = true"
        ></v-list-item>
      </v-list>
    </v-navigation-drawer>
    <v-row>
      <v-text-field
        class="mb-2"
        v-model="phrase"
        :loading="loading"
        density="compact"
        variant="solo"
        label="Search in data.europa.eu, sort by update time in descending order."
        append-inner-icon="mdi-magnify"
        single-line
        hide-details
        @click:append-inner="search"
      ></v-text-field>
    </v-row>
    <v-row
      ><v-card class="mx-auto my-2 w-100" v-if="searchResults.length !== 0"
        ><v-list lines="three">
          <v-list-item
            v-for="result in searchResults"
            :key="result.id"
            :title="result?.title?.en"
            :subtitle="result?.description?.en"
            prepend-icon="mdi-circle-small"
            @click="console.log(result)"
          ></v-list-item></v-list></v-card
    ></v-row>

    <v-row
      ><v-card class="text-center mx-auto my-2">
        <v-pagination
          v-if="maxPage !== 0"
          v-model="page"
          :length="maxPage"
          :total-visible="7"
          @update:modelValue="search"
        ></v-pagination></v-card
    ></v-row>
  </v-container>
  <v-dialog v-model="infoDialog" class="w-auto">
    <v-card class="mx-auto w-75">
      <v-card-item prepend-icon="mdi-tag-outline">
        <v-card-title> Topic: {{ topic }} </v-card-title>
        <v-card-subtitle>Due: {{ due }}</v-card-subtitle>
      </v-card-item>
      <v-textarea
        v-model="description"
        prepend-inner-icon="mdi-attachment"
        disabled
        auto-grow
        class="font-weight-black font-weight-bold"
      ></v-textarea>
      <v-card-actions>
        <v-btn color="primary" class="mx-auto w-50" @click="infoDialog = false"
          >Close</v-btn
        >
      </v-card-actions>
    </v-card>
  </v-dialog>
  <v-dialog v-model="kgDialog" class="w-auto">
    <v-card class="w-100">
      <v-card-text>
        Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod
        tempor incididunt ut labore et dolore magna aliqua.
      </v-card-text>
      <v-card-actions>
        <v-btn color="primary" class="mx-auto w-25" @click="kgDialog = false"
          >Close Dialog</v-btn
        >
      </v-card-actions>
    </v-card>
  </v-dialog>
  <v-dialog v-model="partDialog" class="w-auto">
    <v-card class="mx-auto w-50">
      <v-list>
        <v-list-item
          prepend-icon="mdi-account"
          :title="handler.email"
          :subtitle="handler.username"
        ></v-list-item>
      </v-list>
      <v-list density="compact" nav>
        <v-list-item
          v-for="collaborator in collaborators"
          :key="collaborator.uid"
          prepend-icon="mdi-account-outline"
          :title="collaborator.email"
          :subtitle="collaborator.username"
        ></v-list-item>
      </v-list>
      <v-btn color="primary" block @click="partDialog = false"
        >Close Dialog</v-btn
      >
    </v-card>
  </v-dialog>
</template>
