<template>
  <div class="flex space-x-4 p-4 overflow-x-auto">
    <List
      v-for="list in boardStore.lists"
      :key="list.id"
      :list="list"
      @remove-card="handleRemoveCard"
      @edit-card="handleEditCard"
      @move-card="handleMoveCard"
    />
    <button @click="addNewList" class="bg-blue-500 text-white rounded p-2">Add New List</button>
    <button @click="saveBoardToServer" class="bg-green-500 text-white rounded p-2 ml-2">Save Board</button>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { useBoardStore } from '../store/board';
import List from '../components/List.vue';
import { serializeBoard } from '../utils/boardSerializer';
import { saveBoard } from '../services/api';

export default defineComponent({
  components: { List },
  setup() {
    const boardStore = useBoardStore();
    boardStore.loadFromLocalStorage(); // Optional: Keep this if you want to load initial data from local storage

    const addNewList = () => {
      const title = prompt('Enter list title');
      if (title) {
        boardStore.addList(title);
      }
    };

    const handleRemoveCard = (listId: string, cardId: string) => {
      boardStore.removeCard(listId, cardId);
    };

    const handleEditCard = (listId: string, cardId: string, title: string, color: string) => {
      boardStore.editCard(listId, cardId, title, color);
    };

    const handleMoveCard = (cardId: string, fromListId: string, toListId: string) => {
      boardStore.moveCard(cardId, fromListId, toListId);
    };

    const saveBoardToServer = async () => {
      // Serialize the board data
      const serializedBoard = serializeBoard({
        id: 'board-1', // Replace with a dynamic ID if needed
        title: 'My Board', // Replace with a dynamic title if needed
        lists: boardStore.lists,
      });

      // Send the serialized board data to the server
      try {
        await saveBoard(serializedBoard);
        alert('Board saved successfully!');
      } catch (error) {
        console.error('Error saving board:', error);
        alert('Failed to save board.');
      }
    };

    return {
      boardStore,
      addNewList,
      handleRemoveCard,
      handleEditCard,
      handleMoveCard,
      saveBoardToServer,
    };
  },
});
</script>

<style scoped>
/* Add any additional styling here */
</style>
