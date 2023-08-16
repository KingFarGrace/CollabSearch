import { ref } from 'vue'
import { defineStore } from 'pinia'

export const useWorkspaceStore = defineStore('workspace', () => {
  const wid = ref(0)
  const title = ref('')
  const setTitle = (content) => {
    title.value = content
  }
  const description = ref('')
  const setDescription = (content) => {
    description.value = content
  }
  const handler = ref(0)
  const setHandler = (content) => {
    handler.value = content
  }
  const due = ref('')
  const setDue = (content) => {
    due.value = content
  }
  const getWorkspaceInfo = async () => {}
  return {
    wid,
    title,
    setTitle,
    description,
    setDescription,
    handler,
    setHandler,
    due,
    setDue,
    getWorkspaceInfo
  }
})
