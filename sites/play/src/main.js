import Vue from 'vue';
import { BootstrapVue, IconsPlugin } from 'bootstrap-vue';
import axios from 'axios'
import VueAxios from 'vue-axios'
import VueBootstrapTypeahead from 'vue-bootstrap-typeahead'


import App from './App.vue';
import router from './router';
import store from './store';
import 'bootstrap/dist/css/bootstrap.css';
import 'bootstrap-vue/dist/bootstrap-vue.css';
// import 'bootswatch/dist/cosmo/bootstrap.min.css';

// Make BootstrapVue available throughout your project
Vue.use(BootstrapVue);
// Optionally install the BootstrapVue icon components plugin
Vue.use(IconsPlugin);

// Global registration
Vue.component('vue-bootstrap-typeahead', VueBootstrapTypeahead)

Vue.config.productionTip = false;

Vue.use(VueAxios, axios)

new Vue({
  router,
  store,
  render: (h) => h(App),
}).$mount('#app');
