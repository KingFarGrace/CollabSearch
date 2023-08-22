<script setup>
import { onMounted } from 'vue'
import { useUWStore } from '@/stores/userWorkspace'
import { useAccountStore } from '@/stores/account'
import { useWorkspaceStore } from '@/stores/workspace'
import { useRouter } from 'vue-router'
import { storeToRefs } from 'pinia'
import { workspaceGetCreatedService } from '@/api/workspace'
// List
var store = useUWStore()
var { createdWorkspaces } = storeToRefs(store)
var { setCreated, getCreatedWorkspaceByWid } = store
var { uid } = storeToRefs(useAccountStore())
var { setWorkspace } = useWorkspaceStore()
async function getCreatedWorkspaces() {
  try {
    const data = await workspaceGetCreatedService(uid.value)
    const created = data.Workspaces
    createdWorkspaces.value = created
    setCreated(created)
  } catch (error) {
    setCreated()
  }
}
const router = useRouter()
function toWorkspaceManager(wid) {
  const workspace = getCreatedWorkspaceByWid(wid)
  setWorkspace(workspace)
  router.push('/workspace/manage')
}
onMounted(() => {
  getCreatedWorkspaces()
})
</script>

<template>
  <v-container>
    <v-row>
      <span class="text-h5 text-center mx-auto mb-4 w-100"
        >Workspaces I Created.</span
      >
    </v-row>
    <v-row>
      <v-card class="mx-auto w-75">
        <v-list lines="three">
          <v-list-item
            v-for="workspace in createdWorkspaces"
            :key="workspace?.wid"
            :title="workspace?.topic"
            :subtitle="workspace?.description"
            @click="toWorkspaceManager(workspace.wid)"
          ></v-list-item>
        </v-list>
      </v-card>
    </v-row>
  </v-container>
</template>
