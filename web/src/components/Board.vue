<template>
  <div class="flex space-x-4">
    <List v-for="(list, index) in lists" :key="index" :list="list" @add-card="addCard(index)" />
    <div class="w-64 bg-gray-200 rounded p-4 flex items-center justify-center">
      <button
        @click="addList"
        class="text-blue-600 font-bold hover:underline"
      >
        + Add List
      </button>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { reactive } from 'vue';
import List from './List.vue';

const lists = reactive([
  { title: 'To Do', cards: ['Task 1', 'Task 2'] },
  { title: 'In Progress', cards: ['Task 3'] },
]);

const addList = () => {
  lists.push({ title: `List ${lists.length + 1}`, cards: [] });
};

const addCard = (listIndex: number) => {
  const cardName = prompt('Enter card name:');
  if (cardName) lists[listIndex].cards.push(cardName);
};
</script>

