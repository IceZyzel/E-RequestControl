<template>
  <div class="dashboard">
    <header class="dashboard-header">
    <h1>Ласкаво просимо до панелі управління тікетами</h1>
    <button class="logout-btn" @click="authStore.logout()">
      <i class="fas fa-sign-out-alt"></i> Вийти
    </button>
    </header>
    <!-- Tickets Section -->
    <section class="tickets-section" v-if="tickets.length > 0">
      <div class="section-header">
        <h2>Тікети</h2>
      </div>
      <div class="table-container">
        <table class="data-table">
          <thead>
          <tr>
            <th>Назва</th>
            <th>Опис</th>
            <th>Статус</th>
            <th>Призначено</th>
            <th>Відправник</th>
            <th>Створено</th>
            <th>Оновлено</th>
            <th>Дії</th>
          </tr>
          </thead>
          <tbody>
          <tr v-for="ticket in paginatedTickets" :key="ticket.TicketID">
            <td>{{ ticket.Title }}</td>
            <td class="ticket-description">
                <span class="description-text">
                  {{ truncateDescription(ticket.Description, ticket.TicketID) }}
                </span>
              <button
                  v-if="ticket.Description && ticket.Description.length > 100"
                  class="show-more-btn"
                  @click="toggleDescription(ticket.TicketID)"
              >
                {{ expandedDescriptions[ticket.TicketID] ? 'Менше' : 'Більше' }}
              </button>
            </td>
            <td>
                <span :class="['status-badge', getStatusClass(ticket.Status)]">
                  {{ ticket.Status }}
                </span>
            </td>
            <td>{{ ticket.AssigneeUsername || 'Не призначено' }}</td>
            <td>{{ ticket.SenderUsername }}</td>
            <td>{{ formatDate(ticket.CreatedAt) }}</td>
            <td>{{ formatDate(ticket.UpdatedAt) }}</td>
            <td class="actions">
              <button class="edit-btn" @click="editTicket(ticket)" title="Редагувати">
                <i class="fas fa-edit"></i>
              </button>
              <button class="delete-btn" @click="confirmDelete('ticket', ticket.TicketID, ticket.Title)" title="Видалити">
                <i class="fas fa-trash-alt"></i>
              </button>
            </td>
          </tr>
          </tbody>
        </table>
        <div class="pagination-controls" v-if="totalPages > 1">
          <button class="pagination-btn" @click="prevPage" :disabled="currentPage === 1">
            <i class="fas fa-chevron-left"></i>
          </button>
          <span class="page-info">
            Сторінка {{ currentPage }} з {{ totalPages }}
          </span>
          <button class="pagination-btn" @click="nextPage" :disabled="currentPage === totalPages">
            <i class="fas fa-chevron-right"></i>
          </button>
        </div>
      </div>
    </section>
    <div class="empty-state" v-else>
      <i class="fas fa-ticket-alt empty-icon"></i>
      <p>Список тікетів порожній</p>
    </div>

    <!-- Notifications Section -->
    <section class="notifications-section" v-if="notifications.length > 0">
      <h2>Сповіщення</h2>
      <div class="notifications-list">
        <div v-for="notification in paginatedNotifications" :key="notification.NotificationID" class="notification-item">
          <div class="notification-content">
            <p class="notification-message">{{ notification.Message }}</p>
            <span class="notification-time">{{ formatDate(notification.CreatedAt) }}</span>
          </div>
          <button class="delete-btn small" @click="confirmDelete('notification', notification.NotificationID, 'сповіщення')">
            <i class="fas fa-trash-alt"></i>
          </button>
        </div>
        <div class="pagination-controls" v-if="notifications.length > itemsPerPage">
          <button class="pagination-btn" @click="prevNotificationsPage" :disabled="currentNotificationsPage === 1">
            <i class="fas fa-chevron-left"></i>
          </button>
          <span class="page-info">
          Сторінка {{ currentNotificationsPage }} з {{ totalNotificationsPages }}
        </span>
          <button class="pagination-btn" @click="nextNotificationsPage" :disabled="currentNotificationsPage === totalNotificationsPages">
            <i class="fas fa-chevron-right"></i>
          </button>
        </div>
      </div>

    </section>
    <div class="empty-state" v-else>
      <i class="fas fa-bell empty-icon"></i>
      <p>Немає сповіщень</p>
    </div>

    <!-- Create Ticket Modal -->
    <transition name="modal-fade">
    <div v-if="showCreateTicketModal" class="modal-overlay" @click.self="showCreateTicketModal= false" >
      <div class="modal">
        <div class="modal-header">
          <h3>Створити новий тікет</h3>
          <button class="close-btn" @click="showCreateTicketModal = false">
            <i class="fas fa-times"></i>
          </button>
        </div>
        <div class="modal-body">
          <form @submit.prevent="addTicket">
            <div class="form-group">
              <label>Заголовок</label>
              <input v-model="newTicket.Title" type="text" placeholder="Введіть заголовок" required />
            </div>
            <div class="form-group">
              <label>Опис</label>
              <textarea v-model="newTicket.Description" placeholder="Введіть опис" required></textarea>
            </div>
            <div class="form-group">
              <label>Призначити</label>
              <select v-model="newTicket.AssignedTo" required>
                <option v-for="user in users" :key="user.UserID" :value="user.UserID">
                  {{ user.FirstName }} {{ user.LastName }} ({{ user.Username }})
                </option>
              </select>
            </div>
            <div class="form-actions">
              <button type="button" class="cancel-btn" @click="showCreateTicketModal = false">Скасувати</button>
              <button type="submit" class="submit-btn">Створити</button>
            </div>
          </form>
        </div>
      </div>
    </div>
    </transition>

    <!-- Edit Ticket Modal -->
    <transition name="modal-fade">
    <div v-if="editTicketData" class="modal-overlay" @click.self="editTicketData = null">
      <div class="modal">
        <div class="modal-header">
          <h3>Редагувати тікет</h3>
          <button class="close-btn" @click="editTicketData = null">
            <i class="fas fa-times"></i>
          </button>
        </div>
        <div class="modal-body">
          <form @submit.prevent="updateTicket">
            <div class="form-group">
              <label>Заголовок</label>
              <input v-model="editTicketData.Title" type="text" placeholder="Введіть заголовок" required />
            </div>
            <div class="form-group">
              <label>Опис</label>
              <textarea v-model="editTicketData.Description" placeholder="Введіть опис" required></textarea>
            </div>
            <div class="form-group">
              <label>Призначити</label>
              <select v-model="editTicketData.AssignedTo" required>
                <option v-for="user in users" :key="user.UserID" :value="user.UserID">
                  {{ user.Username }}
                </option>
              </select>
            </div>
            <div class="form-actions">
              <button type="button" class="cancel-btn" @click="editTicketData = null">Скасувати</button>
              <button type="submit" class="submit-btn">Оновити</button>
            </div>
          </form>
        </div>
      </div>
    </div>
    </transition>

    <!-- Confirmation Modal -->
    <transition name="modal-fade">
    <div v-if="confirmModal.show" class="modal-overlay" @click.self="confirmModal.show = false">
      <div class="confirm-modal">
        <div class="modal-header">
          <h3>Підтвердження дії</h3>
        </div>
        <div class="modal-body">
          <p>Ви дійсно хочете видалити тікет <strong>"{{ confirmModal.name }}"</strong>?</p>
        </div>
        <div class="modal-footer">
          <button class="cancel-btn" @click="confirmModal.show = false">Скасувати</button>
          <button class="confirm-btn" @click="executeDelete">Видалити</button>
        </div>
      </div>
    </div>
    </transition>
    <!-- Create Ticket Button -->
    <button class="create-btn floating-btn" @click="showCreateTicketModal = true">
      <i class="fas fa-plus"></i> Додати тікет
    </button>
  </div>
