<script setup>
import { inject } from 'vue'
var createdWorkspaces = inject('created')
</script>

<template>
  <v-card class="mx-auto" max-width="300">
    <span class="text-h5">Workspaces I Created.</span>
    <v-list
      :items="createdWorkspaces"
      item-title="title"
      item-value="description"
    ></v-list>
  </v-card>
</template>
