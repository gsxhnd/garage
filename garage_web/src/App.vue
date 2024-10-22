<template>
  <router-view />
</template>

<script setup lang="ts">
import { onBeforeMount } from "vue";

onBeforeMount(async () => {
  const doc = document.documentElement;
  doc.dataset.theme = "oneDark";
});
</script>
<style scoped></style>
