import { ref } from 'vue'
import { defineStore } from 'pinia'

export const useUWStore = defineStore(
  'user-workspace',
  () => {
    const joinedWorkspaces = ref([])
    const getJoinedWorkspaces = async () => {}
    const createdWorkspaces = ref([])
    const getCreatedWorkspacesByWid = (wid) => {
      for (let i = 0; i < createdWorkspaces.value.length; i++) {
        const workspace = createdWorkspaces[i]
        if (workspace.wid === wid) return workspace
      }
    }
    const getCreatedWorkspaces = async () => {}
    return {
      joinedWorkspaces,
      getJoinedWorkspaces,
      createdWorkspaces,
      getCreatedWorkspacesByWid,
      getCreatedWorkspaces
    }
  },
  { persist: true }
)
