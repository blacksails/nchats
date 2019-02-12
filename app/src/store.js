import Vue from "vue";
import Vuex from "vuex";

Vue.use(Vuex);

let conn;

export default new Vuex.Store({
  state: {
    nickname: "",
    messages: [],
    inputMessage: ""
  },

  mutations: {
    setNickname(state, nickname) {
      state.nickname = nickname;
    },
    connectionClosed(state) {
      state.messages = [...state.messages, { message: "Connection closed." }];
    },
    newMessage(state, msg) {
      state.messages = [...state.messages, msg];
    },
    clearMessages(state) {
      state.messages = [];
    },
    setInputMessage(state, str) {
      state.inputMessage = str;
    }
  },

  actions: {
    connect({ commit }, nickname) {
      conn = new WebSocket("ws://" + document.location.host + "/ws");
      conn.onclose = () => {
        commit("connectionClosed");
        setTimeout(() => commit("clearMessages"));
      };
      conn.onmessage = event => {
        commit("newMessage", JSON.parse(event.data));
      };
      commit("setNickname", nickname);
    },
    sendMessage({ commit }, message) {
      conn.send(JSON.stringify(message));
      commit("setInputMessage", "");
    },
    disconnect({ commit }) {
      conn.close();
      commit("setNickname", "");
    }
  }
});
