<template>
  <div class="dashboard">
    <h1>Ласкаво просимо до панелі управління тікетами</h1>

    <section class="users" v-if="users.length > 0">
    </section>

    <section class="tickets" v-if="tickets.length > 0">
      <h2>Тікети</h2>
      <table class="tickets-table">
        <thead>
        <tr>
          <th>Заголовок</th>
          <th>Опис</th>
          <th>Статус</th>
          <th>Призначено</th>
          <th>Відправник</th>
          <th>Оновлено</th>
          <th>Дії</th>
        </tr>
        </thead>
        <tbody>
        <tr v-for="ticket in tickets" :key="ticket.TicketID">
          <td>{{ ticket.Title }}</td>
          <td class="ticket-description">{{ ticket.Description }}</td>
          <td>{{ ticket.Status }}</td>
          <td>{{ ticket.AssigneeUsername }}</td>
          <td>{{ ticket.SenderUsername }}</td>
          <td>{{ ticket.UpdatedAt }}</td>
          <td>
            <button @click="deleteTicket(ticket.TicketID)">Видалити</button>
            <button @click="editTicket(ticket)">Оновити</button>
          </td>
        </tr>
        </tbody>

      </table>
    </section>
    <p v-else>Список тікетів пустий</p>

    <section class="notifications" v-if="notifications.length > 0">
      <h2>Сповіщення</h2>
      <div v-for="notification in notifications" :key="notification.NotificationID" class="notification-item">
        <span>{{ notification.Message }}</span>
        <span>{{ notification.CreatedAt }}</span>
        <button @click="markAsRead(notification.NotificationID)">Помітити як прочитане</button>
      </div>
    </section>
    <p v-else>Немає сповіщень</p>

    <section class="create-ticket">
      <h2>Додати тікет</h2>
      <form @submit.prevent="addTicket">
        <input v-model="newTicket.Title" type="text" placeholder="Заголовок тікета" required />
        <textarea v-model="newTicket.Description" placeholder="Опис тікета" required></textarea>

        <select v-model="newTicket.AssignedTo" required>
          <option v-for="user in users" :key="user.user_id" :value="user.user_id">
            {{ user.firstname }} {{ user.lastname }} ({{ user.username }})
          </option>
        </select>

        <button :disabled="!newTicket.Title || !newTicket.Description">Додати тікет</button>
      </form>
    </section>

    <section v-if="editTicketData" class="edit-ticket">
      <h2>Редагувати тікет</h2>
      <form @submit.prevent="updateTicket">
        <input v-model="editTicketData.Title" type="text" placeholder="Заголовок тікета" required />
        <textarea v-model="editTicketData.Description" placeholder="Опис тікета" required></textarea>

        <select v-model="editTicketData.AssignedTo" required>
          <option v-for="user in users" :key="user.user_id" :value="user.user_id">
            {{ user.username }}
          </option>
        </select>

        <button :disabled="!editTicketData.Title || !editTicketData.Description">Оновити тікет</button>
      </form>
    </section>
  </div>
</template>

<script>
import { ref, onMounted, watch } from 'vue';
import { requestApi, notificationApi } from '../api';
import { useAuthStore } from '../store/auth';

