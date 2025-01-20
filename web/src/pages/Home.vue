<template>
  <main class="flex flex-col items-center justify-center h-screen">
    <h1 class="text-4xl font-bold mb-8">Welcome to the Game Hub</h1>
    <div class="flex gap-4">
      <router-link
        to="/tictactoe"
        class="px-6 py-3 bg-pink-500 rounded uppercase font-bold hover:bg-pink-600 duration-300">
        Tic Tac Toe
      </router-link>
      <router-link
        to="/canvas"
        class="px-6 py-3 bg-blue-500 rounded uppercase font-bold hover:bg-blue-600 duration-300">
        Canvas
      </router-link>
    </div>
  </main>
</template>

<script lang="ts" setup>
</script>

