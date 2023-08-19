<script setup>
import { useRouter } from 'vue-router'
import { VContainer } from 'vuetify/lib/components/index.mjs'
var router = useRouter()
function goAuthPage() {
  router.push('/login')
}
function goAccountPage() {
  router.push('/user')
}
function goWorkspacePage() {
  router.push('/workspace')
}
function goWorkspaceHandlerPage() {
  router.push('/workspace/handler')
}
function goWorkspaceCreatorPage() {
  router.push('/workspace/create')
}
</script>

<template>
  <v-layout class="rounded rounded-md">
    <v-navigation-drawer expand-on-hover rail>
      <v-list>
        <v-list-item
          prepend-avatar="https://randomuser.me/api/portraits/women/85.jpg"
          title="Sandra Adams"
          subtitle="sandra_a88@gmailcom"
          @click="goAccountPage"
          v-if="true"
        ></v-list-item>
        <v-list-item
          prepend-avatar="src/assets/OIP.jpg"
          title="Login"
          @click="goAuthPage"
          v-if="false"
        >
        </v-list-item>
      </v-list>
      <v-divider></v-divider>
      <v-list nav>
        <v-list-item
          prepend-icon="mdi-briefcase"
          title="Workspace I Joined"
          @click="goWorkspacePage"
        ></v-list-item>
        <v-list-item
          prepend-icon="mdi-briefcase-outline"
          title="Workspace I Created"
          @click="goWorkspaceHandlerPage"
        ></v-list-item>
        <v-list-item
          prepend-icon="mdi-folder-plus-outline"
          title="Workspace Creator"
          @click="goWorkspaceCreatorPage"
        ></v-list-item>
      </v-list>
    </v-navigation-drawer>
    <v-container>
      <v-main
        class="d-flex align-center justify-center"
        style="min-height: 300px"
      >
        <router-view></router-view>
      </v-main>
    </v-container>
  </v-layout>
</template>

<style scoped></style>
