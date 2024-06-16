import Vue from 'vue'
import axios from 'axios'

export const getMusic = () => {return axios.get('http://localhost:8000/api/music/')}

export const postMusic = (formData) => {return axios.post('http://localhost:8000/api/upload/', formData, {
    headers: {
      'Content-Type': 'multipart/form-data'
    }
  })}