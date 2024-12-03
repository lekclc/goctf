<template>
    <div>
      <h1>注册页面</h1>
      <form @submit.prevent="submitForm">
        <div>
          <label for="username">用户名:</label>
          <input type="text" v-model="username" id="username" required />
        </div>
        <div>
          <label for="password">密码:</label>
          <input type="password" v-model="password" id="password" required />
        </div>
        <div>
          <label for="confirmPassword">确认密码:</label>
          <input type="password" v-model="confirmPassword" id="confirmPassword" required />
        </div>
        <button type="submit">注册</button>
      </form>
      <div v-if="errorMessage" class="error">{{ errorMessage }}</div>
      <div v-if="successMessage" class="success">{{ successMessage }}</div>
    </div>
  </template>
  
  <script>
  import { Url  } from '@/config/config';
  export default {
    data() {
      return {
        username: '',
        password: '',
        confirmPassword: '',
        errorMessage: '',
        successMessage: ''
      };
    },
    methods: {
      async submitForm() {
        // 清空之前的消息
        this.errorMessage = '';
        this.successMessage = '';
  
        // 检查密码是否匹配
        if (this.password !== this.confirmPassword) {
          this.errorMessage = '密码不匹配';
          return;
        }
  
        const formData = new URLSearchParams();
        formData.append('name', this.username);
        formData.append('passwd', this.password);
  
        try {
          const response = await fetch(`${Url}/user/register`, {
            method: 'POST',
            headers: {
              'Content-Type': 'application/x-www-form-urlencoded'
            },
            body: formData
          });
  
          if (!response.ok) {
            const errorData = await response.json();
            throw new Error(errorData.message || '注册失败');
          }
  
          this.successMessage = '注册成功！请登录。';
          // 清空输入框
          this.username = '';
          this.password = '';
          this.confirmPassword = '';
          this.$router.push('login');
        } catch (error) {
          this.errorMessage = error.message;
        }
      }
    }
  };
  </script>
  
  <style>
  .error {
    color: red;
  }
  .success {
    color: green;
  }
  </style>