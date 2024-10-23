<template>
  <div class="movie">
    <div class="movie-content">
      <div class="movie-search"></div>
      <div class="movie-pic">
        <MoviePic />
      </div>
    </div>
    <div class="movie-panel">
      <MoviePanel />
    </div>
  </div>
</template>

<script setup lang="ts">
import MoviePic from "./movie/MoviePic.vue";
import MoviePanel from "./movie/MoviePanel.vue";
import { onBeforeMount } from "vue";
import { useMovieStore } from "@/stores/movie";
const movieStore = useMovieStore();

onBeforeMount(async () => {
  await movieStore.getMovies();
});
</script>

<style scoped lang="scss">
.movie {
  background-color: rgb(236, 236, 236);
  padding-top: 20px;
  flex: 1;
  display: flex;
  &-content {
    padding: 20px;
    width: 100%;
    flex-grow: 1;
  }
  &-panel {
    width: 400px;
  }
}
</style>
