import { ref } from 'vue'
import { defineStore } from 'pinia'

export const useUWStore = defineStore(
  'user-workspace',
  () => {
    const joinedWorkspaces = ref([])
    const getJoinedWorkspaces = async () => {}
    const createdWorkspaces = ref([])
    const getCreatedWorkspaces = async () => {}
    return {
      joinedWorkspaces,
      getJoinedWorkspaces,
      createdWorkspaces,
      getCreatedWorkspaces
    }
  },
  { persist: true }
)
