import {createStore} from 'vuex'
import {api,wsUrl} from "../api"


export const store = createStore({
  state () {
    return {
      api: api,
      wsUrl:wsUrl
    }
  },
  mutations: {
    setApi (state, api) {
      state.api = api;
    },
    setUrl (state, wsUrl) {
      state.wsUrl = wsUrl;
    },
  },

})
