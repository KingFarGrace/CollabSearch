import instance from '@/utils/axios'

const groupPath = '/msg'

export const sendMessageService = (sender, receiver, content) => {
  return instance.post(groupPath, { sender, receiver, content, read: false })
}

export const getMessagesService = (receiver) => {
  return instance.get(groupPath + '/' + receiver)
}

export const setReadService = (sender, receiver, content) => {
  return instance.put(groupPath, { sender, receiver, content, read: true })
}

export const removeMessageService = (sender, receiver, content, read) => {
  return instance.delete(groupPath, { sender, receiver, content, read })
}
