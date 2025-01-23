

import { defineStore } from 'pinia';

// Define the shape of the user object
interface User {
  email: string;
}

// Define the auth store
export const useAuthStore = defineStore('auth', {
  // State
  state: () => ({
    user: null as User | null, // Holds the logged-in user or null if no user is logged in
  }),

  // Getters
  getters: {
    isAuthenticated: (state) => !!state.user, // Returns true if a user is logged in
  },

  // Actions
  actions: {
    // Set the logged-in user
    setUser(user: User) {
      this.user = user;
    },

    // Clear the logged-in user (logout)
    clearUser() {
      this.user = null;
    },
  },
});
