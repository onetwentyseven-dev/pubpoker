import Vue from 'vue';
import Vuex from 'vuex';

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    atTop: true
  },
  getters: {
    getTop() {
      return state.atTop;
    }
  },
  mutations: {
    storeAtTop(state, value) {
      state.atTop = value
    }
  },
  actions: {
    setAtTop({ commit }, status) {
      commit("storeAtTop", status)
    }
  },
  modules: {
  },
});
