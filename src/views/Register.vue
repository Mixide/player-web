<template>
  <Content>
    用户名：<input type="text" v-model="username"> <br>
    密码：<input type="pasword" v-model="password"> <br>
    <button @click="submit">注册</button> <br>
    <router-link to="/login">登录</router-link>
  </Content>
</template>

<script>
  import Content from '../components/Content';
  import axios from 'axios';
  export default {
    name: 'Register',
    components:{
      Content
    },
    data() {
      return {
        username: "",
        password: ""
      }
    },
    methods: {
      submit() {
        axios.post('http://localhost:8000/api/user/register/', {"username": this.username, "password": this.password})
        .then((response) => {
            alert(response.data.msg)
        })
      }
    }
  }
</script>

<style scoped>

</style>