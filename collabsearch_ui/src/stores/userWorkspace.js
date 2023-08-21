import { ref } from 'vue'
import { defineStore } from 'pinia'
import {
  workspaceGetJoinedService,
  workspaceGetCreatedService
} from '@/api/workspace'

export const useUWStore = defineStore(
  'user-workspace',
  () => {
    const joinedWorkspaces = ref([])
    const setJoined = (joined) => {
      joinedWorkspaces.value = joined
    }
    const getJoinedWorkspaces = async ({ uid }) => {
      const joinedWorkspaces = await workspaceGetJoinedService({ uid })
      setJoined(joinedWorkspaces)
    }
    const createdWorkspaces = ref([])
    const setCreated = (created) => {
      createdWorkspaces.value = created
    }
    const getCreatedWorkspacesByWid = (wid) => {
      for (let i = 0; i < createdWorkspaces.value.length; i++) {
        const workspace = createdWorkspaces[i]
        if (workspace.wid === wid) return workspace
      }
    }
    const getCreatedWorkspaces = async ({ uid }) => {
      const createdWorkspaces = await workspaceGetCreatedService({ uid })
      setCreated(createdWorkspaces)
    }
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
