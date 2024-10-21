<template>
  <div
    class="dropdown"
    :class="{ 'is-active': showDropdown }"
    v-on-click-outside="closeDropdown"
  >
    <div class="dropdown-trigger">
      <button
        class="button"
        @click="showDropdown = !showDropdown"
        aria-haspopup="true"
        aria-controls="dropdown-menu"
      >
        <span>Dropdown button</span>
        <span class="icon is-small">
          <i class="fas fa-angle-down" aria-hidden="true"></i>
        </span>
      </button>
    </div>
    <div class="dropdown-menu" id="dropdown-menu" role="menu">
      <div class="dropdown-content">
        <a href="#" class="dropdown-item"> Dropdown item </a>
        <a class="dropdown-item"> Other dropdown item </a>
        <a href="#" class="dropdown-item is-active"> Active dropdown item </a>
        <a href="#" class="dropdown-item"> Other dropdown item </a>
        <hr class="dropdown-divider" />
        <a href="#" class="dropdown-item"> With a divider </a>
      </div>
    </div>
  </div>
</template>
<style scoped lang="scss"></style>
<script setup lang="ts">
import { defineEmits, ref } from "vue";
import { vOnClickOutside } from "@vueuse/components";

const showDropdown = ref(false);
const emit = defineEmits<{
  selected: [v: string];
}>();

function closeDropdown() {
  showDropdown.value = false;
}

function aaa() {
  emit("selected", "123");
}
</script>
