import ky from "ky";

const httpv2 = ky.create({
  prefixUrl: "https://example.com/api",
  timeout: 5000,
  headers: {
    "Content-Type": "application/json",
  },
  retry: 1,
  hooks: {
    beforeRequest: [
      (options) => {
        console.log("before request", options);
      },
    ],
    beforeRetry: [
      async ({ error }) => {
        console.log("before retry", error);
      },
    ],
    beforeError: [
      (error) => {
        console.log("before error:", error);

        return error;
      },
    ],
    afterResponse: [
      (request, _options, response) => {
        console.log("after reponse:", response.status);
        console.log("after reponse:", request.headers);
        return new Response("A different response", { status: 200 });
      },
    ],
  },
  throwHttpErrors: false,
});

export { httpv2 };
