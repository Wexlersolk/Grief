export const saveBoardState = (lists) => {
  localStorage.setItem('boardState', JSON.stringify(lists));
};

export const loadBoardState = () => {
  const data = localStorage.getItem('boardState');
  return data ? JSON.parse(data) : [];
};
