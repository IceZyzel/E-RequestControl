import axios from "axios";

const apiClient = axios.create({
    baseURL: "http://localhost:8000",
    headers: {
        "Content-Type": "application/json",
    },
    withCredentials: true,
});

apiClient.interceptors.request.use(
    (config) => {
        const token = localStorage.getItem("token");

        if (token) {
            config.headers['Authorization'] = `Bearer ${token}`;
        }

        return config;
    },
    (error) => Promise.reject(error)
);

apiClient.interceptors.response.use(
    (response) => {
        return response;
    },
    (error) => {
        return Promise.reject(error);
    }
);

export const authApi = {
    login: (username, password) =>
        apiClient.post("/auth/login", { username, password }),
    register: (firstname, lastname, email, username, password) =>
        apiClient.post("/auth/register", {
            firstname,
            lastname,
            email,
            username,
            password,
        }),
    registerAdmin: (firstname, lastname, email, username, password) =>
        apiClient.post("/auth/registerAdmin", {
            firstname,
            lastname,
            email,
            username,
            password,
        }),
};

export const requestApi = {
    getUserTickets: () => apiClient.get("/api/tickets/"),
    getAllUsers: () => apiClient.get("/api/users"),
    createTicket: async (ticketData) => {
        try {
            const response = await apiClient.post("/api/tickets/", ticketData);
            return response.data;
        } catch (error) {
            console.error("Помилка при створенні тікета:", error);
            throw error;
        }
    },
    updateTicket: (ticketID, ticketData) =>
        apiClient.put(`/api/tickets/${ticketID}`, ticketData),
    deleteTicket: async (ticketID) => {
        try {
            await apiClient.delete(`/api/tickets/${ticketID}`);
            return ticketID;
        } catch (error) {
            console.error("Помилка при видаленні тікета:", error);
            throw error;
        }
    },
};

export const notificationApi = {
    getUserNotifications: () => apiClient.get("/api/notifications/"),
    createNotification: async (notificationData) => {
        try {
            const response = await apiClient.post(
                "/api/notifications/",
                notificationData
            );
            return response.data;
        } catch (error) {
            console.error("Помилка при створенні сповіщення:", error);
            throw error;
        }
    },
    markAsRead: async (notificationID) => {
        try {
            await apiClient.delete(`/api/notifications/${notificationID}`);
            return notificationID;
        } catch (error) {
            console.error("Помилка при видаленні сповіщення:", error);
            throw error;
        }
    },
};
