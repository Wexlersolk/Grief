<template>
  <div
    class="bg-white rounded-lg shadow-md cursor-pointer hover:shadow-lg transition-shadow"
    draggable="true"
    @dragstart="onDragStart"
    @dragend="onDragEnd"
    @click="openEditModal"
  >
    <!-- Color strip at the top -->
    <div :style="{ backgroundColor: card.color }" class="h-2 rounded-t-lg"></div>

    <!-- Card content -->
    <div class="p-3">
      <p class="text-sm">{{ card.title }}</p>
    </div>

    <!-- Remove button -->
    <button
      @click.stop="emitRemoveCard"
      class="absolute top-1 right-1 text-gray-500 hover:text-red-500"
    >
      ✕
    </button>
  </div>

  <!-- Edit Modal -->
  <EditCardModal
    v-if="isEditing"
    :card="card"
    :isOpen="isEditing"
    @close="closeEditModal"
    @save="saveChanges"
  />
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue';
import { Card as CardType } from '../store/board';
import EditCardModal from './EditCardModal.vue';

export default defineComponent({
  components: {
    EditCardModal,
  },
  props: {
    card: Object as () => CardType,
    listId: String,
  },
  emits: ['remove-card', 'edit-card'],
  setup(props, { emit }) {
    const isEditing = ref(false);

    const openEditModal = () => {
      isEditing.value = true;
    };

    const closeEditModal = () => {
      isEditing.value = false;
    };

    const saveChanges = (updatedCard: CardType) => {
      emit('edit-card', props.listId, props.card.id, updatedCard.title, updatedCard.color);
      closeEditModal();
    };

    const emitRemoveCard = () => {
      emit('remove-card', props.listId, props.card.id);
    };

    const onDragStart = (event: DragEvent) => {
      if (event.dataTransfer) {
        event.dataTransfer.setData('cardId', props.card.id);
        event.dataTransfer.setData('fromListId', props.listId!);
      }
    };

    const onDragEnd = () => {
      // Handle drag end logic if needed
    };

    return {
      isEditing,
      openEditModal,
      closeEditModal,
      saveChanges,
      emitRemoveCard,
      onDragStart,
      onDragEnd,
    };
  },
});
</script>

<style scoped>
/* Add any additional styling here */
</style>
