<template>
  <div class="movie-card" :class="{ active: active }">
    <img :src="imgUrl" />
    <span>{{ data?.code }}</span>
    <span>{{ data?.title }}</span>
    <span>{{ data?.publishDate }}</span>
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
  padding: 5px;
  cursor: pointer;
  display: flex;
  flex-direction: column;
  box-shadow: 0px 1px 5px 1px rgba(0, 0, 0, 0.1);
  border-radius: 5px;

  &.active {
    box-shadow: 0px 0px 5px 1px rgba(0, 127, 245, 0.493);
  }

  &:hover {
    box-shadow: 0px 0px 5px 1px rgba(0, 127, 245, 0.493);
  }
  span {
    align-self: start;
  }
}
</style>
