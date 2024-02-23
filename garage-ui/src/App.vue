<script setup lang="ts">
import { RouterView } from "vue-router";
import { testStore } from "@/stores/index";
import { userLogin, testApi } from "@/api/user";

testStore();
userLogin();
testApi().catch((err) => {
  console.log(err);
});
</script>

<template>
  <router-view class="router-view" />
</template>

<style scoped lang="less"></style>
