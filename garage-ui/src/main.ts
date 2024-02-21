import { createApp } from "vue";
import { router } from "@/router";
import { createPinia } from "pinia";
import "./style.less";
import App from "./App.vue";

import "vuetify/styles";
import { createVuetify } from "vuetify";
import { aliases, mdi } from "vuetify/iconsets/mdi";
import "@mdi/font/css/materialdesignicons.css";

const vuetify = createVuetify({
  icons: {
    defaultSet: "mdi",
    aliases,
    sets: { mdi },
  },
});

const pinia = createPinia();

const app = createApp(App);
app.use(vuetify);
app.use(router);
app.use(pinia);

app.mount("#app");
