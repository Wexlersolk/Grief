<template>
  <div class="bg-gray-100 rounded-lg p-3 w-72">
    <h2 class="font-bold text-lg mb-3">{{ list.title }}</h2>
    <div
      class="space-y-2"
      @dragover.prevent
      @drop="onDrop"
    >
      <Card
        v-for="card in list.cards"
        :key="card.id"
        :card="card"
        :listId="list.id"
        @remove-card="emitRemoveCard"
        @edit-card="emitEditCard"
      />
      <NewCardButton :listId="list.id" />
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import Card from './Card.vue';
import NewCardButton from './NewCardButton.vue';
import { Card as CardType, List as ListType } from '../store/board';

export default defineComponent({
  components: { Card, NewCardButton },
  props: {
    list: Object as () => ListType,
  },
  emits: ['remove-card', 'edit-card', 'move-card'],
  setup(props, { emit }) {
    const emitRemoveCard = (listId: string, cardId: string) => {
      emit('remove-card', listId, cardId);
    };

    const emitEditCard = (listId: string, cardId: string, title: string, color: string) => {
      emit('edit-card', listId, cardId, title, color);
    };

    const onDrop = (event: DragEvent) => {
      if (event.dataTransfer) {
        const cardId = event.dataTransfer.getData('cardId');
        const fromListId = event.dataTransfer.getData('fromListId');
        emit('move-card', cardId, fromListId, props.list.id);
      }
    };

    return { emitRemoveCard, emitEditCard, onDrop };
  },
});
</script>

<style scoped>
/* Add any additional styling here */
</style>
