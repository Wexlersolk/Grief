import { createRouter, createWebHistory } from 'vue-router';
import { getKeycloakInstance } from '@/services/keycloak'; // Import Keycloak instance
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

// Define public routes that don't require authentication
const publicRoutes = ['Login', 'Signup'];

// Add a global navigation guard
router.beforeEach((to, from, next) => {
  const keycloak = getKeycloakInstance();

  // Check if the route is public
  if (publicRoutes.includes(to.name as string)) {
    next(); // Allow access to public routes
    return;
  }

  // Check if the user is authenticated
  if (keycloak.authenticated) {
    next(); // Allow access to protected routes
  } else {
    // Redirect to Keycloak login if not authenticated
    keycloak.login({ redirectUri: window.location.origin + to.fullPath });
  }
});

export default router;