</template>

<script>
import { ref, onMounted, computed } from "vue";
import { requestApi, notificationApi } from "../api";
import { useAuthStore } from "../store/auth";
import { format } from 'date-fns';
import {useRouter} from "vue-router";

export default {
  name: "TicketDashboard",
  setup() {
    const tickets = ref([]);
    const currentPage = ref(1);
    const notifications = ref([]);
    const users = ref([]);
    const authStore = useAuthStore();
    const router = useRouter();
    const expandedDescriptions = ref({});
    const itemsPerPage = ref(10);
    const currentNotificationsPage = ref(1);

    const newTicket = ref({
      Title: '',
      Description: '',
      AssignedTo: null,
    });
    const editTicketData = ref(null);
    const showCreateTicketModal = ref(false);

    const confirmModal = ref({
      show: false,
      type: '',
      id: null,
      name: '',
      callback: null
    });

    const formatDate = (dateString) => {
      return format(new Date(dateString), 'dd.MM.yyyy HH:mm');
    };

    const fetchUsers = async () => {
      try {
        const response = await requestApi.getAllUsers();
        users.value = response.data || [];
      } catch (error) {
        console.error("Помилка при завантаженні користувачів:", error);
      }
    };

    const fetchTickets = async () => {
      try {
        const response = await requestApi.getUserTickets();
        tickets.value = (response.data || []).sort((a, b) =>
            new Date(b.CreatedAt) - new Date(a.CreatedAt)
        );
      } catch (error) {
        console.error("Помилка при завантаженні тікетів:", error);
        tickets.value = [];
      }
    };
    const truncateDescription = (description, ticketId) => {
      if (!description) return '';
      const shouldTruncate = description.length > 100 && !expandedDescriptions.value[ticketId];
      return shouldTruncate ? `${description.substring(0, 100)}...` : description;
    };
    const toggleDescription = (ticketId) => {
      expandedDescriptions.value = {
        ...expandedDescriptions.value,
        [ticketId]: !expandedDescriptions.value[ticketId]
      };
    };

    const paginatedTickets = computed(() => {
      const start = (currentPage.value - 1) * itemsPerPage.value;
      const end = start + itemsPerPage.value;
      return tickets.value.slice(start, end);
    });

    const nextPage = () => {
      if (currentPage.value < totalPages.value) currentPage.value++;
    };

    const totalPages = computed(() => Math.ceil(tickets.value.length / itemsPerPage.value));

    const prevPage = () => {
      if (currentPage.value > 1) currentPage.value--;
    };

    const fetchNotifications = async () => {
      try {
        const response = await notificationApi.getUserNotifications();
        notifications.value = (response.data || []).sort((a, b) =>
            new Date(b.CreatedAt) - new Date(a.CreatedAt)
        );
      } catch (error) {
        console.error("Помилка при завантаженні сповіщень:", error);
        notifications.value = [];
      }
    };
    const paginatedNotifications = computed(() => {
      const start = (currentNotificationsPage.value - 1) * itemsPerPage.value;
      const end = start + itemsPerPage.value;
      return notifications.value.slice(start, end);
    });

    const totalNotificationsPages = computed(() =>
        Math.ceil(notifications.value.length / itemsPerPage.value)
    );
    const nextNotificationsPage = () => {
      if (currentNotificationsPage.value < totalNotificationsPages.value) currentNotificationsPage.value++;
    };
    const prevNotificationsPage = () => {
      if (currentNotificationsPage.value > 1) currentNotificationsPage.value--;
    };
    const confirmDelete = (type, id, name) => {
      confirmModal.value = {
        show: true,
        type,
        id,
        name,
        callback: type === 'ticket' ? deleteTicket : deleteNotification
      };
    };

    const executeDelete = async () => {
      try {
        await confirmModal.value.callback(confirmModal.value.id);
        confirmModal.value.show = false;
        await fetchTickets();
      } catch (error) {
        console.error("Помилка при видаленні:", error);
      }
    };

    const getStatusClass = (status) => {
      switch (status) {
        case 'Новий': return 'status-new';
        case 'Оновлено': return 'status-in-progress';
        default: return '';
      }
    };
    const deleteTicket = async (ticketID) => {
      try {
        await requestApi.deleteTicket(ticketID);
        tickets.value = tickets.value.filter(ticket => ticket.TicketID !== ticketID);
      } catch (error) {
        console.error("Помилка при видаленні тікета:", error);
      }
    };
    const deleteNotification = async (notificationID) => {
      await notificationApi.markAsRead(notificationID);
      notifications.value = notifications.value.filter(notification => notification.NotificationID !== notificationID);
    };
    const addTicket = async () => {
      try {
        await requestApi.createTicket(newTicket.value);
        newTicket.value = { Title: '', Description: '', AssignedTo: null };
        showCreateTicketModal.value = false;
        await fetchTickets();
        await fetchNotifications();
      } catch (error) {
        console.error("Помилка при додаванні тікета:", error);
      }
    };

    const editTicket = (ticket) => {
      editTicketData.value = { ...ticket };
    };

    const updateTicket = async () => {
      try {
        await requestApi.updateTicket(editTicketData.value.TicketID, editTicketData.value);
        editTicketData.value = null;
        await Promise.all([
          fetchTickets(),
          fetchNotifications()
        ]);
      } catch (error) {
        console.error("Помилка при оновленні тікета:", error);
      }
    };

    const markAsRead = async (notificationID) => {
      try {
        await notificationApi.markAsRead(notificationID);
        await fetchNotifications();
      } catch (error) {
        console.error("Помилка при оновленні повідомлення:", error);
      }
    };

    onMounted(() => {
      fetchTickets();
      fetchNotifications();
      fetchUsers();
    });

    return {
      tickets,
      notifications,
      users,
      newTicket,
      editTicketData,
      showCreateTicketModal,
      confirmModal,
      authStore,
      formatDate,
      confirmDelete,
      executeDelete,
      deleteTicket,
      addTicket,
      editTicket,
      updateTicket,
      markAsRead,

      paginatedNotifications,
      nextNotificationsPage,
      prevPage,
      nextPage,
      paginatedTickets,
      currentPage,
      totalPages,
      truncateDescription,
      toggleDescription,
      getStatusClass,
      expandedDescriptions,
      currentNotificationsPage,
      totalNotificationsPages,
      prevNotificationsPage,
      itemsPerPage
    };
  }
};
</script>

