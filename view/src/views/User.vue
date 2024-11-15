<template>
    <div>
      <h1>用户个人信息</h1>
      <div v-if="userInfo">
        <p>用户名: {{ userInfo.uname }}</p>
        <p>权限: {{ userInfo.auth }}</p>
      </div>
      <div v-else>
        <p>请先登录</p>
        </div>
    </div>
  </template>
  
  <script>

  import { checkUserLoggedIn } from '../logic/check_jwt';
  export default {
    name: 'User',
    beforeRouteEnter(to, from, next) {
      // 假设您有一个方法来检查用户是否已登录
      const isLoggedIn = checkUserLoggedIn(); // 替换为实际的登录检查逻辑
  
      if (isLoggedIn) {
        GetInfo();
        next(); // 用户已登录，继续访问
      } else {
        next({ name: 'Login' }); // 用户未登录，跳转到登录页面
      }
    }
  }

  async function GetInfo(){
    try {
      const response = await fetch(`${Url}/user/info`, {
        headers: {
          'token': `${localStorage.getItem('jwt')}`
        }
      });
      const data = await response.json();
      console.log(data);
      this.userInfo = data;

    } catch (error) {
      console.error('Error fetching user info:', error);
    }
  }
  
  // 模拟的登录检查函数

  </script>
  
  <style>
  /* 添加样式 */
  </style>