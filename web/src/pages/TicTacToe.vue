<script lang="ts" setup>
import { ref, computed } from 'vue';

const player = ref<'X' | 'O'>('X');
const board = ref<string[][]>([
  ['', '', ''],
  ['', '', ''],
  ['', '', '']
]);

const CalculateWinner = (board: string[][]): string | null => {
  const lines = [
    [0, 1, 2], [3, 4, 5], [6, 7, 8],
    [0, 3, 6], [1, 4, 7], [2, 5, 8],
    [0, 4, 8], [2, 4, 6]
  ];

  const flatBoard = board.flat();

  for (const [a, b, c] of lines) {
    if (flatBoard[a] && flatBoard[a] === flatBoard[b] && flatBoard[a] === flatBoard[c]) {
      return flatBoard[a];
    }
  }

  return null;
};

const winner = computed(() => CalculateWinner(board.value));

const MakeMove = (x: number, y: number): void => {
  if (winner.value || board.value[x][y]) return;

  board.value[x][y] = player.value;
  player.value = player.value === 'X' ? 'O' : 'X';
};

const ResetGame = (): void => {
  board.value = [
    ['', '', ''],
    ['', '', ''],
    ['', '', '']
  ];
  player.value = 'X';
};
</script>

<template>
  <main class="pt-8 text-center">
    <h1 class="mb-8 text-3xl font-bold uppercase">Tic Tac Toe</h1>
    <h3 class="text-xl mb-4">Player {{ player }}'s turn</h3>

    <div class="flex flex-col items-center mb-8">
      <div v-for="(row, x) in board" :key="x" class="flex">
        <div 
          v-for="(cell, y) in row" 
          :key="y" 
          @click="MakeMove(x, y)" 
          class="border border-white w-24 h-24 hover:bg-gray-700 flex items-center justify-center material-icons-outlined text-4xl cursor-pointer">
          {{ cell }}
        </div>
      </div>
    </div>

    <div>
      <h2 v-if="winner" class="text-6xl font-bold mb-8">Player '{{ winner }}' wins!</h2>
      <button @click="ResetGame" class="px-4 py-2 bg-pink-500 rounded uppercase font-bold">Reset</button>
    </div>
  </main>
</template>
