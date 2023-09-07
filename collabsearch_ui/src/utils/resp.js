export const success = (resp) => {
  if (resp.Code % 10 === 0) return true
  return false
}