export default {
  name: 'Dashboard',
  setup() {
    const tickets = ref([]);
    const notifications = ref([]);
    const users = ref([]);
    const authStore = useAuthStore();
    const newTicket = ref({
      Title: '',
      Description: '',
      AssignedTo: null,
    });
    const editTicketData = ref(null);

    const loggedInUser = authStore.user;

    const fetchUsers = async () => {
      try {
        const response = await requestApi.getAllUsers();
        users.value = response.data || [];
      } catch (error) {
        console.error('Помилка при завантаженні користувачів:', error);
      }
    };

    const fetchTickets = async () => {
      try {
        const response = await requestApi.getUserTickets();
        tickets.value = response.data || [];
      } catch (error) {
        console.error('Помилка при завантаженні тікетів:', error);
      }
    };

    const fetchNotifications = async () => {
      try {
        const response = await notificationApi.getUserNotifications();
        notifications.value = response.data || [];
      } catch (error) {
        console.error('Помилка при завантаженні повідомлень:', error);
      }
    };

    const deleteTicket = async (ticketID) => {
      try {
        await requestApi.deleteTicket(ticketID);
        tickets.value = tickets.value.filter(ticket => ticket.TicketID !== ticketID);
      } catch (error) {
        console.error('Помилка при видаленні тікета:', error);
      }
    };

    const addTicket = async () => {
      try {
        await requestApi.createTicket(newTicket.value);
        newTicket.value = { Title: '', Description: '', AssignedTo: null };
        await fetchTickets();
        await fetchNotifications();
      } catch (error) {
        console.error('Помилка при додаванні тікета:', error);
      }
    };

    const editTicket = (ticket) => {
      editTicketData.value = { ...ticket };
    };

    const updateTicket = async () => {
      try {
        await requestApi.updateTicket(editTicketData.value.TicketID, editTicketData.value);
        editTicketData.value = null;
        await fetchTickets();
        await fetchNotifications();
      } catch (error) {
        console.error('Помилка при оновленні тікета:', error);
      }
    };

    const markAsRead = async (notificationID) => {
      try {
        await notificationApi.markAsRead(notificationID);
        await fetchNotifications();
      } catch (error) {
        console.error('Помилка при оновленні повідомлення:', error);
      }
    };

    watch(() => loggedInUser, fetchUsers);
    onMounted(() => {
      fetchTickets();
      fetchNotifications();
      fetchUsers();
    });

    return { tickets, notifications, users, newTicket, editTicketData, deleteTicket, addTicket, editTicket, updateTicket, markAsRead };
  },
};
</script>

<style scoped>
.dashboard {
  padding: 30px;
  background: #f8f9fa;
  border-radius: 12px;
  max-width: 1100px;
  margin: 0 auto;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
}

h1 {
  margin-top:50px;
  text-align: center;
  margin-bottom: 25px;
  color: #222;
  font-size: 24px;
}
section {
  margin-bottom: 40px;
  background: white;
  padding: 20px;
  border-radius: 12px;
  box-shadow: 0 2px 6px rgba(0, 0, 0, 0.1);
}

h2 {
  text-align: center;
  color: #333;
  margin-bottom: 15px;
  font-size: 20px;
}
td button {
  width: 120px;
}

table {
  width: 100%;
  border-collapse: separate;
  border-spacing: 0;
  margin-bottom: 20px;
  border-radius: 10px;
  overflow: hidden;
}

th, td {
  padding: 12px;
  text-align: left;
  border-bottom: 1px solid #ddd;
}
td button {
  margin-bottom: 10px;
}

td button:last-child {
  margin-right: 0;
}

th {
  background-color: #007bff;
  color: white;
  font-weight: bold;
}

td {
  background-color: #fff;
}

tr:hover td {
  background-color: #f1f3f5;
}

.ticket-description {
  white-space: normal;
  word-break: break-word;
  max-width: 300px;
}

button {
  background-color: #007bff;
  color: white;
  border: none;
  padding: 10px 15px;
  cursor: pointer;
  border-radius: 8px;
  font-size: 14px;
  transition: 0.2s ease-in-out;
}

button:hover {
  background-color: #0056b3;
}

button:disabled {
  background-color: #d3d3d3;
  cursor: not-allowed;
}

.notification-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 15px;
  background: #f8f9fa;
  border-left: 6px solid #007bff;
  border-radius: 10px;
  margin-bottom: 10px;
  box-shadow: 0 2px 5px rgba(0, 0, 0, 0.1);
}

.create-ticket form,
.edit-ticket form {
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.create-ticket input,
.create-ticket textarea,
.create-ticket select,
.edit-ticket input,
.edit-ticket textarea,
.edit-ticket select {
  padding: 12px;
  font-size: 16px;
  border-radius: 8px;
  border: 1px solid #ccc;
  transition: 0.2s;
}

.create-ticket input:focus,
.create-ticket textarea:focus,
.create-ticket select:focus,
.edit-ticket input:focus,
.edit-ticket textarea:focus,
.edit-ticket select:focus {
  border-color: #007bff;
  outline: none;
  box-shadow: 0 0 5px rgba(0, 123, 255, 0.3);
}

.create-ticket button,
.edit-ticket button {
  font-size: 16px;
  padding: 12px;
  border-radius: 8px;
}

p {
  text-align: center;
  color: #777;
  font-style: italic;
  font-size: 16px;
}

</style>
