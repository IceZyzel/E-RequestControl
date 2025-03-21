import { createRouter, createWebHistory } from 'vue-router';
import Login from '../views/Login.vue';
import Registration from '../views/Registration.vue';
import RegisterAdmin from '../views/RegisterAdmin.vue';
import Dashboard from "../views/Dashboard.vue";
import AdminDashboard from "../views/AdminDashboard.vue";

const routes = [
    { path: '/', redirect: '/login' },
    { path: '/login', component: Login },
    { path: '/register', component: Registration },
    { path: '/registerAdmin', component: RegisterAdmin },
    { path: '/dashboard', component: Dashboard, meta: { requiresUser: true } },
    { path: '/admin-dashboard', component: AdminDashboard, meta: { requiresAdmin: true } },
];

const router = createRouter({
    history: createWebHistory(),
    routes,
});

router.beforeEach((to, from, next) => {
    const token = localStorage.getItem('token');
    const role = localStorage.getItem('role');

    if (!token && !['/login', '/register', '/registerAdmin'].includes(to.path)) {
        return next('/login');
    }

    if (to.meta.requiresAdmin && role !== '1') {
        return next('/dashboard');
    }

    if (to.meta.requiresUser && role === '1') {
        return next('/admin-dashboard');
    }

    next();
});

export default router;
