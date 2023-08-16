<script setup>
import { useRouter } from 'vue-router'
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
          prepend-avatar="src\\assets\\OIP.jpg"
          title="Login"
          @click="goAuthPage"
          v-if="false"
        >
        </v-list-item>
      </v-list>
      <v-divider></v-divider>

      <v-list density="compact" nav>
        <v-list-item
          prepend-icon="mdi-account-multiple"
          title="My Workspaces"
          value="myWorkspaces"
          @click="goWorkspacePage"
        ></v-list-item>
        <v-list-item
          prepend-icon="mdi-star"
          title="Starred"
          value="starred"
        ></v-list-item>
      </v-list>
    </v-navigation-drawer>

    <v-app-bar title="Application bar"></v-app-bar>

    <v-main
      class="d-flex align-center justify-center"
      style="min-height: 300px"
    >
      <router-view></router-view>
    </v-main>
  </v-layout>
</template>

<style scoped></style>
