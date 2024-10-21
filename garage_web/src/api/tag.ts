import { http, Response } from "@/utils/http";

export interface Tag {
  id: number;
  name: string;
  pid: number;
  createdAt?: Date | null;
  updatedAt?: Date | null;
}

export const GetTags = () => {
  return http.get<Response<Tag[]>>("/tag");
};
