<template>
    <div>
        <h1>Login Page</h1>
        <form @submit.prevent="handleLogin">
            <input v-model="uname" placeholder="用户名" required />
            <input type="password" v-model="passwd" placeholder="密码" required />
            <button type="submit">登录</button>
            <button @click="goToRegister" type="button">注册</button> <!-- 添加注册按钮 -->
        </form>
    </div>
</template>

<script>
import { Url } from '../config/config';
export default {

    data() {
        return {
            uname: '',
            passwd: ''
        };
    },
    methods: {
        async handleLogin() {
            try {
                const response = await fetch(`${Url}/user/login`, {
                    method: 'POST', 
                    headers: {
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify({
                        uname: this.uname,
                        passwd: this.passwd
                    })
                });

                if (!response.ok) {
                    throw new Error('登录失败'); // 处理登录失败
                }

                const data = await response.json();
                console.log(data.token);
                localStorage.setItem('jwt',data.token); // 存储 JWT
                // 登录成功后的处理逻辑，例如跳转到用户页面
                this.$router.push({ name: 'User' });
            } catch (error) {
                console.error(error); // 处理错误
            }
        },
        goToRegister() {
            this.$router.push({ name: 'Register' }); // 跳转到注册页面
        }
    }
};
</script>