<style scoped>

:root {
  --primary-color: #4361ee;
  --secondary-color: #3f37c9;
  --success-color: #4cc9f0;
  --danger-color: #f72585;
  --warning-color: #f8961e;
  --info-color: #4895ef;
  --light-color: #f8f9fa;
  --dark-color: #212529;
  --gray-color: #6c757d;
  --border-color: #dee2e6;
  --white: #ffffff;
}

* {
  box-sizing: border-box;
  margin: 0;
  padding: 0;
}

body {
  font-family: 'Roboto', sans-serif;
  line-height: 1.6;
  color: var(--dark-color);
  background-color: #f5f7fa;
}

.dashboard {
  max-width: 1400px;
  margin: 0 auto;
  padding: 2rem;
}
.dashboard-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 2rem;
  position: relative;
}
h1 {
  text-align: center;
  margin-bottom: 2rem;
  color: var(--primary-color);
  font-size: 2.2rem;
  font-weight: 700;
}

h2 {
  font-size: 1.5rem;
  font-weight: 600;
  color: var(--dark-color);
  margin-bottom: 1.5rem;
}

.section-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 1.5rem;
}

.empty-state {
  text-align: center;
  padding: 2rem;
  background: var(--white);
  border-radius: 8px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.05);
  margin-bottom: 2rem;
}

