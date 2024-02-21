import axios from "axios";
import { InternalAxiosRequestConfig, AxiosResponse } from "axios";

const http = axios.create({
  baseURL: "",
  timeout: 5000,
  headers: {
    "Content-Type": "application/json",
  },
});

http.interceptors.request.use(
  (config: InternalAxiosRequestConfig) => {
    console.log(config.baseURL);
    return config;
  },
  (error: any) => {
    console.error(error);
    return Promise.reject(error);
  }
);

http.interceptors.response.use(
  (resp: AxiosResponse) => {
    return resp;
  },
  (error: any) => {
    console.error("resp error: ", error);
    return Promise.reject(error);
  }
);

export default http;
