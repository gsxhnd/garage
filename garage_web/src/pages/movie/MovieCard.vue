<template>
  <div class="movie-card" :class="{ active: active }">
    <img
      :src="imgUrl"
      onerror="this.src='/img_not_fount.svg'; this.onerror=null;"
    />
    <div class="movie-card-detail">
      <span>{{ data?.code }}</span>
      <span>{{ data?.title }}</span>
      <span>{{ data?.publishDate }}</span>
    </div>
  </div>
</template>
<script setup lang="ts">
import { ref } from "vue";
import { Movie } from "@/api/movie";

const { data, active } = defineProps<{
  data?: Movie;
  active?: boolean;
}>();

const imgUrl = ref(`/api/v1/img/${data?.id}/${data?.code}`);
</script>
<style scoped lang="scss">
.movie-card {
  cursor: pointer;
  display: flex;
  flex-direction: column;
  box-shadow: 0px 1px 5px 1px rgba(0, 0, 0, 0.1);
  border-radius: 5px;
  background-color: var(--app-card-background);
  color: var(--app-card-color);
  position: relative;

  &.active {
    box-shadow: 0px 0px 5px 1px rgba(0, 127, 245, 0.493);
  }

  &:hover {
    background-color: var(--app-card-hover-background);
    box-shadow: 0px 0px 5px 1px rgba(0, 127, 245, 0.493);
  }

  &-detail {
    padding: 0 5px;
    position: absolute;
    bottom: 0;
    color: #ffffff;
  }
  span {
    align-self: start;
  }
}
</style>
