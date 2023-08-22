<script setup>
import { onMounted } from 'vue'
import { useUWStore } from '@/stores/userWorkspace'
import { useAccountStore } from '@/stores/account'
import { useWorkspaceStore } from '@/stores/workspace'
import { useRouter } from 'vue-router'
import { storeToRefs } from 'pinia'
import { workspaceGetJoinedService } from '@/api/workspace'
// List
var store = useUWStore()
var { joinedWorkspaces } = storeToRefs(store)
var { setJoined, getJoinedWorkspaceByWid } = store
var { uid } = storeToRefs(useAccountStore())
var { setWorkspace } = useWorkspaceStore()
async function getJoinedWorkspaces() {
  try {
    const data = await workspaceGetJoinedService(uid.value)
    const joined = data.Workspaces
    joinedWorkspaces.value = joined
    setJoined(joined)
  } catch (error) {
    setJoined()
  }
}
const router = useRouter()
function toWorkspace(wid) {
  const workspace = getJoinedWorkspaceByWid(wid)
  setWorkspace(workspace)
  router.push('/workspace/search')
}
onMounted(() => {
  getJoinedWorkspaces()
})
</script>

<template>
  <v-container>
    <v-row>
      <span class="text-h5 text-center mx-auto mb-4 w-100"
        >Workspaces I Joined.</span
      >
    </v-row>
    <v-row>
      <v-card class="mx-auto w-75">
        <v-list lines="three">
          <v-list-item
            v-for="workspace in joinedWorkspaces"
            :key="workspace?.wid"
            :title="workspace?.topic"
            :subtitle="workspace?.description"
            @click="toWorkspace(workspace.wid)"
          ></v-list-item>
        </v-list>
      </v-card>
    </v-row>
  </v-container>
</template>
