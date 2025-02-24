<template>
  <div class="login-container">
    <div class="login-form">
      <h2>Вхід в систему</h2>
      <form @submit.prevent="handleSubmit">
        <div class="input-group">
          <label for="username">Логін</label>
          <input
              type="text"
              v-model="username"
              id="username"
              placeholder="Введіть ваш логін"
              required
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
            />
            <button type="button" @click="togglePasswordVisibility" class="eye-button">
              👁️
            </button>
          </div>
        </div>

        <div class="actions">
          <button type="submit" class="btn-submit">Увійти</button>
          <p class="forgot-password">
            <router-link to="/register">Зареєструватися</router-link>
          </p>
          <p class="admin-registration">
            <router-link to="/registerAdmin">Зареєструватися як адміністратор</router-link>
          </p>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import { authApi } from '../api';

const username = ref('');
const password = ref('');
const isPasswordVisible = ref(false);
const router = useRouter();

const togglePasswordVisibility = () => {
  isPasswordVisible.value = !isPasswordVisible.value;
};

const handleSubmit = async () => {
  try {
    const response = await authApi.login(username.value, password.value);
    localStorage.setItem('token', response.data.token);
    localStorage.setItem('role', response.data.role);
    router.push('/dashboard');
  } catch (error) {
    alert('Увійти не вдалося');
  }
};
</script>

<style scoped>
.login-container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  background-color: #f4f7fc;
}

.login-form {
  background-color: #fff;
  padding: 2rem;
  border-radius: 8px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
  width: 400px;
  max-width: 100%;
}

h2 {
  text-align: center;
  margin-bottom: 1.5rem;
  color: #333;
}

.input-group {
  margin-bottom: 1rem;
}

.input-group label {
  display: block;
  font-weight: 500;
  margin-bottom: 0.5rem;
}

.input-group input {
  width: 100%;
  padding: 0.8rem;
  font-size: 1rem;
  border-radius: 6px;
  border: 1px solid #ccc;
  box-sizing: border-box;
}

.password-wrapper {
  position: relative;
}

.eye-button {
  position: absolute;
  top: 50%;
  right: 10px;
  transform: translateY(-50%);
  background: none;
  border: none;
  font-size: 1.2rem;
  cursor: pointer;
}

.actions {
  text-align: center;
}

.btn-submit {
  background-color: #007bff;
  color: white;
  padding: 0.8rem 1.5rem;
  border: none;
  border-radius: 6px;
  cursor: pointer;
  font-size: 1.1rem;
  width: 100%;
  transition: background-color 0.3s ease;
}

.btn-submit:hover {
  background-color: #0056b3;
}

.forgot-password {
  margin-top: 1rem;
  font-size: 0.9rem;
}

.forgot-password a {
  color: #007bff;
  text-decoration: none;
}

.forgot-password a:hover {
  text-decoration: underline;
}

.admin-registration {
  margin-top: 1rem;
  font-size: 0.9rem;
}

.admin-registration a {
  color: #007bff;
  text-decoration: none;
}

.admin-registration a:hover {
  text-decoration: underline;
}
</style>
