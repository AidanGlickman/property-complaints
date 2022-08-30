import { createApp } from "vue";
import Oruga from '@oruga-ui/oruga-next';
import '@oruga-ui/oruga-next/dist/oruga-full-vars.css';
import { createPinia } from "pinia";

import App from "./App.vue";
import router from "./router";

import { library } from '@fortawesome/fontawesome-svg-core'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';

import { faSearch } from '@fortawesome/free-solid-svg-icons'

library.add(faSearch);

import "./assets/main.css";

const app = createApp(App);

app.component('VueFontawesome', FontAwesomeIcon);

app.use(Oruga, {
    iconComponent: 'vue-fontawesome',
    iconPack: 'fas',
});
app.use(createPinia());
app.use(router);

app.mount("#app");
