<template>
  <div class="movie">
    <div class="movie-content">
      <div class="movie-search">
        <MovieSearch />
      </div>
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
import MovieSearch from "./movie/MovieSearch.vue";
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
  flex: 1;
  display: flex;
  &-content {
    background-color: var(--app-surface-background-2);
    padding: 20px;
    width: 100%;
    flex-grow: 1;
  }
  &-panel {
    background-color: var(--app-surface-background-1);
    padding: 10px;
    width: 400px;
  }
}
</style>
