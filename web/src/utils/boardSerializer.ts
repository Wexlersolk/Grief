// src/utils/boardSerializer.ts
import { Board, List, Card } from '../types/board';

export const serializeBoard = (board: Board): Board => {
  // Ensure all data is included (e.g., lists, cards, colors, etc.)
  return {
    id: board.id,
    title: board.title,
    lists: board.lists.map((list: List) => ({
      id: list.id,
      title: list.title,
      cards: list.cards.map((card: Card) => ({
        id: card.id,
        title: card.title,
        description: card.description || '',
        color: card.color || '#ffffff', // Default color if not provided
      })),
    })),
  };
};
