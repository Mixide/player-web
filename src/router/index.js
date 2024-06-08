import { createRouter, createWebHistory } from 'vue-router'
import Home from '../views/Home.vue';
import Login from '../views/Login.vue';
import MusicSpace from '../views/MusicSpace.vue';
import Register from '../views/Register.vue';
import NotFound from '../views/NotFound.vue';

const routes = [
  {
    path: '/',
    name: 'home',
    component:Home
  },
  {
    path: '/musicspace',
    name: 'musicspace',
    component:MusicSpace
  },
  {
    path: '/login',
    name: 'login',
    component:Login
  },
  {
    path: '/register',
    name: 'register',
    component:Register
  },
  {
    path: '/404',
    name: '404',
    component:NotFound
  },
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router
