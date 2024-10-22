import { http, Response } from "@/utils/http";
import { AxiosRequestConfig } from "axios";

export interface Actor {
  id: number;
  name: string;
  aliasName?: string | null;
  cover?: string | null;
  createdAt?: Date | null;
  updatedAt?: Date | null;
}

export const CreateActor = (data: Array<Actor>) => {
  return http.post<Response<null>>("/actor", data);
};

export const DeleteActor = (ids: Array<number>) => {
  let config: AxiosRequestConfig<Array<number>> = {
    data: ids,
  };
  return http.delete<Response<null>>(`/actor`, config);
};

export const GetActors = () => {
  return http.get<Response<Actor[]>>("/actor");
};
