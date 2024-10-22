import { http, Response } from "@/utils/http";
import { AxiosRequestConfig } from "axios";

export interface Tag {
  id: number;
  name: string;
  pid: number;
  createdAt?: Date | null;
  updatedAt?: Date | null;
}

export const CreateTag = (data: Array<Tag>) => {
  return http.post<Response<null>>("/tag", data);
};

export const DeleteTag = (data: Array<number>) => {
  let config: AxiosRequestConfig<Array<number>> = {
    data: data,
  };
  return http.delete<Response<null>>("/tag", config);
};

export const UpdateTag = (data: Tag) => {
  return http.put<Response<null>>("/tag", data);
};

export const GetTag = () => {
  return http.get<Response<Tag[]>>("/tag");
};
