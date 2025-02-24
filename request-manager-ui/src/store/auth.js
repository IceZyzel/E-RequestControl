import { defineStore } from 'pinia';
import axios from 'axios';


export const useAuthStore = defineStore('auth', {
    state: () => ({
        user: null,
        token: localStorage.getItem('token') || ''
    }),
    actions: {
        async login(credentials) {
            try {
                const res = await axios.post('/auth/login', credentials);
                this.user = res.data.user;
                this.token = res.data.token;
                localStorage.setItem('token', res.data.token);
                localStorage.setItem('role', res.data.user.role);
                axios.defaults.headers.common['Authorization'] = `Bearer ${res.data.token}`;
            } catch (error) {
                alert('Помилка авторизації!');
            }
        },
        logout() {
            this.user = null;
            this.token = '';
            localStorage.removeItem('token');
            localStorage.removeItem('role');
            delete axios.defaults.headers.common['Authorization'];
        },
        setUser(user) {
            this.user = user;
        }
    }
});
