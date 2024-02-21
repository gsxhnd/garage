import http from "@/utils/http";
import { AxiosResponse } from "axios";

export const userLogin = (): Promise<AxiosResponse<any, any>> =>
  http.get("/user/logion");
