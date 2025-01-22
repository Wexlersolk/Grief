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
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { useBoardStore } from '../store/board';
import List from '../components/List.vue';

export default defineComponent({
  components: { List },
  setup() {
    const boardStore = useBoardStore();
    boardStore.loadFromLocalStorage();

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

    return { boardStore, addNewList, handleRemoveCard, handleEditCard, handleMoveCard };
  },
});
</script>

<style scoped>
/* Add any additional styling here */
</style>
