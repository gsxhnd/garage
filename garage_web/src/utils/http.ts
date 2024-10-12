import axios from "axios";

const http = axios.create({
  baseURL: "",
  timeout: 1000,
});

http.interceptors.request.use(
  (config) => {
    return config;
  },
  (_err) => {}
);

http.interceptors.response.use(
  (resp) => {
    return resp;
  },
  (_err) => {}
);

export { http };
