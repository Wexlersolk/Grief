<template>
  <button @click="addCard" class="mt-2 bg-green-500 text-white rounded p-2">+ Create New Card</button>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { useBoardStore } from '../store/board';

export default defineComponent({
  props: {
    listId: String,
  },
  setup(props) {
    const boardStore = useBoardStore();

    const addCard = () => {
      const newCard = {
        id: Date.now().toString(),
        title: '',
        color: boardStore.getLastCardColor(props.listId!),
      };
      boardStore.addCard(props.listId!, newCard.title);
    };

    return { addCard };
  },
});
</script>

<style scoped>
/* Add any additional styling here */
</style>
