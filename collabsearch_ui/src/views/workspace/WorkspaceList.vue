<script setup>
import { inject } from 'vue'
var joinedWorkspaces = inject('joined')
</script>

<template>
  <v-card class="mx-auto" max-width="300">
    <span class="text-h5">Workspaces I Joined.</span>
    <v-list
      :items="joinedWorkspaces"
      item-title="title"
      item-value="description"
    ></v-list>
  </v-card>
</template>
