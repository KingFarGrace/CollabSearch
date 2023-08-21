import instance from '@/utils/axios'

const groupPath = '/workspace'

export const workspaceCreateService = ({
  topic,
  description,
  handler,
  due
}) => {
  return instance.post(groupPath, { topic, description, handler, due })
}

export const workspaceGetJoinedService = ({ uid }) => {
  return instance.get(groupPath, { uid })
}

export const workspaceGetCreatedService = ({ uid }) => {
  return instance.get(groupPath + '/handler', { uid })
}

export const workspaceUpdateService = ({
  wid,
  topic,
  description,
  handler,
  due
}) => {
  return instance.put(groupPath, { wid, topic, description, handler, due })
}

export const workspaceGetInfoService = (wid) => {
  return instance.get(groupPath + wid)
}
