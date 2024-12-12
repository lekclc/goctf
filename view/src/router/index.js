import { createRouter, createWebHistory } from 'vue-router';
import Home from '../views/Home.vue';
import About from '../views/About.vue';
import User from '../views/User.vue';
import Login from '../views/Login.vue'; 
import Register from '../views/Register.vue';
import Team from '../views/Team.vue';
import Game from '../views/Game.vue';
import GameList from '../views/GameList.vue';
import Rank from '../views/Rank.vue';
const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/about',
    name: 'About',
    component: About
  },
  {
    path: '/user',
    name: 'User',
    component: User
  },
  {
    path: '/login',
    name: 'Login',
    component: Login
  },
  {
    path: '/register',
    name: 'Register',
    component: Register
  },
  {
    path: '/team',
    name: 'Team',
    component: Team
  },
  {
    path: '/game/:id',
    name: 'Game',
    component: Game
  },
  {
    path: '/game/:id&:teamid&:teamname',
    name: 'Game',
    component: Game
  },
  {
    path: '/gamelist',
    name: 'GameList',
    component: GameList
  },
  {
    path: '/rank/:id&:page',
    name: 'Rank',
    component: Rank
  }
  
];

const router = createRouter({
  history: createWebHistory(),
  routes
});

export default router;