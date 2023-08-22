import instance from '@/utils/axios'
import { outside } from '@/utils/axios'

const groupPath = '/workspace/result'

export const searchSetResultService = (uid, wid, title, url, phrase) => {
  return instance.post(groupPath, { uid, wid, title, url, phrase })
}

export const searchGetHintService = (uid, wid, phrase) => {
  return instance.get(groupPath + '/search', { uid, wid, phrase })
}

export const searchRequestService = (phrase, page, limit) => {
  const requestParams = {
    q: phrase,
    page: page,
    limit: limit,
    filter: 'dataset',
    sort: 'modified'
  }
  const url = '/api?' + new URLSearchParams(requestParams)
  return outside(url)
}
