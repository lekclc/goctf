<template>
    <div id="app" class="container">
      <nav class="sidebar">
        <ul>
          <li><router-link to="/">Home</router-link></li>
          <li><router-link to="/about">About</router-link></li>
          <li><router-link to="/game">Game</router-link></li>
          <li><router-link to="/team">Team</router-link></li>
          <li><router-link to="/user">User</router-link></li>

          <!-- 可以在这里添加更多的导航链接 -->
        </ul>
      </nav>
      <main class="content">
        <router-view />
      </main>
    </div>
  </template>
  
  <script>
  export default {
    name: 'App'
  }
  </script>
  
  <style>
  * {
  margin: 0; /* 去掉所有元素的外边距 */
  padding: 0; /* 去掉所有元素的内边距 */
  box-sizing: border-box; /* 包括边框和内边距在内的宽高计算 */
}

 .container {
   display: flex;
   height: 100vh; /* 使容器高度占满整个视口 */
 }
  .sidebar {
    width: 60px; /* 设置导航栏的固定宽度 */
    max-width: 250px; /* 设置导航栏的最大宽度 */
    background-color: #f4f4f4; /* 设置导航栏的背景色 */
    padding: 5px; /* 设置内边距 */
  }
  
  .sidebar ul {
    list-style-type: none; /* 去掉列表的默认样式 */
    padding: 0; /* 去掉内边距 */
  }
  
  .sidebar li {
    margin: 10px 0; /* 设置列表项的上下间距 */
  }
  
  .sidebar a {
    text-decoration: none; /* 去掉链接的下划线 */
    color: #333; /* 设置链接的颜色 */
  }
  
  .sidebar a:hover {
    color: #007BFF; /* 设置链接的悬停颜色 */
  }
  
  .content {
    flex: 1; /* 使内容区域占据剩余空间 */
    padding: 20px; /* 设置内容区域的内边距 */
  }
  </style>