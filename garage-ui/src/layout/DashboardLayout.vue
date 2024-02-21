<template>
  <v-app id="inspire">
    <v-navigation-drawer
      model-value
      class="pt-4"
      color="grey-lighten-3"
      expand-on-hover
      rail
    >
      <v-avatar
        v-for="n in 6"
        :key="n"
        :color="`grey-${n === 1 ? 'darken' : 'lighten'}-1`"
        :size="n === 1 ? 36 : 20"
        class="d-block text-center mx-auto mb-9"
      ></v-avatar>
    </v-navigation-drawer>

    <v-main>
      <!--  -->
      <router-view> </router-view>
    </v-main>
  </v-app>
</template>

<script setup lang="ts"></script>

<style scoped lang="less"></style>
