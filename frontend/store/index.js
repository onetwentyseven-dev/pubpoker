export const state = () => ({
  atTop: true,
})

export const getters = {
  getTop() {
    return state.atTop
  },
}

export const actions = {
  setAtTop({ commit }, status) {
    commit('storeAtTop', status)
  },
}

export const mutations = {
  storeAtTop(state, value) {
    state.atTop = value
  },
}
