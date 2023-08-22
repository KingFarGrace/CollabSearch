import instance from '@/utils/axios'

const groupPath = '/user'

export const accountCheckStatusService = () => {
  return instance.get(groupPath + '/check')
}

export const accountLoginService = (email, password) => {
  const jsonObj = { email, password }
  return instance.post(groupPath + '/login', jsonObj)
}

export const accountRegisterService = (email, username, password, confirm) => {
  return instance.post(groupPath, { email, username, password, confirm })
}

export const accountUpdateService = (uid, email, username, profile, avatar) => {
  return instance.put(groupPath, { uid, email, username, profile, avatar })
}

export const accountJoinWorkspaceService = (uid, wid) => {
  return instance.post(groupPath + '/workspace', { uid, wid })
}

export const accountSearchUsersByEmailService = (key) => {
  return instance.get(groupPath + '/' + key)
}
