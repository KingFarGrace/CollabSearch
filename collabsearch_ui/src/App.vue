<script setup>
import { useRouter } from 'vue-router'
import { VContainer } from 'vuetify/lib/components/index.mjs'
import { useAccountStore } from '@/stores/account'
import { storeToRefs } from 'pinia'
var store = useAccountStore()
var { isLogin, email, username, avatar } = storeToRefs(store)
var defaultAvatar = store.defaultAvatarPath
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
          :prepend-avatar="defaultAvatar"
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
      <v-main class="d-flex align-center justify-center mx-auto">
        <router-view></router-view>
      </v-main>
    </v-container>
  </v-layout>
</template>

<style scoped></style>
