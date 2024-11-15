function checkUserLoggedIn() {
    const token = localStorage.getItem('jwt'); // 从 localStorage 中获取 JWT
    if (!token) return false; // 如果没有 token，用户未登录

    try {
      // 检查 JWT 是否过期
      const payload = JSON.parse(atob(token.split('.')[1])); // 解码 JWT
      const isExpired = payload.exp * 1000 < Date.now(); // 检查过期时间

      return !isExpired; // 返回用户是否已登录
    } catch (error) {
      return false; // 解码失败，返回未登录
    }
  }

  function checkAdmin(){
    const token = localStorage.getItem('jwt'); // 从 localStorage 中获取 JWT
    if (!token) return false; // 如果没有 token，用户未登录
    const payload = JSON.parse(atob(token.split('.')[1])); // 解码 JWT
    return payload.auth === true; // 返回用户是否为管理员
  }

  export { checkUserLoggedIn, checkAdmin };