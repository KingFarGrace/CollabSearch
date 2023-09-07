import { computed, ref } from 'vue'
import { defineStore } from 'pinia'
import { stringToDatetime, datetimeToString } from '@/utils/formatter'

export const useWorkspaceStore = defineStore(
  'workspace',
  () => {
    const wid = ref(0)
    const setWid = (id) => {
      wid.value = id
    }
    const topic = ref('')
    const setTopic = (content) => {
      topic.value = content
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
    const computedDue = computed({
      get: () => stringToDatetime(due.value),
      set: (datetime) => {
        due.value = datetimeToString(datetime)
      }
    })
    const setWorkspace = (workspace) => {
      if (!workspace) return
      setWid(workspace.wid)
      setTopic(workspace.topic)
      setDescription(workspace.description)
      setDue(workspace.due)
      setHandler(workspace.handler)
    }
    return {
      wid,
      topic,
      setTopic,
      description,
      setDescription,
      handler,
      setHandler,
      due,
      setDue,
      computedDue,
      setWorkspace
    }
  },
  { persist: true }
)
