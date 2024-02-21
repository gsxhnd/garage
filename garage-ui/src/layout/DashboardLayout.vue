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

      <v-app-bar-title>Application Bar</v-app-bar-title>
    </v-app-bar>
    <v-navigation-drawer
      model-value
      class="pt-4"
      color="grey-lighten-3"
      expand-on-hover
      :rail="layoutStore.app_bar.rail"
    >
      <!-- <v-avatar
        v-for="n in 6"
        :key="n"
        :color="`grey-${n === 1 ? 'darken' : 'lighten'}-1`"
        :size="n === 1 ? 36 : 20"
        class="d-block text-center mx-auto mb-9"
      ></v-avatar> -->
    </v-navigation-drawer>

    <v-main>
      <!--  -->
      <router-view> </router-view>
    </v-main>
  </v-app>
</template>

<script setup lang="ts">
import { useLayoutStore } from "@/stores/layout";
const layoutStore = useLayoutStore();
</script>

<style scoped lang="less"></style>
