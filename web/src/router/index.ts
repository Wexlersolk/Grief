import { createRouter, createWebHistory } from 'vue-router';
import Home from '@/pages/Home.vue';
import TicTacToe from '@/pages/TicTacToe.vue';
import Canvas from '@/pages/Canvas.vue';

const routes = [
  { path: '/', name: 'Home', component: Home },
  { path: '/tictactoe', name: 'TicTacToe', component: TicTacToe },
  { path: '/canvas', name: 'Canvas', component: Canvas },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;

