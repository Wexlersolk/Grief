// src/types/board.ts
export interface Card {
  id: string;
  title: string;
  description?: string;
  color?: string;
}

export interface List {
  id: string;
  title: string;
  cards: Card[];
}

export interface Board {
  id: string;
  title: string;
  lists: List[];
}
