import { createApp } from "vue";
import { createPinia } from "pinia";
import { router } from "@/router";
import { i18n } from "@/locales/i18n.ts";
import { autoAnimatePlugin } from "@formkit/auto-animate/vue";

// import ContextMenu from "@imengyu/vue3-context-menu";
// import "@imengyu/vue3-context-menu/lib/vue3-context-menu.css";
// import "splitpanes/dist/splitpanes.css";
import "primeicons/primeicons.css";
import "vue-multiselect/dist/vue-multiselect.min.css";
import "@fortawesome/fontawesome-free/css/all.min.css"
import "./style.scss";

import { tooltip } from "@/directive/tooltip.ts";
import App from "./App.vue";

const pinia = createPinia();
const app = createApp(App);

app.use(i18n);
app.use(pinia);
app.use(router);
app.use(autoAnimatePlugin);
app.directive("tooltip", tooltip);
app.mount("#app");
