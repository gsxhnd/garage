import { http, Response } from "@/utils/http";

export interface MovieTag {
  id: number;
  movieId: number;
  tagId: number;
  tagName?: string | null;
}

export const GetMovieTag = () => {
  return http.get<Response<MovieTag[]>>("/movie_tag");
};
