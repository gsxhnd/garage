import { http, Response } from "@/utils/http";
import { MovieTag } from "./movie_tag";
import { MovieActor } from "./movie_actor";

export interface Movie {
  id: number;
  code: string;
  title: string;
  cover?: string | null;
  publishDate?: Date | null;
  director?: string | null;
  produceCompany?: string | null;
  publishCompany?: string | null;
  series?: string | null;
  createdAt?: Date | null;
  updatedAt?: Date | null;
}

export interface MovieInfo {
  movie: Movie;
  actors: MovieActor[];
  tags: MovieTag[];
}

export const GetMovies = () => {
  return http.get<Response<Array<Movie>>>("/movie");
};

export const GetMovieInfo = (code: string) => {
  return http.get<Response<MovieInfo>>(`/movie/info/${code}`);
};

export const CreateMovies = () => {
  return http.post<Response<null>>("/movie");
};
