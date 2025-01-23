<template>
  <div class="min-h-screen flex items-center justify-center bg-gray-100">
    <div class="bg-white p-8 rounded-lg shadow-md w-96">
      <h2 class="text-2xl font-bold mb-6 text-center">Login</h2>

      <!-- Email/Password Login Form -->
      <form @submit.prevent="login">
        <div class="mb-4">
          <label for="email" class="block text-sm font-medium text-gray-700">Email</label>
          <input
            v-model="email"
            type="email"
            id="email"
            class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
            required
          />
        </div>
        <div class="mb-6">
          <label for="password" class="block text-sm font-medium text-gray-700">Password</label>
          <input
            v-model="password"
            type="password"
            id="password"
            class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
            required
          />
        </div>
        <button
          type="submit"
          class="w-full bg-blue-500 text-white py-2 px-4 rounded-md hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-blue-500"
        >
          Login
        </button>
      </form>

      <!-- Signup Link -->
      <p class="mt-6 text-center text-sm text-gray-600">
        Don't have an account?
        <router-link to="/signin" class="text-blue-500 hover:text-blue-600">Sign up</router-link>
      </p>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue';
import { useRouter } from 'vue-router';
import { useAuthStore } from '@/store/auth'; // Import the mock auth store

export default defineComponent({
  setup() {
    const router = useRouter();
    const authStore = useAuthStore();
    const email = ref('');
    const password = ref('');

    const login = () => {
      // Mock login logic
      if (email.value === 'test@example.com' && password.value === 'password') {
        authStore.setUser({ email: email.value }); // Set user in the store
        router.push('/'); // Redirect to the home page
      } else {
        alert('Invalid email or password');
      }
    };

    return { email, password, login };
  },
});
</script>

<style scoped>
/* Add any additional styling here */
</style>
