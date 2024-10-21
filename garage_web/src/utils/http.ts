import axios from "axios";
import { AxiosResponse } from "axios";

const http = axios.create({
  baseURL: "/api/v1",
  timeout: 1000,
});

http.interceptors.request.use(
  (config) => {
    return config;
  },
  (_err) => {}
);

http.interceptors.response.use(
  (resp: AxiosResponse<Response<any>>) => {
    console.log(resp);
    if (resp.status !== 200) {
    }
    if (resp.data.code !== 0) {
      alert(resp.data.message);
    }
    return resp;
  },
  (_err) => {}
);

export { http };

export interface Response<T> {
  code: number;
  message: string;
  data: T;
}
