<template>
  <div class="sidebar" :class="{ expand: isExpand }" v-auto-animate>
    <div class="header">
      <img v-tooltip="'logo'" src="/vite.svg" />
    </div>
    <div class="sidebar-menu">
      <template v-for="v in menu">
        <div
          class="menu-item"
          :class="[
            isExpand ? 'expand' : 'collapsed',
            activeMenu == v.router ? 'active' : '',
          ]"
          @click="toPage(v.router)"
          v-tooltip="v.tooltip"
        >
          <span class="icon">
            <i :class="v.icon"></i>
          </span>
          <span v-if="isExpand" class="text">{{ v.tooltip }}</span>
        </div>
      </template>
    </div>
    <div class="footer" @click="toggleExpand">
      <span class="icon">
        <i class="fas fa-bars"></i>
      </span>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onBeforeMount } from "vue";
import { useRoute, useRouter } from "vue-router";

const router = useRouter();
const route = useRoute();
const isExpand = ref(true);
const activeMenu = ref("");
const menu = ref([
  {
    name: "movie",
    tooltip: "Movie",
    icon: "fas fa-film",
    router: "movie",
  },
  {
    name: "actor",
    tooltip: "Actor",
    icon: "fas fa-user",
    router: "actor",
  },
  { name: "anime", tooltip: "Anime", icon: "fas fa-tv", router: "anime" },
  {
    name: "tag",
    tooltip: "Tag",
    icon: "fas fa-tags",
    router: "tag",
  },
  {
    name: "setting",
    tooltip: "Setting",
    icon: "fas fa-gear",
    router: "setting",
  },
]);

onBeforeMount(() => {
  activeMenu.value = route.name as string;
});

function toggleExpand() {
  isExpand.value = !isExpand.value;
}

function toPage(name: string) {
  router.push({ name });
  activeMenu.value = name;
}
</script>

<style scoped lang="scss">
.sidebar {
  background-color: var(--app-sidebar-background);
  color: var(--app-sidebar-color);
  height: 100vh;
  width: 50px;
  &.expand {
    width: 200px;
  }

  &-menu {
    display: flex;
    flex-direction: column;
    align-items: flex-start;
    cursor: pointer;
    .menu-item {
      &.collapsed {
        justify-content: space-around;
      }
      &.active {
        background-color: var(--app-sidebar-menu-item-active);
      }
      padding: 10px 5px;
      display: flex;
      width: 100%;
      .icon {
        margin-right: 5px;
      }
      span {
        display: flex;
        align-items: center;
      }
      &:hover {
        background-color: var(--app-sidebar-menu-item-hover);
      }
    }
  }

  .header {
    padding-top: 10px;
  }

  .footer {
    position: absolute;
    bottom: 0;
    display: flex;
    justify-content: center;
    width: 100%;
  }
}
</style>
