import instance from '@/utils/axios'
import { outside } from '@/utils/axios'

const groupPath = '/workspace/result'

export const searchSetResultService = (wid, url, phrase) => {
  return instance.post(groupPath, { wid, url, phrase })
}

export const searchGetHintsService = (wid, phrase) => {
  return instance.post(groupPath + '/search', { wid, phrase })
}

export const searchSetNoteService = (wid, url, content, feedback) => {
  return instance.post(groupPath + '/note', { wid, url, content, feedback })
}

export const searchGetNotesService = (wid, url) => {
  return instance.post(groupPath + '/note/all', { wid, url })
}

export const searchGetAvgScoreService = (wid, url) => {
  return instance.post(groupPath + '/avg', { wid, url })
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
