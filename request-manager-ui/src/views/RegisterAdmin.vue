<template>
  <div class="register-admin login-container">
    <div class="login-form">
      <h1>Реєстрація адміністратора</h1>
      <form @submit.prevent="registerAdmin">
        <div class="input-group">
          <label for="firstname">Ім'я</label>
          <input
              v-model="firstname"
              type="text"
              id="firstname"
              required
              class="form-input"
              placeholder="Введіть ваше ім'я"
          />
        </div>

        <div class="input-group">
          <label for="lastname">Прізвище</label>
          <input
              v-model="lastname"
              type="text"
              id="lastname"
              required
              class="form-input"
              placeholder="Введіть ваше прізвище"
          />
        </div>

        <div class="input-group">
          <label for="email">Email</label>
          <input
              v-model="email"
              type="email"
              id="email"
              required
              class="form-input"
              placeholder="Введіть ваш email"
          />
        </div>

        <div class="input-group">
          <label for="username">Юзернейм</label>
          <input
              v-model="username"
              type="text"
              id="username"
              required
              class="form-input"
              placeholder="Введіть юзернейм"
          />
        </div>

        <div class="input-group password-wrapper">
          <label for="password">Пароль</label>
          <input
              v-model="password"
              :type="isPasswordVisible ? 'text' : 'password'"
              id="password"
              required
              class="form-input"
              placeholder="Введіть пароль"
          />
          <button
              type="button"
              class="eye-button"
              :class="{ 'active': isPasswordVisible }"
              @click="togglePasswordVisibility"
          >
            <i class="fas" :class="isPasswordVisible ? 'fa-eye-slash' : 'fa-eye'"></i>
          </button>
        </div>

        <div class="input-group password-wrapper">
          <label for="confirmPassword">Підтвердіть пароль</label>
          <input
              v-model="confirmPassword"
              :type="isPasswordVisible ? 'text' : 'password'"
              id="confirmPassword"
              required
              class="form-input"
              placeholder="Підтвердіть пароль"
          />
          <button
              type="button"
              class="eye-button"
              :class="{ 'active': isPasswordVisible }"
              @click="togglePasswordVisibility"
          >
            <i class="fas" :class="isPasswordVisible ? 'fa-eye-slash' : 'fa-eye'"></i>
          </button>
        </div>

        <button type="submit" class="btn-submit">
          <i class="fas fa-user-shield"></i> Зареєструватися як адміністратор
        </button>
      </form>

      <div class="links">
        <router-link to="/login" class="link">
          <i class="fas fa-sign-in-alt"></i> Вхід в систему
        </router-link>
        <router-link to="/register" class="link">
          <i class="fas fa-user-plus"></i> Реєстрація користувача
        </router-link>
      </div>
    </div>
  </div>
</template>

<script>
import { ref } from 'vue';
import { authApi } from '../api';
import { useRouter } from 'vue-router';

export default {
  name: 'RegisterAdmin',
  setup() {
    const router = useRouter();
    const firstname = ref('');
    const lastname = ref('');
    const email = ref('');
    const username = ref('');
    const password = ref('');
    const confirmPassword = ref('');
    const isPasswordVisible = ref(false);

    const togglePasswordVisibility = () => {
      isPasswordVisible.value = !isPasswordVisible.value;
    };

    const registerAdmin = async () => {
      if (password.value !== confirmPassword.value) {
        alert("Паролі не співпадають.");
        return;
      }

      try {
        const response = await authApi.registerAdmin(
            firstname.value,
            lastname.value,
            email.value,
            username.value,
            password.value
        );
        console.log('Адміністратора зареєстровано:', response.data);
        alert('Ви успішно зареєструвались. Вас буде автоматично перенаправлено на сторінку входу.');
        router.push('/login');
      } catch (error) {
        alert('Помилка реєстрації. Спробуйте ще раз.');
      }
    };

    return {
      firstname,
      lastname,
      email,
      username,
      password,
      confirmPassword,
      isPasswordVisible,
      togglePasswordVisibility,
      registerAdmin,
    };
  },
};
</script>

<style scoped>
.register-admin {
  background-color: var(--light-color);
}

.login-container {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 100vh;
  padding: 20px;
}

.login-form {
  background-color: var(--white);
  padding: 2.5rem;
  border-radius: 12px;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.1);
  width: 100%;
  max-width: 500px;
}

h1 {
  text-align: center;
  margin-bottom: 2rem;
  color: var(--danger-color);
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

.btn-submit {
  background-color: var(--danger-color);
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
  margin-top: 1rem;
}

.btn-submit:hover {
  background-color: #e5177e;
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
  justify-content: center;
}

.link:hover {
  color: var(--secondary-color);
  text-decoration: underline;
}

.link i {
  font-size: 0.8rem;
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