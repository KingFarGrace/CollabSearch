export const convertBinaryToImageUrl = (binaryData) => {
  const blob = new Blob([binaryData], { type: 'image/jpeg' })
  return URL.createObjectURL(blob)
}
