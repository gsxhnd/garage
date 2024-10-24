<template>
  <div
    v-if="show"
    ref="target"
    class="tag-tree"
    :style="{ top: position.y + 'px', left: position.x + 'px' }"
  >
    tag
  </div>
</template>

<script setup lang="ts">
// import { ref, useTemplateRef } from "vue";

const { show, position = { x: 0, y: 0 } } = defineProps<{
  show?: boolean;
  position?: {
    x: number;
    y: number;
  };
}>();
</script>

<style scoped lang="scss">
.tag-tree {
  position: fixed;
}
</style>
