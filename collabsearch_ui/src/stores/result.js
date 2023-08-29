import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useResultStore = defineStore(
  'result',
  () => {
    // Searching
    const currentPhrase = ref('')
    const setCurrentPhrase = (phrase) => {
      currentPhrase.value = phrase
    }
    const currentPage = ref(1)
    const setCurrentPage = (num) => {
      currentPage.value = num
    }
    const maxPage = ref(0)
    const setMaxPage = (num) => {
      maxPage.value = num
    }
    const resultsInPage = ref([])
    const setResultsInPage = async (results) => {
      resultsInPage.value = results
    }
    const colors = ref(new Map())
    const setColors = (newColors) => {
      colors.value = newColors
    }
    const resetSearchingPage = () => {
      setCurrentPhrase('')
      setCurrentPage(1)
      setMaxPage(0)
      setResultsInPage([])
      setColors(new Map())
    }
    // Details
    const title = ref('')
    const description = ref('')
    // destribution in destributions:
    // destribution
    // - title
    // - description
    // - format
    // -- id
    // - access_url
    // - download_url
    const distributions = ref(null)
    const resource = ref('')
    const setResult = (result) => {
      title.value = result?.title?.en
      description.value = result?.description?.en
      distributions.value = result?.distributions
      resource.value = result?.resource
    }
    return {
      currentPhrase,
      setCurrentPhrase,
      currentPage,
      setCurrentPage,
      maxPage,
      setMaxPage,
      resultsInPage,
      setResultsInPage,
      colors,
      setColors,
      resetSearchingPage,
      title,
      description,
      distributions,
      resource,
      setResult
    }
  },
  { persist: true }
)
