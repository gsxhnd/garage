<script setup lang="ts">
import { ref } from "vue";

const count = ref(0);
</script>

<template>
  <div class="card">
    <v-btn prepend-icon="$vuetify" variant="tonal" @click="count++">
      count is {{ count }}
    </v-btn>
    <p>
      Edit
      <code>components/HelloWorld.vue</code> to test HMR
    </p>

    <v-select
      :items="[
        'California',
        'Colorado',
        'Florida',
        'Georgia',
        'Texas',
        'Wyoming',
      ]"
      variant="outlined"
    ></v-select>
  </div>
</template>

<style scoped>
.read-the-docs {
  color: #888;
}
</style>
