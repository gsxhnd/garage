import { http, Response } from "@/utils/http";

export interface Actor {
  id: number;
  name: string;
  aliasName?: string | null;
  cover?: string | null;
  createdAt?: Date | null;
  updatedAt?: Date | null;
}

export const GetActors = () => {
  return http.get<Response<Actor[]>>("/actor");
};
