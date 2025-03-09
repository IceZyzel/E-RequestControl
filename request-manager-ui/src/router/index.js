import { createRouter, createWebHistory } from 'vue-router';
import Login from '../views/Login.vue';
import Registration from '../views/Registration.vue';
import RegisterAdmin from '../views/RegisterAdmin.vue';
import Dashboard from "../views/Dashboard.vue";

const routes = [
    { path: '/', redirect: '/login' },
    { path: '/login', component: Login },
    { path:'/register', component: Registration },
    { path:'/registerAdmin', component: RegisterAdmin },
    { path:'/dashboard', component: Dashboard },
];

const router = createRouter({
    history: createWebHistory(),
    routes
});

router.beforeEach((to, from, next) => {
    const isAuthenticated = localStorage.getItem('token');
    const isAdmin = localStorage.getItem('role') === 'admin';


    if (to.meta.requiresAuth && !isAuthenticated) {
        next('/login');
    } else if (to.meta.requiresAdmin && !isAdmin) {
        next('/dashboard');
    } else {
        next();
    }
});

export default router;
