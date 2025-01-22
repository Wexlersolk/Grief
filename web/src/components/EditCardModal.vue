<template>
  <div v-if="isOpen" class="fixed inset-0 flex items-center justify-center bg-black bg-opacity-50" @click.self="closeModal">
    <div class="bg-white p-5 rounded shadow-md w-96">
      <h2 class="text-lg font-bold mb-4">Edit Card</h2>
      <input v-model="card.title" placeholder="Card Title" class="border p-2 w-full rounded mb-4" />
      <input type="color" v-model="card.color" class="border p-1 w-full mb-4" />
      <div class="flex justify-end">
        <button @click="saveChanges" class="bg-blue-500 text-white rounded p-2">Save</button>
        <button @click="closeModal" class="bg-gray-300 rounded p-2 ml-2">Cancel</button>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { Card } from '../store/board';

export default defineComponent({
  props: {
    card: Object as () => Card,
    isOpen: Boolean,
  },
  emits: ['close', 'save'],
  setup(props, { emit }) {
    const closeModal = () => {
      emit('close');
    };

    const saveChanges = () => {
      emit('save', { ...props.card });
      closeModal();
    };

    return { closeModal, saveChanges };
  },
});
</script>

<style scoped>
/* Add any additional styling here */
</style>
