<script setup>
import { ref } from 'vue'
var page = ref(1)
var loading = ref(false)
var drawer = ref(true)
var rail = ref(true)
var infoDialog = ref(false)
var kgDialog = ref(false)
var clbDialog = ref(false)
var phrase = ref('')
var searchItems = ref([])
function search() {
  loading.value = true
  searchItems.value = [
    { title: 'one', detail: 'one detail' },
    { title: 'two', detail: 'two detail' }
  ]
  loading.value = false
}
</script>
<template>
  <v-container class="mx-auto">
    <v-navigation-drawer
      v-model="drawer"
      :rail="rail"
      permanent
      @click="rail = false"
      location="top"
    >
      <v-list-item prepend-icon="mdi-file-cabinet" title="Workspace Detail" nav>
        <template v-slot:append>
          <v-btn
            variant="text"
            icon="mdi-chevron-down"
            @click.stop="rail = !rail"
          ></v-btn>
        </template>
      </v-list-item>

      <v-divider></v-divider>

      <v-list density="compact" nav>
        <v-list-item
          prepend-icon="mdi-information-box"
          title="Information"
          @click="infoDialog = true"
        ></v-list-item>
        <v-list-item
          prepend-icon="mdi-graph-outline"
          title="Knowledge Graph"
          @click="kgDialog = true"
        ></v-list-item>
        <v-list-item
          prepend-icon="mdi-account-group-outline"
          title="Collaborators"
          @click="clbDialog = true"
        ></v-list-item>
      </v-list>
    </v-navigation-drawer>
    <v-row>
      <v-autocomplete
        v-model="phrase"
        :items="searchItems"
        append-inner-icon="mdi-search-web"
        auto-select-first
        class="flex-full-width my-1"
        density="comfortable"
        item-props
        menu-icon=""
        placeholder="Search in data.europa.eu"
        prepend-inner-icon="mdi-magnify"
        rounded
        theme="light"
        variant="solo"
        @click:appendInner="search"
      ></v-autocomplete>
    </v-row>
    <v-row
      ><v-card class="mx-auto my-1 w-100" v-if="searchItems.value != null"
        ><v-list lines="three">
          <v-list-item
            v-for="searchItem in searchItems"
            :key="searchItem.title"
            :title="searchItem.title"
            subtitle="..."
          ></v-list-item></v-list></v-card
    ></v-row>

    <v-row
      ><v-card class="text-center mx-auto">
        <v-pagination
          v-model="page"
          :length="15"
          :total-visible="7"
        ></v-pagination></v-card
    ></v-row>
  </v-container>
  <v-dialog v-model="infoDialog" class="w-auto">
    <v-card class="mx-auto w-75">
      <v-card-title> Topic here </v-card-title>
      <v-card-subtitle>Due here</v-card-subtitle>
      <v-card-text>Description here</v-card-text>
      <v-card-actions>
        <v-btn color="primary" class="mx-auto w-50" @click="infoDialog = false"
          >Close</v-btn
        >
      </v-card-actions>
    </v-card>
  </v-dialog>
  <v-dialog v-model="kgDialog" class="w-auto">
    <v-card class="w-100">
      <v-card-text>
        Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod
        tempor incididunt ut labore et dolore magna aliqua.
      </v-card-text>
      <v-card-actions>
        <v-btn color="primary" class="mx-auto w-25" @click="kgDialog = false"
          >Close Dialog</v-btn
        >
      </v-card-actions>
    </v-card>
  </v-dialog>
  <v-dialog v-model="clbDialog" class="w-auto">
    <v-card class="mx-auto w-50">
      <v-list>
        <v-list-item title="handler" prepend-icon="mdi-account"></v-list-item>
      </v-list>
      <v-list>
        <v-list-item
          title="user1"
          prepend-icon="mdi-account-outline"
        ></v-list-item>
        <v-list-item
          title="user2"
          prepend-icon="mdi-account-outline"
        ></v-list-item>
      </v-list>
      <v-btn color="primary" block @click="clbDialog = false"
        >Close Dialog</v-btn
      >
    </v-card>
  </v-dialog>
</template>
