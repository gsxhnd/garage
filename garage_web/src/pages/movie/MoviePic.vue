<template>
  <div class="grid is-gap-0 is-col-min-1">
    <div v-for="i in movieStore.movies" class="cell">
      <MovieCard
        @click="clickMovie(i.code)"
        :data="i"
        :active="movieStore.selectMovieInfo?.movie.id == i.id"
      />
    </div>
  </div>
</template>

<script setup lang="ts">
import MovieCard from "./MovieCard.vue";
import { useMovieStore } from "@/stores/movie";
const movieStore = useMovieStore();

function clickMovie(code: string) {
  movieStore.selectMovie(code);
}
</script>

<style scoped lang="scss"></style>
