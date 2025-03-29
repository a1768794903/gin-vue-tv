<template>
  <div id="app">
    <!-- 确保有路由出口 -->
    <router-view></router-view>
  </div>
</template>

<script>
export default {
  name: 'HomePage',
  data() {
    return {
      content: 'home' // 默认展示的内容
    };
  },
  methods: {
    changeContent(newContent) {
      this.content = newContent;
    }
  }
}
</script>

<style scoped>
/* 样式部分 */
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

.navbar ul {
  list-style-type: none;
  padding: 0;
}

.navbar li {
  display: inline;
  margin-right: 10px;
}

.main-content img {
  width: 100%;
  height: auto;
}

.footer {
  margin-top: 20px;
  padding: 10px;
  background-color: #f8f9fa;
}
</style>