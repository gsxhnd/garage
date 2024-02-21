<script setup lang="ts">
import { RouterView } from "vue-router";
import { testStore } from "@/stores/index";

testStore();
</script>

<template>
  <router-view class="router-view" />
</template>

<style scoped lang="less"></style>
