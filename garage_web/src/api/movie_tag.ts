import { http, Response } from "@/utils/http";

export interface MovieTag {
  id: number;
  movie_id: number;
  tag_id: number;
  tag_name: string;
}

export const GetMovieTag = () => {
  return http.get<Response<MovieTag[]>>("/movie_tag");
};