.empty-icon {
  font-size: 3rem;
  color: var(--gray-color);
  margin-bottom: 1rem;
}

.empty-state p {
  color: var(--gray-color);
  font-size: 1.1rem;
}
.table-container {
  overflow-x: auto;
  background: var(--white);
  border-radius: 8px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.05);
  margin-bottom: 2rem;
}

.data-table {
  width: 100%;
  border-collapse: collapse;
}

.data-table th {
  background-color: var(--primary-color);
  color: var(--white);
  padding: 1rem;
  text-align: left;
  font-weight: 500;
}

.data-table th,
.data-table td {
  padding: 12px 15px;
  text-align: left;
  border-bottom: 1px solid #e0e0e0;
}

.data-table tr:last-child td {
  border-bottom: none;
}

.data-table tr:hover {
  background-color: rgba(67, 97, 238, 0.05);
}

.status-badge {
  display: inline-block;
  padding: 0.35rem 0.65rem;
  border-radius: 50px;
  font-weight: 500;
}

.status-badge.open {
  background-color: rgba(59, 130, 246, 0.1);
  color: #3b82f6;
}

.status-badge.in-progress {
  background-color: rgba(249, 115, 22, 0.1);
  color: #f97316;
}

.status-badge.resolved {
  background-color: rgba(16, 185, 129, 0.1);
  color: #10b981;
}

