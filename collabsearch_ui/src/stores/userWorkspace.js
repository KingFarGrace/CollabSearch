import { ref } from 'vue'
import { defineStore } from 'pinia'

export const useUWStore = defineStore(
  'user-workspace',
  () => {
    const joinedWorkspaces = ref([])
    const setJoined = (joined) => {
      joinedWorkspaces.value = joined
    }
    const getJoinedWorkspaceByWid = (wid) => {
      for (let i = 0; i < joinedWorkspaces.value.length; i++) {
        const workspace = joinedWorkspaces.value[i]
        if (workspace.wid === wid) return workspace
      }
    }
    const createdWorkspaces = ref([])
    const setCreated = (created) => {
      createdWorkspaces.value = created
    }
    const getCreatedWorkspaceByWid = (wid) => {
      for (let i = 0; i < createdWorkspaces.value.length; i++) {
        const workspace = createdWorkspaces.value[i]
        if (workspace.wid === wid) return workspace
      }
    }
    return {
      joinedWorkspaces,
      setJoined,
      getJoinedWorkspaceByWid,
      createdWorkspaces,
      setCreated,
      getCreatedWorkspaceByWid
    }
  },
  { persist: true }
)
