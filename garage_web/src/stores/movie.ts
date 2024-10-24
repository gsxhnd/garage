import { defineStore } from "pinia";
import { Ref, ref } from "vue";
import { Movie, MovieInfo, GetMovies, GetMovieInfo } from "@/api/movie";

export const useMovieStore = defineStore("movie", () => {
  const movies: Ref<Array<Movie>> = ref([]);
  const selectMovieInfo: Ref<MovieInfo | null> = ref(null);

  async function getMovies() {
    await GetMovies().then((data) => {
      movies.value = data.data.data;
    });
  }

  async function selectMovie(code: string) {
    selectMovieInfo.value = null;
    await GetMovieInfo(code).then((data) => {
      selectMovieInfo.value = data.data.data;
    });
  }

  return { movies, selectMovieInfo, selectMovie, getMovies };
});
