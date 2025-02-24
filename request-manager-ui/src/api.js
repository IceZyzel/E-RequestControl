import axios from 'axios';

const apiClient = axios.create({
    baseURL: 'http://localhost:8000',
    headers: {
        'Content-Type': 'application/json',
    },
    withCredentials: true,
});

apiClient.interceptors.request.use(config => {
    const token = localStorage.getItem('token');
    console.log('Intercepting request, token:', token);
    if (token) {
        config.headers.Authorization = `Bearer ${token}`;
    }
    console.log('Config headers:', config.headers);
    return config;
}, error => {
    return Promise.reject(error);
});

apiClient.interceptors.response.use(response => {
    console.log('Response headers:', response.headers);
    return response;
}, error => {
    if (error.response) {
        console.error('Response error:', error.response);
    } else {
        console.error('Error without response:', error);
    }
    return Promise.reject(error);
});

export const authApi = {
    login: (username, password) => apiClient.post('/auth/login', { username, password }),
    register: (firstname,lastname,email,username,password ) => apiClient.post('/auth/register', { firstname,lastname,email,username,password }),
    registerAdmin: (firstname,lastname,email,username,password ) => apiClient.post('/auth/registerAdmin', { firstname,lastname,email,username,password }),
};


