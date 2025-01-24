// src/router/index.ts

import { createRouter, createWebHistory } from 'vue-router';
import Home from '@/pages/Home.vue';
import TicTacToe from '@/pages/TicTacToe.vue';
import Canvas from '@/pages/Canvas.vue';
import Board from '@/pages/Board.vue'; // Import the Board page
import Login from '@/pages/Login.vue'; // Import the Login page
import Signin from '@/pages/Signup.vue'; // Import the Signin page

const routes = [
  { path: '/', name: 'Home', component: Home },
  { path: '/tictactoe', name: 'TicTacToe', component: TicTacToe },
  { path: '/canvas', name: 'Canvas', component: Canvas },
  { path: '/board', name: 'Board', component: Board }, // Route for Board
  { path: '/login', name: 'Login', component: Login }, // Route for Login
  { path: '/signup', name: 'Signup', component: Signin }, // Route for Signin
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
