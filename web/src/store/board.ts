import { defineStore } from 'pinia';

export interface Card {
  id: string;
  title: string;
  color: string;
  description?: string;
}

export interface List {
  id: string;
  title: string;
  cards: Card[];
}

export const useBoardStore = defineStore('board', {
  state: () => ({
    lists: [] as List[],
  }),
  actions: {
    // Load data from localStorage
    loadFromLocalStorage() {
      const data = localStorage.getItem('kanban-board');
      if (data) {
        this.lists = JSON.parse(data);
      }
    },

    // Save data to localStorage
    saveToLocalStorage() {
      localStorage.setItem('kanban-board', JSON.stringify(this.lists));
    },

    // Add a new list
    addList(title: string) {
      const newList: List = {
        id: Date.now().toString(),
        title,
        cards: [],
      };
      this.lists.push(newList);
      this.saveToLocalStorage();
    },

    // Add a new card to a list
    addCard(listId: string, title: string = '') {
      const list = this.lists.find((list) => list.id === listId);
      if (list) {
        const lastCardColor = list.cards.length > 0 ? list.cards[list.cards.length - 1].color : '#000000';
        const newCard: Card = {
          id: Date.now().toString(),
          title,
          color: lastCardColor,
        };
        list.cards.push(newCard);
        this.saveToLocalStorage();
      }
    },

    // Remove a card from a list
    removeCard(listId: string, cardId: string) {
      const list = this.lists.find((list) => list.id === listId);
      if (list) {
        list.cards = list.cards.filter((card) => card.id !== cardId);
        this.saveToLocalStorage();
      }
    },

    // Edit a card's details
    editCard(listId: string, cardId: string, newTitle: string, newColor: string) {
      const list = this.lists.find((list) => list.id === listId);
      if (list) {
        const card = list.cards.find((card) => card.id === cardId);
        if (card) {
          card.title = newTitle;
          card.color = newColor;
          this.saveToLocalStorage();
        }
      }
    },

    // Move a card between lists
    moveCard(cardId: string, fromListId: string, toListId: string) {
      const fromList = this.lists.find((list) => list.id === fromListId);
      const toList = this.lists.find((list) => list.id === toListId);

      if (fromList && toList) {
        const cardIndex = fromList.cards.findIndex((card) => card.id === cardId);
        if (cardIndex !== -1) {
          const [card] = fromList.cards.splice(cardIndex, 1);
          toList.cards.push(card);
          this.saveToLocalStorage();
        }
      }
    },

    // Get the color of the last card in a list
    getLastCardColor(listId: string): string {
      const list = this.lists.find((list) => list.id === listId);
      if (list && list.cards.length > 0) {
        return list.cards[list.cards.length - 1].color;
      }
      return '#000000'; // Default color
    },
  },
});
