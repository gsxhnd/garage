<template>
  <v-app id="inspire">
    <v-app-bar :elevation="0" rounded>
      <template v-slot:prepend>
        <v-app-bar-nav-icon
          :border="false"
          @click="() => (layoutStore.app_bar.rail = !layoutStore.app_bar.rail)"
        ></v-app-bar-nav-icon>
        <a style="width: 20px">
          <v-img src="vite.svg"></v-img>
        </a>
      </template>
    </v-app-bar>
    <v-navigation-drawer
      model-value
      class="pt-4"
      expand-on-hover
      :rail="layoutStore.app_bar.rail"
    >
      <v-list density="compact" nav>
        <v-list-item
          prepend-icon="mdi-folder"
          title="My Files"
          value="myfiles"
        ></v-list-item>
        <v-list-item
          prepend-icon="mdi-account-multiple"
          title="Shared with me"
          value="shared"
        ></v-list-item>
        <v-list-item
          prepend-icon="mdi-star"
          title="Starred"
          value="starred"
        ></v-list-item>
      </v-list>
    </v-navigation-drawer>

    <v-main>
      <router-view></router-view>
    </v-main>
  </v-app>
</template>

<script setup lang="ts">
import { useLayoutStore } from "@/stores/layout";
const layoutStore = useLayoutStore();
</script>

<style scoped lang="less">
.v-app-bar {
  background-color: #e9ebf2;
}
</style>
