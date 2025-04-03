<template>
  <div class="admin-dashboard">
    <header class="admin-header">
    <h1>Адмін Панель Управління</h1>
    <button class="logout-btn" @click="authStore.logout()">
      <i class="fas fa-sign-out-alt"></i> Вийти
    </button>
    </header>
    <!-- Users Section -->
    <section class="users-section" v-if="users.length > 0">
      <div class="section-header">
        <h2>Користувачі</h2>
        <button class="create-btn" @click="showCreateUserModal = true">
          <i class="fas fa-plus"></i> Додати користувача
        </button>
      </div>
      <div class="table-container">
        <table class="data-table">
          <thead>
          <tr>
            <th>Ім'я</th>
            <th>Прізвище</th>
            <th>Юзернейм</th>
            <th>Пошта</th>
            <th>Роль</th>
            <th>Час створення користувача</th>
            <th>Час оновлення користувача</th>
            <th>Дії</th>
          </tr>
          </thead>
          <tbody>
          <tr v-for="user in users" :key="user.UserID">
            <td>{{ user.FirstName }}</td>
            <td>{{ user.LastName }}</td>
            <td>{{ user.Username }}</td>
            <td>{{ user.Email }}</td>
            <td>
                <span :class="['role-badge', {'admin': user.RoleID === 1, 'user': user.RoleID === 2}]">
                  {{ user.RoleID === 1 ? 'Адмін' : user.RoleID === 2 ? 'Користувач' : 'Невідома роль' }}
                </span>
            </td>
            <td>{{ user.CreatedAt }}</td>
            <td>{{ user.UpdatedAt }}</td>
            <td class="actions">
              <button class="edit-btn" @click="editUser(user)">
                <i class="fas fa-edit"></i>
              </button>
              <button class="delete-btn" @click="confirmDelete('user', user.UserID, user.Username)">
                <i class="fas fa-trash-alt"></i>
              </button>
            </td>
          </tr>
          </tbody>
        </table>
      </div>
    </section>
    <div class="empty-state" v-else>
      <i class="fas fa-users empty-icon"></i>
      <p>Немає користувачів</p>
    </div>

    <!-- Tickets Section -->
    <section class="tickets-section" v-if="tickets.length > 0">
      <h2>Тікети</h2>
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
          <tr v-for="ticket in tickets" :key="ticket.TicketID">
            <td>{{ ticket.Title }}</td>
            <td class="ticket-description">{{ ticket.Description }}</td>
            <td>{{ ticket.Status }}</td>
            <td>{{ ticket.AssigneeUsername }}</td>
            <td>{{ ticket.SenderUsername }}</td>
            <td>{{ formatDate(ticket.CreatedAt) }}</td>
            <td>{{ formatDate(ticket.UpdatedAt) }}</td>
            <td class="actions">
              <button class="delete-btn" @click="confirmDelete('ticket', ticket.TicketID, ticket.Title)">
                <i class="fas fa-trash-alt"></i>
              </button>
            </td>
          </tr>
          </tbody>
        </table>
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
        <div v-for="notification in notifications" :key="notification.NotificationID" class="notification-item">
          <div class="notification-content">
            <p class="notification-message">{{ notification.Message }}</p>
            <span class="notification-time">{{ formatDate(notification.CreatedAt) }}</span>
          </div>
          <button class="delete-btn small" @click="confirmDelete('notification', notification.NotificationID, 'сповіщення')">
            <i class="fas fa-trash-alt"></i>
          </button>
        </div>
      </div>
    </section>
    <div class="empty-state" v-else>
      <i class="fas fa-bell empty-icon"></i>
      <p>Немає сповіщень</p>
    </div>

    <!-- Data Management Section -->
    <section class="data-management-section">
      <h2>Управління Даними</h2>
      <div class="data-actions">
        <button class="action-btn backup" @click="backupData">
          <i class="fas fa-database"></i> Резервне Копіювання
        </button>
        <button class="action-btn restore" @click="restoreData">
          <i class="fas fa-redo"></i> Відновити Дані
        </button>
        <button class="action-btn export" @click="exportData">
          <i class="fas fa-file-export"></i> Експортувати
        </button>
        <button class="action-btn import" @click="importData">
          <i class="fas fa-file-import"></i> Імпортувати
        </button>
      </div>
    </section>

    <!-- Create User Modal -->
    <div v-if="showCreateUserModal" class="modal-overlay">
      <div class="modal">
        <div class="modal-header">
          <h3>Створити нового користувача</h3>
          <button class="close-btn" @click="showCreateUserModal = false">
            <i class="fas fa-times"></i>
          </button>
        </div>
        <div class="modal-body">
          <form @submit.prevent="submitCreateUser">
            <div class="form-group">
              <label>Ім'я</label>
              <input v-model="newUser.FirstName" type="text" placeholder="Введіть ім'я" required />
            </div>
            <div class="form-group">
              <label>Прізвище</label>
              <input v-model="newUser.LastName" type="text" placeholder="Введіть прізвище" required />
            </div>
            <div class="form-group">
              <label>Юзернейм</label>
              <input v-model="newUser.Username" type="text" placeholder="Введіть юзернейм" required />
            </div>
            <div class="form-group">
              <label>Email</label>
              <input v-model="newUser.Email" type="email" placeholder="Введіть email" required />
            </div>
            <div class="form-group">
              <label>Пароль</label>
              <input v-model="newUser.Password" type="password" placeholder="Введіть пароль" required />
            </div>
            <div class="form-group">
              <label>Роль</label>
              <select v-model="newUser.RoleID" required>
                <option value="2">Користувач</option>
                <option value="1">Адміністратор</option>
              </select>
            </div>
            <div class="form-actions">
              <button type="button" class="cancel-btn" @click="showCreateUserModal = false">Скасувати</button>
              <button type="submit" class="submit-btn">Створити</button>
            </div>
          </form>
        </div>
      </div>
    </div>

    <!-- Edit User Modal -->
    <div v-if="editUserData" class="modal-overlay">
      <div class="modal">
        <div class="modal-header">
          <h3>Редагувати користувача</h3>
          <button class="close-btn" @click="editUserData = null">
            <i class="fas fa-times"></i>
          </button>
        </div>
        <div class="modal-body">
          <form @submit.prevent="submitEditUser">
            <div class="form-group">
              <label>Ім'я</label>
              <input v-model="editUserData.FirstName" type="text" placeholder="Введіть ім'я" required />
            </div>
            <div class="form-group">
              <label>Прізвище</label>
              <input v-model="editUserData.LastName" type="text" placeholder="Введіть прізвище" required />
            </div>
            <div class="form-group">
              <label>Юзернейм</label>
              <input v-model="editUserData.Username" type="text" placeholder="Введіть юзернейм" required />
            </div>
            <div class="form-group">
              <label>Роль</label>
              <select v-model="editUserData.RoleID" required>
                <option value="2">Користувач</option>
                <option value="1">Адміністратор</option>
              </select>
            </div>
            <div class="form-actions">
              <button type="button" class="cancel-btn" @click="editUserData = null">Скасувати</button>
              <button type="submit" class="submit-btn">Зберегти</button>
            </div>
          </form>
        </div>
      </div>
    </div>

    <!-- Confirmation Modal -->
    <div v-if="confirmModal.show" class="modal-overlay">
      <div class="confirm-modal">
        <div class="modal-header">
          <h3>Підтвердження дії</h3>
        </div>
        <div class="modal-body">
          <p>Ви дійсно хочете видалити {{ confirmModal.type === 'user' ? 'користувача' :
              confirmModal.type === 'ticket' ? 'тікет' : 'сповіщення' }} <strong>"{{ confirmModal.name }}"</strong>?</p>
        </div>
        <div class="modal-footer">
          <button class="cancel-btn" @click="confirmModal.show = false">Скасувати</button>
          <button class="confirm-btn" @click="executeDelete">Видалити</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { ref, onMounted } from "vue";
