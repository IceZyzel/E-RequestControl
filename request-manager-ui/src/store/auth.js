import { defineStore } from 'pinia';
import { authApi } from '../api';
import { jwtDecode } from 'jwt-decode';

export const useAuthStore = defineStore('auth', {
    state: () => ({
        user: null,
        token: localStorage.getItem('token') || '',
        role: localStorage.getItem('role') || null,
    }),
    actions: {
        async login(credentials) {
            try {
                const res = await authApi.login(credentials.username, credentials.password);

                this.user = res.data.user;
                this.token = res.data.token;

                const decodedToken = jwtDecode(this.token);
                this.role = decodedToken.RoleID;

                localStorage.setItem('token', this.token);
                localStorage.setItem('role', this.role);

                return true;
            } catch (error) {
                console.error("Помилка авторизації:", error);
                alert('Невірний логін або пароль!');
                return false;
            }
        },
        logout() {
            this.user = null;
            this.token = '';
            this.role = null;
            localStorage.removeItem('token');
            localStorage.removeItem('role');
        },
        setRole(role) {
            this.role = role;
            localStorage.setItem('role', role);
        }
    },
    getters: {
        isAuthenticated: (state) => !!state.token,
        userRole: (state) => state.role,
    }
});
