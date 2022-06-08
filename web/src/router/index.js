import { createRouter, createWebHistory } from "vue-router";
import HomeView from "../views/HomeView.vue";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      name: "home",
      component: HomeView,
    },
    {
      path: "/email",
      name: "email",
      component: () => import("../views/EmailView.vue"),
    },
    {
      path: "/register",
      name: "emailregister",
      component: () => import("../views/EmailRegister.vue"),
    },
    {
      path: "/sso",
      name: "sso",
      component: () => import("../views/SsoLogin.vue"),
    },
  ],
});

export default router;
