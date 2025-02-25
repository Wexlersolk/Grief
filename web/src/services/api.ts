// src/services/api.ts
import { Board } from '../types/board';

const API_BASE_URL = 'https://your-server-url.com/api'; // Replace with your server URL

export const saveBoard = async (board: Board) => {
  try {
    const response = await fetch(`${API_BASE_URL}/boards`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(board),
    });

    if (!response.ok) {
      throw new Error('Failed to save board');
    }

    return await response.json();
  } catch (error) {
    console.error('Error saving board:', error);
    throw error;
  }
};
