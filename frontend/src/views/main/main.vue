<template>
  <div class="common-layout">
    <el-container>
      <el-aside>
        <AppMenu></AppMenu>
      </el-aside>
      <el-container>
        <div class="app-main">
          <!-- 子路由出口 -->
          <router-view></router-view>
        </div>
      </el-container>
    </el-container>
  </div>
</template>
<script setup lang="ts">
import AppMenu from './nvmenu/index.vue'
// import AppMain from './components/AppMain/index.vue'

</script>
<script lang="ts">
export default {
  name: 'MainIt'
}
</script>
<style scoped>
.common-layout {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
}
.el-aside {
  width: auto;

  box-sizing: border-box;
  height: 100%;
}
.app-main {
  min-height: 100%;
  width: 100%;
  position: relative;
  overflow: hidden;
  padding: 20px 20px 20px 20px;
  box-sizing: border-box;
}
.layout-container-demo .el-main {
  padding: 0;
}
</style>
