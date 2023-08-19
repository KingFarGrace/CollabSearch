<script setup>
import { inject } from 'vue'
var createdWorkspaces = inject('created')
</script>

<template>
  <span class="text-h5 text-center mx-auto w-100">Workspaces I Created.</span>
  <v-divider></v-divider>
  <v-list
    :items="createdWorkspaces"
    item-title="title"
    item-value="description"
  ></v-list>
</template>
