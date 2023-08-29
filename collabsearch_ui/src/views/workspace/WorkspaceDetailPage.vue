<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { storeToRefs } from 'pinia'
import { useResultStore } from '@/stores/result'
import { useWorkspaceStore } from '@/stores/workspace'
import {
  searchSetResultService,
  searchSetNoteService,
  searchGetNotesService
} from '@/api/search'
// SnackBar
var messager = ref(false)
var timeout = ref(2000)
var crtMsg = ref('')
// Notes
var dialog = ref(false)
var notes = ref([])
// Scroller
var call = ref(0)
function onScroll() {
  call.value++
}
// Go back
const router = useRouter()
// Form
const store = useResultStore()
var { currentPhrase, title, description, distributions, resource } =
  storeToRefs(store)
var comment = ref('')
var score = ref(0)
var ticks = {
  0: 'Totally useless to our topic.',
  2.5: 'Maybe relevant to our topic, maybe.',
  5: 'This is just what I need.'
}
var valid = ref(false)
var commentRules = [
  (value) => {
    if (value?.length > 0 && value?.length <= 255) return true
    return 'Comment cannot be empty or longer than 255 characters.'
  }
]
var { wid } = storeToRefs(useWorkspaceStore())
async function setNote() {
  if (!valid.value) {
    return
  }
  try {
    var setNoteResp = await searchSetNoteService(
      wid.value,
      resource.value,
      comment.value,
      score.value
    )
    crtMsg.value = setNoteResp.Msg
    messager.value = true
  } catch (error) {
    crtMsg.value = error.Msg
    messager.value = true
    console.log(error)
    return
  }
  try {
    await searchSetResultService(wid.value, resource.value, currentPhrase.value)
  } catch (error) {
    crtMsg.value = error.Msg
    messager.value = true
    console.log(error)
  }
}
async function getNotes() {
  dialog.value = true
  try {
    var data = await searchGetNotesService(wid.value, resource.value)
    notes.value = data.Notes
  } catch (error) {
    console.log(error)
  }
  console.log()
}
onMounted(() => {
  distributions.value
})
</script>
<template>
  <v-card class="mx-auto w-100" variant="outlined">
    <v-row>
      <v-col cols="9">
        <v-card v-scroll.self="onScroll" class="overflow-y-auto" height="400"
          ><v-banner class="justify-center text-h5 font-weight-light" sticky>
            {{ title }}
          </v-banner>

          <br />

          <v-card-text class="font-weight-regular text-subtitle-1">{{
            description
          }}</v-card-text>
        </v-card>
      </v-col>

      <v-col cols="3">
        <v-row>
          <v-btn
            prepend-icon="mdi-home-search"
            stacked
            variant="text"
            class="my-16 mx-auto w-auto"
            @click="router.go(-1)"
            >Back to workspace</v-btn
          >
        </v-row>
        <v-row>
          <v-btn
            prepend-icon="mdi-satellite-uplink"
            stacked
            variant="text"
            class="my-16 mx-auto w-auto"
            :href="resource"
            >Go to resource page</v-btn
          >
        </v-row>
      </v-col>
    </v-row>

    <v-row>
      <v-card class="w-100" variant="text"
        ><v-card-title class="text-center text-h5">Distributions</v-card-title>
        <v-list variant="tonal">
          <v-list-item
            v-for="distribution in distributions"
            :key="distribution?.title?.en"
            :title="distribution?.title?.en || 'Empty Distribution Title'"
            lines="three"
          >
            <template v-slot:append>
              <v-btn-group>
                <v-btn
                  v-if="distribution?.access_url"
                  icon="mdi-information"
                  variant="tonal"
                  :href="distribution?.access_url[0]"
                  class="mx-4"
                ></v-btn>
                <v-btn
                  v-if="distribution?.download_url"
                  icon="mdi-download-circle"
                  variant="tonal"
                  :href="distribution?.download_url[0]"
                  class="mx-4"
                ></v-btn>
              </v-btn-group>
            </template>
            <v-tooltip activator="parent" location="top center">{{
              distribution?.description?.en || 'No description.'
            }}</v-tooltip>
          </v-list-item>
        </v-list>
      </v-card>
    </v-row>
    <v-row>
      <v-sheet class="w-100 mx-auto">
        <v-form
          fast-fail
          @submit.prevent
          :model-value="valid"
          @update:modelValue="valid = $event"
        >
          <v-textarea
            v-model="comment"
            label="Leave your comment to this searching result. Recommened comment should contain your email and never be empty."
            :rules="commentRules"
          ></v-textarea>
          <v-slider
            :ticks="ticks"
            :max="5"
            step="1"
            v-model="score"
            thumb-label
            show-ticks="always"
            tick-size="4"
            class="mx-16"
          ></v-slider>
          <v-btn
            type="submit"
            variant="tonal"
            class="my-2 mx-auto"
            @click="setNote"
            >Leave your note
          </v-btn>

          <v-btn
            type="submit"
            variant="tonal"
            class="my-2 mx-auto float-right"
            @click="getNotes"
            >Show others notes</v-btn
          >
        </v-form>
      </v-sheet>
    </v-row>
  </v-card>
  <v-dialog v-model="dialog" class="w-auto">
    <v-card class="w-75 mx-auto">
      <v-list lines="three">
        <v-list-item
          v-for="note in notes"
          :key="note.wid"
          :title="note.content"
          :subtitle="'Score: ' + note.feedback + '/5'"
        ></v-list-item>
      </v-list>
      <v-card-actions>
        <v-btn color="primary" class="mx-auto w-50" @click="dialog = false"
          >Close</v-btn
        >
      </v-card-actions>
    </v-card>
  </v-dialog>
  <v-snackbar v-model="messager" :timeout="timeout">
    {{ crtMsg }}
    <template v-slot:actions>
      <v-btn color="blue" variant="text" @click="messager = false">
        Close
      </v-btn>
    </template>
  </v-snackbar>
</template>
