// src/router/index.ts

import { createRouter, createWebHistory } from 'vue-router';
import Home from '@/pages/Home.vue';
import TicTacToe from '@/pages/TicTacToe.vue';
import Canvas from '@/pages/Canvas.vue';
import Board from '@/pages/Board.vue'; // Import the Board page

const routes = [
  { path: '/', name: 'Home', component: Home },
  { path: '/tictactoe', name: 'TicTacToe', component: TicTacToe },
  { path: '/canvas', name: 'Canvas', component: Canvas },
  { path: '/board', name: 'Board', component: Board }, // Add route for Board
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