.status-badge.closed {
  background-color: rgba(107, 114, 128, 0.1);
  color: #6b7280;
}

.actions {
  display: flex;
  gap: 0.5rem;
}

button {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  gap: 0.5rem;
  padding: 0.6rem 1.2rem;
  border-radius: 6px;
  font-size: 0.95rem;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.2s ease;
  border: none;
}

button i {
  font-size: 0.9rem;
}

.create-btn {
  background-color: var(--primary-color);
  color: var(--white);
}

.create-btn:hover {
  background-color: var(--secondary-color);
}

.edit-btn {
  background-color: var(--info-color);
  color: var(--white);
  padding: 0.5rem;
}

.edit-btn:hover {
  background-color: #3a7bd5;
}

.delete-btn {
  background-color: var(--danger-color);
  color: var(--white);
  padding: 0.5rem;
}

.delete-btn:hover {
  background-color: #e5177e;
}

.delete-btn.small {
  padding: 0.4rem 0.6rem;
}

.pagination-controls {
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 1rem;
  gap: 1rem;
  background-color: #f8f9fa;
}

.page-info {
  font-size: 0.9rem;
  color: var(--gray-color);
}
.notifications-list {
  background: var(--white);
  border-radius: 8px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.05);
  padding: 1rem;
}

.notification-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 1rem;
  border-bottom: 1px solid var(--border-color);
}

.notification-item:last-child {
  border-bottom: none;
}

.notification-content {
  flex: 1;
}

.notification-message {
  font-weight: 500;
  margin-bottom: 0.25rem;
}

.notification-time {
  font-size: 0.85rem;
  color: var(--gray-color);
}


.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  opacity: 1;
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 1000;
}

.modal, .confirm-modal {
  background-color: white;
  border-radius: 10px;
  width: 100%;
  max-width: 500px;
  box-shadow: 0 5px 20px rgba(0, 0, 0, 0.2);
  overflow: hidden;
}

.confirm-modal {
  max-width: 400px;
}

.modal-header {
  padding: 1.5rem;
  border-bottom: 1px solid var(--border-color);
  display: flex;
  justify-content: space-between;
  align-items: center;
  background-color: white;
}

.modal-header h3 {
  font-size: 1.3rem;
  font-weight: 600;
  color: var(--dark-color);
}

.close-btn {
  background: none;
  border: none;
  color: var(--gray-color);
  font-size: 1.2rem;
  cursor: pointer;
  padding: 0.2rem;
}

.close-btn:hover {
  color: var(--danger-color);
}

.modal-body {
  padding: 1.5rem;
  background-color: white;
}

.form-group {
  margin-bottom: 1.25rem;
}

.form-group label {
  display: block;
  margin-bottom: 0.5rem;
  font-weight: 500;
  color: var(--dark-color);
}

.form-group input,
.form-group select {
  width: 100%;
  padding: 0.75rem;
  border: 1px solid var(--border-color);
  border-radius: 6px;
  font-size: 1rem;
  transition: border-color 0.2s;
}

.form-group input:focus,
.form-group select:focus {
  outline: none;
  border-color: var(--primary-color);
  box-shadow: 0 0 0 3px rgba(67, 97, 238, 0.1);
}

