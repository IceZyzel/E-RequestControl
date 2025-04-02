<template>
  <div class="login-container">
    <div class="login-form">
      <h1>Вхід в систему</h1>
      <form @submit.prevent="handleSubmit">
        <div class="input-group">
          <label for="username">Логін</label>
          <input
              type="text"
              v-model="username"
              id="username"
              placeholder="Введіть ваш логін"
              required
              class="form-input"
          />
        </div>

        <div class="input-group">
          <label for="password">Пароль</label>
          <div class="password-wrapper">
            <input
                :type="isPasswordVisible ? 'text' : 'password'"
                v-model="password"
                id="password"
                placeholder="Введіть ваш пароль"
                required
                class="form-input"
            />
            <button
                type="button"
                @click="togglePasswordVisibility"
                class="eye-button"
                :class="{ 'active': isPasswordVisible }"
            >
              <i class="fas" :class="isPasswordVisible ? 'fa-eye-slash' : 'fa-eye'"></i>
            </button>
          </div>
        </div>

        <div class="actions">
          <button type="submit" class="btn-submit">
            <i class="fas fa-sign-in-alt"></i> Увійти
          </button>
          <div class="links">
            <router-link to="/register" class="link">
              <i class="fas fa-user-plus"></i> Зареєструватися
            </router-link>
            <router-link to="/registerAdmin" class="link admin">
              <i class="fas fa-user-shield"></i> Зареєструватися як адміністратор
            </router-link>
          </div>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import { useAuthStore } from '../store/auth';

const username = ref('');
const password = ref('');
const isPasswordVisible = ref(false);
const router = useRouter();
const authStore = useAuthStore();

const togglePasswordVisibility = () => {
  isPasswordVisible.value = !isPasswordVisible.value;
};

const handleSubmit = async () => {
  try {
    const success = await authStore.login({
      username: username.value,
      password: password.value,
    });

    if (!success) return;

    if (authStore.role === 1) {
      router.push('/admin-dashboard');
    } else {
      router.push('/dashboard');
    }
  } catch (error) {
    console.error('Помилка при авторизації:', error);
    alert('Не вдалося увійти: ' + (error.response?.data?.message || error.message));
  }
};
</script>

<style scoped>
.login-container {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 100vh;
  background-color: #f5f7fa;
  padding: 20px;
}

.login-form {
  background-color: var(--white);
  padding: 2.5rem;
  border-radius: 12px;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.1);
  width: 100%;
  max-width: 450px;
}

h1 {
  text-align: center;
  margin-bottom: 2rem;
  color: var(--primary-color);
  font-size: 1.8rem;
  font-weight: 600;
}

.input-group {
  margin-bottom: 1.5rem;
}

.input-group label {
  display: block;
  margin-bottom: 0.75rem;
  font-weight: 500;
  color: var(--dark-color);
  font-size: 0.95rem;
}

.form-input {
  width: 100%;
  padding: 0.9rem 1rem;
  font-size: 1rem;
  border-radius: 8px;
  border: 1px solid var(--border-color);
  transition: all 0.2s ease;
}

.form-input:focus {
  outline: none;
  border-color: var(--primary-color);
  box-shadow: 0 0 0 3px rgba(67, 97, 238, 0.1);
}

.password-wrapper {
  position: relative;
}

.eye-button {
  position: absolute;
  top: 50%;
  right: 12px;
  transform: translateY(-50%);
  background: none;
  border: none;
  color: var(--gray-color);
  font-size: 1rem;
  cursor: pointer;
  padding: 0.5rem;
  transition: color 0.2s ease;
}

.eye-button:hover,
.eye-button.active {
  color: var(--primary-color);
}

.actions {
  margin-top: 2rem;
}

.btn-submit {
  background-color: var(--primary-color);
  color: var(--white);
  padding: 1rem;
  border: none;
  border-radius: 8px;
  cursor: pointer;
  font-size: 1rem;
  font-weight: 500;
  width: 100%;
  transition: background-color 0.2s ease;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 0.5rem;
}

.btn-submit:hover {
  background-color: var(--secondary-color);
}

.btn-submit i {
  font-size: 0.9rem;
}

.links {
  margin-top: 1.5rem;
  display: flex;
  flex-direction: column;
  gap: 0.75rem;
}

.link {
  color: var(--primary-color);
  text-decoration: none;
  font-size: 0.9rem;
  display: flex;
  align-items: center;
  gap: 0.5rem;
  transition: color 0.2s ease;
}

.link:hover {
  color: var(--secondary-color);
  text-decoration: underline;
}

.link i {
  font-size: 0.8rem;
}

.link.admin {
  color: var(--danger-color);
}

.link.admin:hover {
  color: #e5177e;
}

@media (max-width: 480px) {
  .login-form {
    padding: 1.5rem;
  }

  h1 {
    font-size: 1.5rem;
    margin-bottom: 1.5rem;
  }
}
</style>