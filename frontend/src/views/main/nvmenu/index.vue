<template>
  <div>
    <el-radio-group v-model="isCollapse" style="margin-bottom: 20px">
      <el-radio-button v-if="isCollapse" :label="false">打开</el-radio-button>
      <el-radio-button v-else :label="true">隐藏</el-radio-button>
    </el-radio-group>
    <el-menu
      class="el-menu-vertical-demo"
      default-active="2"
      :collapse="isCollapse"
      router
    >
      <el-menu-item index="/arplist">
        <el-icon><icon-menu /></el-icon>
        <template #title>主机ARP查询</template>
      </el-menu-item>
      <el-menu-item index="/switches">
        <el-icon><document /></el-icon>
        <template #title>交换机管理</template>
      </el-menu-item>
    </el-menu>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { Document, Menu as IconMenu } from '@element-plus/icons-vue'
const isCollapse = ref(true)
</script>

<style lang="scss" scoped>
.el-menu {
  border-right: none;
}
</style>