.form-actions {
  display: flex;
  justify-content: flex-end;
  gap: 1rem;
  margin-top: 1.5rem;
}

.cancel-btn {
  background-color: var(--gray-color);
  color: var(--white);
}

.cancel-btn:hover {
  background-color: #5a6268;
}

.submit-btn {
  background-color: var(--primary-color);
  color: var(--white);
}
.show-more-btn {
  background: none;
  border: none;
  color: var(--primary-color);
  cursor: pointer;
  font-size: 0.8rem;
  padding: 0.2rem 0.5rem;
  margin-left: 0.5rem;
  white-space: nowrap;
}
.submit-btn:hover {
  background-color: var(--secondary-color);
}

.confirm-btn {
  background-color: var(--danger-color);
  color: var(--white);
}

.confirm-btn:hover {
  background-color: #e5177e;
}

.modal-footer {
  padding: 1rem 1.5rem;
  border-top: 1px solid var(--border-color);
  display: flex;
  justify-content: flex-end;
  gap: 1rem;
}

.data-table tr:last-child td {
  border-bottom: none;
}

.ticket-description {
  max-width: 300px;
  word-break: break-word;
}

.mark-read-btn {
  background-color: var(--success-color);
  color: var(--white);
  padding: 0.5rem 1rem;
}

.mark-read-btn:hover {
  background-color: #3aa8d8;
}

.floating-btn {
  position: fixed;
  bottom: 2rem;
  right: 2rem;
  z-index: 100;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
}

textarea {
  min-height: 120px;
  padding: 0.75rem;
  width: 100%;
  border: 1px solid var(--border-color);
  border-radius: 6px;
  font-size: 1rem;
  transition: border-color 0.2s;
  resize: none;
}

textarea:focus {
  outline: none;
  border-color: var(--primary-color);
  box-shadow: 0 0 0 3px rgba(67, 97, 238, 0.1);
}
.logout-btn {
  background-color: var(--danger-color);
  color: var(--white);
  padding: 0.8rem 1.5rem;
  border-radius: 8px;
  transition: all 0.2s ease;
  display: flex;
  align-items: center;
  gap: 0.7rem;
  font-size: 1rem;
}

.logout-btn:hover {
  background-color: #e5177e;
  transform: translateY(-1px);
  box-shadow: 0 2px 8px rgba(244, 63, 94, 0.3);
}

.logout-btn i {
  font-size: 1.1rem;
}

.pagination-btn {
  background-color: var(--primary-color);
  color: white;
  border: none;
}

.pagination-btn:disabled {
  background-color: #e0e0e0 !important;
  color: #9e9e9e !important;
}

.cancel-btn {
  background-color: var(--gray-color) !important;
  color: white !important;
}

.submit-btn, .confirm-btn {
  background-color: var(--primary-color) !important;
  color: white !important;
}

.confirm-btn {
  background-color: var(--danger-color) !important;
}

button:not(:disabled):hover {
  transform: translateY(-1px);
  box-shadow: 0 2px 5px rgba(0,0,0,0.2);
}
button:disabled {
  background-color: var(--gray-color) !important;
  color: var(--light-color) !important;
  cursor: not-allowed;
  opacity: 0.7;
}
.modal-fade-enter-active,
.modal-fade-leave-active {
  transition: opacity 0.3s ease;
}

.modal-fade-enter-from,
.modal-fade-leave-to {
  opacity: 0;
}

.modal-fade-enter-active .modal,
.modal-fade-leave-active .modal {
  transition: transform 0.3s ease, opacity 0.3s ease;
}

.modal-fade-enter-from .modal,
.modal-fade-leave-to .modal {
  transform: translateY(-20px);
  opacity: 0;
}
@media (max-width: 768px) {
  .dashboard {
    padding: 1rem;
  }

  .modal {
    width: 95%;
  }
  .dashboard-header {
    flex-direction: column;
    gap: 1rem;
    text-align: center;
  }

  .logout-btn {
    width: 100%;
    justify-content: center;
  }
}
</style>