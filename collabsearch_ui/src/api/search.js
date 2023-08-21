import instance from '@/utils/axios'

const groupPath = '/workspace/result'

export const searchSetResultService = ({ uid, wid, title, url, phrase }) => {
  return instance.post(groupPath, { uid, wid, title, url, phrase })
}

export const searchGetHintService = ({ uid, wid, phrase }) => {
  return instance.get(groupPath + '/search', { uid, wid, phrase })
}
