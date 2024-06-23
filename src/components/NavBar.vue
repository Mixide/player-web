<template>
<nav class="navbar navbar-expand-lg bg-body-tertiary">
  <div class="container-fluid">
    <a class="navbar-brand" href="#">在线音乐</a>
    <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarText" aria-controls="navbarText" aria-expanded="false" aria-label="Toggle navigation">
      <img src="../assets/audio256.png" alt="Logo" width="35" height="35">
      <span class="navbar-toggler-icon"></span>
    </button>
    <div class="collapse navbar-collapse" id="navbarText">
      <ul class="navbar-nav me-auto mb-2 mb-lg-0">
        <li class="nav-item">
          <router-link class="nav-link" :to="{name:'musicspace'}">音乐空间</router-link>
        </li>
      </ul>
      <ul class="navbar-nav" v-if="!store.getters.isLoggedIn">
        <li class="nav-item">
          <router-link class="nav-link" :to="{name:'login'}">登录</router-link>
        </li>
        <li class="nav-item">
          <router-link class="nav-link" :to="{name:'register'}">注册</router-link>
        </li>
      </ul>
      <ul class="navbar-nav" v-else>
        <li class="nav-item">
          <a class="nav-link" style="cursor: pointer" @click="handleLogout">退出</a>
        </li>
      </ul>
    </div>
  </div>
</nav>
</template>

<script>
import { useStore } from 'vuex';
import { mapActions } from 'vuex';
export default{
  name:'NavBar',
  setup() {
    const store = useStore();
    return {store};
  },
  methods: {
    ...mapActions(['logout']),
    handleLogout() {
      this.logout();
      this.$router.push('/login/');
    },
  },
}
</script>
<style scoped>

</style>