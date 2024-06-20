import Vue from 'vue'
import axios from 'axios'

export const getMusic = (id) => {return axios.get('http://localhost:8000/api/music/',{
    params: {
      userid: id
    }
  });
}

export const postMusic = (formData,id) => {
  formData.append('userid', id);
  return axios.post('http://localhost:8000/api/upload/', formData, {
    headers: {
      'Content-Type': 'multipart/form-data'
    }
  });
}

export const postLogin = (username,password) => {
  return axios.post('http://localhost:8000/api/user/login/', {
    username,
    password,
  });
}

export const postRegister = (username,password) => {
  return axios.post('http://localhost:8000/api/user/register/', {
    "username": username, 
    "password": password
  });   
}
