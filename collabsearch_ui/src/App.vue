<script setup>
import { useRouter } from 'vue-router'
import { VContainer } from 'vuetify/lib/components/index.mjs'
import { useAccountStore } from '@/stores/account'
import { storeToRefs } from 'pinia'
import { useTheme } from 'vuetify'
const theme = useTheme()
function toggleTheme() {
  theme.global.name.value = theme.global.current.value.dark ? 'light' : 'dark'
}
var store = useAccountStore()
var { isLogin, email, username, avatar } = storeToRefs(store)
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
          :prepend-avatar="avatar"
          :title="username"
          :subtitle="email"
          @click="goAccountPage"
          v-if="isLogin"
        ></v-list-item>
        <v-list-item
          prepend-avatar="/src/assets/OIP.jpg"
          title="Login"
          @click="goAuthPage"
          v-if="!isLogin"
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
      <v-app-bar>
        <v-spacer></v-spacer>
        <v-btn icon>
          <v-icon>mdi-message-reply-text-outline</v-icon>
        </v-btn>
        <template v-slot:append>
          <v-btn icon="mdi-theme-light-dark" @click="toggleTheme"></v-btn>
        </template>
      </v-app-bar>
      <v-main class="d-flex align-center justify-center mx-auto">
        <router-view></router-view>
      </v-main>
    </v-container>
  </v-layout>
</template>

<style scoped></style>
