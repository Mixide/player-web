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
    UPDATE_USER_PHOTO(state, photoUrl) {
      state.user.photo = photoUrl;
      localStorage.setItem('user', JSON.stringify(state.user));
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
    updateUserPhoto({ commit }, photoUrl) {
      commit('UPDATE_USER_PHOTO', photoUrl);
    },
  },
  getters: {
    isLoggedIn: state => !!state.user,
    user: state => state.user,
  },
});