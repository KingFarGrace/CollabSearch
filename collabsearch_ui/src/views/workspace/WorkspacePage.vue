<script setup>
import { ref } from 'vue'
var phrase = ref('')
</script>
<template>
  <v-responsive class="mx-auto" max-width="400">
    <v-text-field
      v-model="phrase"
      prepend-inner-icon="mdi-search-web"
      label="Input your searching key phrase."
      type="input"
      hint="Default searching engine is google, your mates' searching history will be shown below."
    ></v-text-field>
  </v-responsive>
</template>
