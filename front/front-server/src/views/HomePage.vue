<template>
  <div id="app">
    <!-- 导航栏 -->
    <nav class="navbar">
      <ul>
        <li><a href="#" @click.prevent="changeContent('home')">主页</a></li>
        <li><a href="#" @click.prevent="changeContent('tvbox')">TVBOX接口</a></li>
        <li><a href="#" @click.prevent="changeContent('recommend')">相关推荐</a></li>
      </ul>
    </nav>

    <!-- 主要内容区域 -->
    <main class="main-content">
      <div v-if="content === 'home'" class="image-container">
        <!-- <img src="./assets/home.jpg" alt="Homepage Image"/> -->
        <p>「没有一个冬天不可逾越 没有一个春天不会来临」</p>
      </div>
      <div v-if="content === 'tvbox'" class="image-container">
        <!-- <img src="./assets/tvbox.jpg" alt="TVBOX Image"/> -->
        <p>这里是关于TVBOX接口的内容。</p>
      </div>
      <div v-if="content === 'recommend'" class="image-container">
        <!-- <img src="./assets/recommend.jpg" alt="Recommendations Image"/> -->
        <p>这里是一些相关推荐的内容。</p>
      </div>
    </main>

    <!-- 底部信息 -->
    <footer class="footer">
      <p>&copy; 2025 [你的名字]</p>
    </footer>
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