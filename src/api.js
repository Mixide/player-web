import Vue from 'vue'
import axios from 'axios'

export const getMusic = () => {return axios.get('http://localhost:8000/api/music/')}

export const postMusic = (formData) => {return axios.post('http://localhost:8000/api/upload/', formData, {
  headers: {
    'Content-Type': 'multipart/form-data'
  }
})}

export const getToken = (data) => {return axios.post("https://localhost:8000/api/token/", {
  username: data.username,
  password: data.password,
})}

export const postRefresh = (refresh) => {return axios.post("https://localhost:8000/api/token/refresh/", {
  refresh: refresh,
})}

export const getInfo = (user_id,access) => {return axios.get("https://app165.acapp.acwing.com.cn/myspace/getinfo/", {
  params: {
    user_id: user_id,
  },
  headers: {
    'Authorization': "Bearer " + access,
  }
})}