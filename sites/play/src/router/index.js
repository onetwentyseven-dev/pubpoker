import Vue from 'vue';
import VueRouter from 'vue-router';

Vue.use(VueRouter);

const routes = [
  {
    path: '/',
    name: 'home',
    component: () => import(/* webpackChunkName: "index" */ '../views/HomeView.vue'),
  },
  {
    path: "/tournaments/:tournamentID",
    name: "tournaments",
    component: () => import(/* webpackChunkName: "tournaments" */ "../views/TournamentView.vue"),
  },
  {
    path: "/tournaments/:tournamentID/players",
    name: "players",
    component: () => import(/* webpackChunkName: "tournaments" */ "../views/TournamentPlayersView.vue"),
  }
];

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes,
});

export default router;
