import { createRouter, createWebHashHistory, RouteRecordRaw } from "vue-router";

import HelloWorld from "@/components/HelloWorld.vue";
import DashboardLayout from "@/layout/DashboardLayout.vue";
const Login = { template: "<div>Login</div>" };

const RootRoute: RouteRecordRaw = {
  path: "/",
  name: "Root",
  component: DashboardLayout,
  meta: {
    title: "Root",
  },
  children: [
    {
      path: "/aaa/",
      component: HelloWorld,
    },
  ],
};

const LoginRoute: RouteRecordRaw = {
  path: "/login",
  name: "Login",
  component: Login,
  meta: {
    title: "",
  },
};

const router = createRouter({
  history: createWebHashHistory(),
  routes: [RootRoute, LoginRoute],
  // strict: true,
});

export { router };
