import { defineStore } from "pinia";

export const useLayoutStore = defineStore("layout", {
  state: () => {
    return {
      app_bar: {
        hover: true,
        rail: true,
      },
    };
  },
});
