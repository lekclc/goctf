<script setup>
import { RouterLink, RouterView } from 'vue-router'
</script>

<template>
  <div class="fullscreen-div">
    <div class="sidebar">
      <router-link to="/" class="sidebar-item"><i class="icon-home"></i>HOME</router-link>
      <router-link to="/doc" class="sidebar-item"><i class="icon-document"></i>DOC</router-link>
      <router-link to="/game" class="sidebar-item"><i class="icon-flag"></i>GAME</router-link>
      <router-link to="/info" class="sidebar-item"><i class="icon-user"></i>INFO</router-link>
      <router-link to="/user" class="sidebar-item"><i class="icon-info"></i>USER</router-link>
      <router-link to="/setting" class="sidebar-item"><i class="icon-search"></i>SETTING</router-link>
      <router-link to="/control" class="sidebar-item"><i class="icon-settings"></i>管理</router-link>
    </div>
    <div class="main-content">
      <RouterView />
    </div>
  </div>
</template>

<script>
export default {
  name: 'FullscreenDiv'
}
</script>

<style scoped>
html, body {
  margin: 0;
  padding: 0;
  height: 100%;
  overflow: hidden;
}

.fullscreen-div {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: #1e1e1e;
  color: #ffffff;
  display: flex;
}

.sidebar {
  width: 60px;
  background-color: #66ccff;
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 20px 0;
  color: #000000;
}

.sidebar-item {
  margin-bottom: 20px;
  cursor: pointer;
}

.sidebar-item i {
  font-size: 24px;
  color: #ffffff;
}

.sidebar-item:hover i {
  color: #ffffff;
}

.main-content {
  flex-grow: 1;
  padding: 20px;
  overflow-y: auto;
}
</style>