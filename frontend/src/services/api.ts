import axios from 'axios';

const api = axios.create({
  baseURL: 'http://localhost:5000/', 
});

// Exemplo de funções CRUD:
export const getUsers = () => api.get('/users');
export const createUser = (data: any) => api.post('/user', data);
export const updateUser = (id: number, data: any) => api.put(`/user/${id}`, data);
export const deleteUser = (id: number) => api.delete(`/user/${id}`);
export const login = (data: any) => api.post('/login', data);

export default api;
