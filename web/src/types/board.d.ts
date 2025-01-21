// src/types/board.d.ts

export interface Card {
  id: string;
  title: string;
  color: string;
}

export interface List {
  id: string;
  title: string;
  cards: Card[];
}
