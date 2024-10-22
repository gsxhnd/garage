import { createRouter, createWebHashHistory, RouteRecordRaw } from "vue-router";
// import { usePreferencesStore } from "@/stores/preferences";

import Root from "@/layout/Root.vue";
// import InitElectron from "@/pages/InitElectron.vue";
import Anime from "@/pages/Anime.vue";
import Actor from "@/pages/Actor.vue";
import Movie from "@/pages/Movie.vue";
import Tag from "@/pages/Tag.vue";
import Setting from "@/pages/Setting.vue";

const RootRoute: RouteRecordRaw = {
  path: "/",
  name: "Root",
  component: Root,
  meta: {
    title: "Root",
  },
  children: [
    { path: "anime", name: "anime", component: Anime },
    { path: "actor", name: "actor", component: Actor },
    { path: "movie", name: "movie", component: Movie },
    { path: "tag", name: "tag", component: Tag },
    { path: "setting", name: "setting", component: Setting },
  ],
};

const router = createRouter({
  history: createWebHashHistory(),
  routes: [RootRoute],
  strict: true,
});

router.beforeEach(async (_to, _from) => {
  //   const preferencesStore = usePreferencesStore();
  //   console.log(preferencesStore.useElectron);
  // if (
  //   to.name != "InitElectron" &&
  //   !preferencesStore.useBrowser &&
  //   (preferencesStore.preference === null ||
  //     preferencesStore.preference.appConfig.libraries.length === 0)
  // ) {
  //   preferencesStore.preference?.appConfig.libraries.forEach((library) => {
  //     const { path, use } = library;
  //     console.log(path, use);
  //   });
  //   return { name: "InitElectron" };
  // }
});

export { router };
