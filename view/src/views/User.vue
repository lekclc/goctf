<template>
  <div>
    <h1>用户个人信息</h1>
    <div v-if="userInfo">
      <p>用户名: {{ userInfo.name }}</p>
      <p>权限: {{ userInfo.auth }}</p>
      <div v-if="userInfo.team">
        <p>队伍: {{ userInfo.team }}</p>
      </div>
      <button @click="logout">登出</button> <!-- 添加登出按钮 -->
      <br>
      <button @click="update">修改信息</button> <!-- 添加登出按钮 -->
    </div>
    <div v-else>
      <p>请先登录</p>
    </div>
  </div>
  <div v-if="showModalUpdate" class="modal">
    <div class="modal-content">
      <span class="close" @click="closeModal">&times;</span>
      <h2>修改信息</h2>
      <form @submit.prevent="update_submit">
        <div>
          <label>用户名:</label>
          {{ this.userInfo.name }}
        </div>
        <div>
          <label>原密码:</label>
          <input type="password" v-model="oldPasswd" placeholder="" required />
        </div>
        <div>
          <label>新密码:</label>
          <input type="password" v-model="newPasswd" placeholder="" required />
        </div>
        <button type="submit">提交</button>
        <button type="button" @click="closeModal">取消</button>
      </form>
    </div>
  </div>
</template>

<script>
import { checkUserLoggedIn } from '../logic/check_jwt'; // 假设您有一个检查登录状态的函数
import { Url } from '../config/config'; // 假设您有一个配置文件

export default {
  name: 'User',
  data() {
    return {
      userInfo: {},
      showModalUpdate: false,
    };
  },
  async created() {
    // 在组件创建时获取用户信息
    if (checkUserLoggedIn()) {
      await this.getUserInfo();
    } else {
      this.$router.push('/login');
    }
  },
  methods: {
    async getUserInfo() {
      const name = localStorage.getItem('name');
      const token = localStorage.getItem('jwt');


      if (!name || !token) {
        console.error('Name or token is missing');
        return; // 如果缺少必要的值，提前返回
      }

      const formData = new URLSearchParams();
      formData.append('name', name);
      formData.append('token', token); //token 是用于验证的

      try {
        const response = await fetch(`${Url}/user/info`, {
          method: 'POST',
          headers: {
            'Content-Type': 'application/x-www-form-urlencoded'
          },
          body: formData
        });

        if (!response.ok) {
          throw new Error('获取用户信息失败');
        }

        const data = await response.json();
        this.userInfo = data;
        if (data.admin == true) {
          this.userInfo.auth = "管理员"
        } else {
          this.userInfo.auth = "用户"
        }
      } catch (error) {
        console.error('Fetch error:', error);
      }
    },
    logout() {
      // 清除 localStorage 中的用户信息
      localStorage.removeItem('name');
      localStorage.removeItem('jwt');
      this.userInfo = null; // 清空用户信息
      this.$router.push('/login'); // 跳转到登录页面
    },
    update() {
      this.showModalUpdate = true;
    },
    closeModal(){
      this.showModalUpdate = false;
    },
    async update_submit() {
      const formData = new FormData();
      formData.append('name', localStorage.getItem('name'));
      formData.append('token', localStorage.getItem('jwt'));
      formData.append('old_passwd', this.oldPasswd);
      formData.append('new_passwd', this.newPasswd);
      try {
        const response = await fetch(`${Url}/user/update`, {
          method: 'POST',
          body: formData
        });
        const res = response.json();

        if (response.ok) {
          alert('修改成功');
          this.logout();
        } else {
          alert('修改失败');
        }

      } catch (error) {
        console.error('Fetch error:', error);
      }
    }
  }
};
</script>

<style>
/* 添加样式 */
</style>