<template>
  <div class="register-admin login-container">
    <div class="login-form">
      <h2>Реєстрація адміністратора</h2>
      <form @submit.prevent="registerAdmin">
        <div class="input-group">
          <label for="firstname">Ім'я</label>
          <input v-model="firstname" type="text" id="firstname" required />
        </div>
        <div class="input-group">
          <label for="lastname">Прізвище</label>
          <input v-model="lastname" type="text" id="lastname" required />
        </div>
        <div class="input-group">
          <label for="email">Email</label>
          <input v-model="email" type="email" id="email" required />
        </div>
        <div class="input-group">
          <label for="username">Юзернейм</label>
          <input v-model="username" type="text" id="username" required />
        </div>
        <div class="input-group password-wrapper">
          <label for="password">Пароль</label>
          <input
              v-model="password"
              :type="isPasswordVisible ? 'text' : 'password'"
              id="password"
              required
          />
          <button type="button" class="eye-button" @click="togglePasswordVisibility">👁️</button>
        </div>
        <div class="input-group password-wrapper">
          <label for="confirmPassword">Підтвердіть пароль</label>
          <input
              v-model="confirmPassword"
              :type="isPasswordVisible ? 'text' : 'password'"
              id="confirmPassword"
              required
          />
          <button type="button" class="eye-button" @click="togglePasswordVisibility">👁️</button>
        </div>
        <button type="submit" class="btn-submit">Зареєструватися як адміністратор</button>
      </form>
      <div class="links">
        <p><router-link to="/login">Вхід в систему</router-link></p>
        <p><router-link to="/register">Реєстрація користувача</router-link></p>
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

      console.log('Registering with data:', {
        first_name: firstname.value,
        last_name: lastname.value,
        email: email.value,
        username: username.value,
        password: password.value,
      });

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
html, body {
  height: 100%;
  margin: 0;
  padding: 0;
}

.login-container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100%;
  background-color: #f4f7fc;
  overflow: auto; /* Ensure scrollability */
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
  top: 65%;
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

.error-message {
  color: red;
  margin-top: 10px;
}

.links {
  text-align: center;
  margin-top: 1rem;
}

.links p {
  margin-top: 1rem;
  font-size: 0.9rem;
}

.links a {
  color: #007bff;
  text-decoration: none;
}

.links a:hover {
  text-decoration: underline;
}
</style>
