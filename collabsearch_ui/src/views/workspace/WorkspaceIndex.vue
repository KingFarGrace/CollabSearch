<script setup>
import { ref, provide } from 'vue'
import { useRouter } from 'vue-router'
import { storeToRefs } from 'pinia'
import { useUWStore } from '@/stores/userWorkspace'
var router = useRouter()
var { joinedWorkspaces, createdWorkspaces } = storeToRefs(useUWStore())
provide('joined', joinedWorkspaces)
provide('created', createdWorkspaces)
var tab = ref(null)
function getWorkspacesJoined() {
  router.push('/workspace')
}
function getWorkspacesCreated() {
  router.push('/workspace/handler')
}
function goCreateWorkspace() {
  router.push('/workspace/create')
}
</script>

<template>
  <v-container>
    <v-row
      ><v-card class="float-left">
        <v-tabs v-model="tab" align-tabs="center" color="deep-purple-accent-4">
          <v-tab value="1" @click="getWorkspacesJoined"
            >Workspaces I joined</v-tab
          >
          <v-tab value="2" @click="getWorkspacesCreated"
            >Workspaces I created</v-tab
          >
          <v-tab value="3" @click="goCreateWorkspace">Create a workspace</v-tab>
        </v-tabs>
      </v-card></v-row
    >
    <v-row>
      <v-card class="w-100"><router-view></router-view></v-card>
    </v-row>
  </v-container>
</template>
