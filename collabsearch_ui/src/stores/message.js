import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useMessageStore = defineStore(
  'message',
  () => {
    const messages = ref([])
    const setMessages = (msgs) => {
      messages.value = msgs
    }
    return {
      messages,
      setMessages
    }
  },
  {
    persist: true
  }
)
