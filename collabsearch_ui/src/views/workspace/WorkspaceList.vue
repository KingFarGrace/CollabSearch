<script setup>
import { inject } from 'vue'
var joinedWorkspaces = inject('joined')
</script>

<template>
  <span class="text-h5 text-center mx-auto w-100">Workspaces I Joined.</span>
  <v-divider></v-divider>
  <v-list
    :items="joinedWorkspaces"
    item-title="title"
    item-value="description"
  ></v-list>
</template>
