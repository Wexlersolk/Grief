import { defineStore } from 'pinia';
import { loadBoardState, saveBoardState } from '../utils/localStorage';
//import { Card, List } from "../types/board";

export const useBoardStore = defineStore('board', {
  state: () => ({
    lists: loadBoardState() || [], // Load initial state from localStorage
  }),
  actions: {
    updateCard(updatedCard: Card) {
      const list = this.lists.find(l => l.cards.some(c => c.id === updatedCard.id));
      if (list) {
        const cardIndex = list.cards.findIndex(c => c.id === updatedCard.id);
        if (cardIndex !== -1) {
          list.cards[cardIndex] = updatedCard;
          saveBoardState(this.lists); // Persist changes
        }
      }
    },
    // Add more actions for adding/removing lists and cards
  }
});
