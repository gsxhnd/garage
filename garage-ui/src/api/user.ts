import { httpv2 } from "@/utils/http";

const userLogin = () => httpv2.get("user/logion");

const testApi = () => {
  return httpv2.get("api");
};

export { userLogin, testApi };
