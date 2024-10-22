import axios, { AxiosResponse } from "axios";

export interface Response<T> {
  code: number;
  message: string;
  data: T;
}

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
    if (resp.status !== 200) {
      alert(resp.status);
    }
    if (resp.data.code !== 0) {
      alert(resp.data.message);
    }
    return resp;
  },
  (err) => {
    console.log(err.response);
  }
);

export { http };