import { adminApi } from "../api";
import { format } from 'date-fns';
import { useAuthStore } from '../store/auth';
import { useRouter } from 'vue-router';

export default {
  name: "AdminDashboard",
  setup() {
    const tickets = ref([]);
    const notifications = ref([]);
    const users = ref([]);
    const editUserData = ref(null);
    const showCreateUserModal = ref(false);
    const authStore = useAuthStore();
    const router = useRouter();

    const newUser = ref({
      Username: '',
      Password: '',
      FirstName: '',
      LastName: '',
      Email: '',
      RoleID: '2',
    });

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

    const fetchTickets = async () => {
      try {
        const response = await adminApi.getAllTickets();
        tickets.value = response.data || [];
      } catch (error) {
        console.error("Помилка при завантаженні тікетів:", error);
      }
    };

    const fetchNotifications = async () => {
      try {
        const response = await adminApi.getAllNotifications();
        notifications.value = response.data || [];
      } catch (error) {
        console.error("Помилка при завантаженні сповіщень:", error);
      }
    };

    const fetchUsers = async () => {
      try {
        const response = await adminApi.getAllUsers();
        users.value = response.data || [];
      } catch (error) {
        console.error("Помилка при завантаженні користувачів:", error);
      }
    };

    const confirmDelete = (type, id, name) => {
      confirmModal.value = {
        show: true,
        type,
        id,
        name,
        callback: type === 'user' ? deleteUser :
            type === 'ticket' ? deleteTicket :
                deleteNotification
      };
    };

    const executeDelete = async () => {
      try {
        await confirmModal.value.callback(confirmModal.value.id);
        confirmModal.value.show = false;

        // Refresh data
        if (confirmModal.value.type === 'user') await fetchUsers();
        else if (confirmModal.value.type === 'ticket') await fetchTickets();
        else await fetchNotifications();
      } catch (error) {
        console.error("Помилка при видаленні:", error);
      }
    };

    const deleteTicket = async (ticketID) => {
      await adminApi.adminDeleteTicket(ticketID);
      tickets.value = tickets.value.filter(ticket => ticket.TicketID !== ticketID);
    };

    const deleteUser = async (userID) => {
      await adminApi.deleteUser(userID);
      users.value = users.value.filter(user => user.UserID !== userID);
    };

    const deleteNotification = async (notificationID) => {
      await adminApi.deleteNotification(notificationID);
      notifications.value = notifications.value.filter(notification => notification.NotificationID !== notificationID);
    };

    const editUser = (user) => {
      editUserData.value = {...user};
    };

    const submitEditUser = async () => {
      try {
        const userDataToSend = {
          ...editUserData.value,
          RoleID: parseInt(editUserData.value.RoleID),
        };

        await adminApi.updateUser(editUserData.value.UserID, userDataToSend);
        editUserData.value = null;
        await fetchUsers();
      } catch (error) {
        console.error("Помилка при оновленні користувача:", error);
      }
    };

    const submitCreateUser = async () => {
      try {
        await adminApi.createUser({
          ...newUser.value,
          RoleID: parseInt(newUser.value.RoleID)
        });

        showCreateUserModal.value = false;
        newUser.value = {
          Username: '',
          Password: '',
          FirstName: '',
          LastName: '',
          Email: '',
          RoleID: '2'
        };

        await fetchUsers();
      } catch (error) {
        console.error('Помилка при створенні користувача:', error);
      }
    };

    const backupData = async () => {
      try {
        const response = await adminApi.backupData();

        const blob = new Blob([response.data], { type: 'application/sql' });
        const downloadUrl = URL.createObjectURL(blob);

        const link = document.createElement('a');
        link.href = downloadUrl;

        const contentDisposition = response.headers['content-disposition'];
        let filename = 'backup_' + new Date().toISOString().slice(0, 10) + '.sql';

        if (contentDisposition) {
          const filenameMatch = contentDisposition.match(/filename="?(.+)"?/);
          if (filenameMatch && filenameMatch[1]) {
            filename = filenameMatch[1];
          }
        }

        link.download = filename;
        document.body.appendChild(link);
        link.click();

        setTimeout(() => {
          document.body.removeChild(link);
          URL.revokeObjectURL(downloadUrl);
        }, 100);

      } catch (error) {
        console.error('Backup error:', error);
        alert(`Backup failed: ${error.response?.data?.message || error.message}`);
      }
    };


    const restoreData = async () => {
      const fileInput = document.createElement('input');
      fileInput.type = 'file';
      fileInput.accept = '.sql';

      fileInput.onchange = async (e) => {
        const file = e.target.files[0];
        if (!file) return;

        try {
          const formData = new FormData();
          formData.append('file', file);

          await adminApi.restoreData(formData);
          alert('Дані успішно відновлені!');

          await Promise.all([
            fetchUsers(),
            fetchTickets(),
            fetchNotifications()
          ]);
        } catch (error) {
          console.error('Помилка при відновленні даних:', error);
          alert(`Помилка при відновленні даних: ${error.response?.data?.error || error.message}`);
        }
      };

      fileInput.click();
    };

    const exportData = async () => {
      try {
        const data = await adminApi.exportData();
        const blob = new Blob([data], {type: 'application/vnd.openxmlformats-officedocument.spreadsheetml.sheet'});
        const link = document.createElement('a');
        link.href = URL.createObjectURL(blob);
        link.download = `export_${new Date().toISOString().slice(0, 10)}.xlsx`;
        link.click();
        alert('Дані успішно експортовані!');
      } catch (error) {
        console.error('Помилка при експорті даних:', error);
        alert('Помилка при експорті даних!');
      }
    };

    const importData = async () => {
      const fileInput = document.createElement('input');
      fileInput.type = 'file';
      fileInput.accept = '.xlsx,.xls';

      fileInput.onchange = async (e) => {
        const file = e.target.files[0];
        if (file) {
          try {
            await adminApi.importData(file);
            alert('Дані успішно імпортовані!');
            await Promise.all([fetchUsers(), fetchTickets(), fetchNotifications()]);
          } catch (error) {
            console.error('Помилка при імпорті даних:', error);
            alert('Помилка при імпорті даних!');
          }
        }
      };

      fileInput.click();
    };

    onMounted(() => {
      Promise.all([fetchUsers(), fetchTickets(), fetchNotifications()]);
    });

    return {
      tickets,
      notifications,
      users,
      editUserData,
      newUser,
      showCreateUserModal,
      confirmModal,
      authStore,
      formatDate,
      confirmDelete,
      executeDelete,
      editUser,
      submitEditUser,
      submitCreateUser,
      backupData,
      restoreData,
      exportData,
      importData
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

.admin-dashboard {
  max-width: 1400px;
  margin: 0 auto;
  padding: 2rem;
}
.admin-header {
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
  font-size: 0.95rem;
}

.data-table th {
  background-color: var(--primary-color);
  color: var(--white);
  padding: 1rem;
  text-align: left;
  font-weight: 500;
}

.data-table td {
  padding: 1rem;
  border-bottom: 1px solid var(--border-color);
  vertical-align: middle;
}

.data-table tr:last-child td {
  border-bottom: none;
}

.data-table tr:hover {
  background-color: rgba(67, 97, 238, 0.05);
}

.role-badge {
  display: inline-block;
  padding: 0.35rem 0.65rem;
  border-radius: 50px;
  font-size: 0.8rem;
  font-weight: 500;
}

.role-badge.admin {
  background-color: rgba(244, 63, 94, 0.1);
  color: #f43f5e;
}

.role-badge.user {
  background-color: rgba(16, 185, 129, 0.1);
  color: #10b981;
}

.status-badge {
  display: inline-block;
  padding: 0.35rem 0.65rem;
  border-radius: 50px;
  font-size: 0.8rem;
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

.action-btn {
  padding: 0.8rem 1.5rem;
  border-radius: 8px;
  font-weight: 500;
}
.data-management-section h2{
  padding-top: 30px;
}
.backup {
  background-color: #10b981;
  color: white;
}

.backup:hover {
  background-color: #0d9e6e;
}

.restore {
  background-color: #3b82f6;
  color: white;
}

.restore:hover {
  background-color: #2563eb;
}

.export {
  background-color: #f59e0b;
  color: white;
}

.export:hover {
  background-color: #d97706;
}

.import {
  background-color: #8b5cf6;
  color: white;
}

.import:hover {
  background-color: #7c3aed;
}

.data-actions {
  display: flex;
  gap: 1rem;
  flex-wrap: wrap;
  margin-bottom: 2rem;
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

/* Modal Styles */
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
  background-color: white;
  align-items: center;
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
@media (max-width: 768px) {
  .admin-dashboard {
    padding: 1rem;
  }

  .data-actions {
    flex-direction: column;
  }

  .action-btn {
    width: 100%;
  }

  .modal {
    width: 95%;
  }
  .admin-header {
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