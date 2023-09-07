import { DateTime } from 'luxon'
export const datetimeToString = (time) => {
  const date = new Date(time)

  const day = String(date.getDate()).padStart(2, '0')
  const month = String(date.getMonth() + 1).padStart(2, '0') // +1 because months are 0-indexed
  const year = date.getFullYear()

  const hours = String(date.getHours()).padStart(2, '0')
  const minutes = String(date.getMinutes()).padStart(2, '0')
  const seconds = String(date.getSeconds()).padStart(2, '0')

  return `${day}-${month}-${year} ${hours}:${minutes}:${seconds}`
}

export const stringToDatetime = (dateTimeStr) => {
  return DateTime.fromFormat(dateTimeStr, 'dd-MM-yyyy HH:mm:ss')
}
