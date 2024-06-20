<template>
  <Content>
    登录
    <form @submit.prevent="handleLogin">
      <input v-model="username" type="text" placeholder="用户名" required />
      <input v-model="password" type="password" placeholder="密码" required />
      <button type="submit">登录</button>
    </form>
  </Content>
</template>

<script>
import { mapActions } from 'vuex';
import Content from '../components/Content';

export default {
  name: 'Login',
  components: {
    Content,
  },
  data() {
    return {
      username: '',
      password: '',
    };
  },
  methods: {
    ...mapActions(['login']),
    handleLogin() {
      this.login({
        username: this.username,
        password: this.password,
        success: () => {
          // 登录成功后的其他操作，例如跳转到主页
          this.$router.push('/');
        },
        error: () => {
          // 处理登录失败逻辑
          console.error('登录失败');
          // 显示错误消息
          alert('登录失败，请检查用户名和密码');
        },
      });
    },
  },
};
</script>

<style scoped>
/* 你的样式 */
</style>