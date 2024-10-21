import { http, Response } from "@/utils/http";

export interface MovieActor {
  id: number;
  movieId: number;
  actorId: number;
  actorName?: string | null;
}

export const GetMovieActor = () => {
  return http.get<Response<MovieActor[]>>("/movie_actor");
};
