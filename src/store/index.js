// store/index.js
import Vue from 'vue';
import Vuex from 'vuex';
import axios from 'axios';
import { postLogin } from '@/api';



export default new Vuex.Store({
  state: {
    user: localStorage.getItem('user') != "undefined" ? JSON.parse(localStorage.getItem('user')) : null,
  },
  mutations: {
    SET_USER(state, user) {
      state.user = user;
      localStorage.setItem('user', JSON.stringify(user));
    },
    CLEAR_USER(state) {
      state.user = null;
      localStorage.removeItem('user');
    },
  },
  actions: {
    async login({ commit }, { username, password, success, error }) {
      try {
        const response = await postLogin(username,password);
        const user = response.data.userinfo;
        commit('SET_USER', user);
        success(response);
      } catch (err) {
        error();
      }
    },
    logout({ commit }) {
      commit('CLEAR_USER');
    },
  },
  getters: {
    isLoggedIn: state => !!state.user,
    user: state => state.user,
  },
});