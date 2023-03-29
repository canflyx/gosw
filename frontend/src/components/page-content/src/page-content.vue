<template>
  <hf-table
    :listData="dataList"
    v-bind="props.contentConfig"
    :listCount="dataCount"
    v-model:page="pageInfo"
    @selectionChange="selectionChange"
  >
    <!-- v-model如果不指定名称默认对应的是 modeValue update:modeValue,用名字后
  就是 page update:page
  -->
    <!-- header 插槽 -->

    <template #headHandler>
      <slot name="headHandler"></slot>
      <el-button
        type="primary"
        v-if="permission?.isCreate"
        @click="handleAddItem"
        :icon="Plus"
        >新增</el-button
      >
    </template>

    <!-- 表中间的插槽 -->
    <!-- 使用插槽来修改值 或 格式化数值 -->
    <template #status="scope">
      <el-button
        plain
        size="small"
        :type="scope.row.enable ? 'success' : 'danger'"
        >{{ scope.row.enable ? '启用' : '禁用' }}</el-button
      >
    </template>
    <template #createAt="scope">
      <!-- $filters 为全局格式化函数 -->
      <span>{{ $filters.formatTime(scope.row.createAt) }}</span>
    </template>
    <template #updateAt="scope">
      <span>{{ $filters.formatTime(scope.row.updateAt) }}</span>
    </template>
    <!-- 后面的操作列 -->
    <template #handle="scope">
      <el-button
        type="primary"
        v-if="permission?.isUpdate"
        @click="handleEdit(scope.row)"
        text
        >编辑</el-button
      >

      <el-popconfirm
        title="确认删除吗?"
        width="220"
        confirm-button-text="OK"
        cancel-button-text="No, Thanks"
        icon-color="#626AEF"
        v-if="permission?.isDelete"
        @confirm="handleDelete(scope.row)"
      >
        <template #reference>
          <el-button type="danger" text>删除</el-button></template
        >
      </el-popconfirm>
    </template>
    <template
      v-for="item in otherPropSlots"
      :key="item.slotName"
      #[item.slotName]="scope"
    >
      <template v-if="item.slotName">
        <slot :name="item.slotName" :row="scope.row"></slot>
      </template>
    </template>
  </hf-table>
</template>

<script setup lang="ts">
import { computed, defineProps, PropType, ref, watch } from 'vue'
import HfTable from '@/base-ui/table'
import type { IContent, IPermission } from '@/base-ui/table/types'

import { Plus } from '@element-plus/icons-vue'
import { mainStore } from '@/store/'

const props = defineProps({
  contentConfig: { type: Object as PropType<IContent>, required: true },
  pageName: { type: String, required: true },
  permission: { type: Object as PropType<IPermission> }
})
// 按钮权限

// 双向绑定 pageInfo
const pageInfo = ref({
  pageSize: 10,
  currentPage: 1
})
watch(pageInfo, () => getPageData())
// 1.发送网络请求
const store = mainStore()

// 获取内容的 url和参数
const getPageData = (queryInfo: any = {}) => {
  if (!props.permission?.isQuery) return
  store.getPageListAction({
    pageName: props.pageName,
    queryInfo: {
      page_number: (pageInfo.value.currentPage - 1) * pageInfo.value.pageSize,
      page_size: pageInfo.value.pageSize,
      kws: { ...queryInfo }
    }
  })
}
getPageData()

// 2 从pinia 中获取数据
let dataList: any
let dataCount: any
switch (props.pageName) {
  case 'maclist':
    dataList = computed(() => store.arpList)
    dataCount = computed(() => store.arpCount)
    break
  case 'switches':
    dataList = computed(() => store.switchesList)
    dataCount = computed(() => store.switchesCount)
    break
}

// 3.获取其它动态插槽名称(一般在显示页面中自己定义)
const otherPropSlots = props.contentConfig?.propList.filter((item: any) => {
  if (item.slotName === 'createAt') return false
  if (item.slotName === 'updateAt') return false
  if (item.slotName === 'handle') return false
  return true
})

const selectionChange = (value: any) => {
  store.selectItem = []
  for (const item of value) {
    store.selectItem.push(item.ID)
  }
}

// 4.增删改
const handleDelete = (item: any) => {
  store.deletePageDataAction({
    pageName: props.pageName,
    id: item.ID
  })
  getPageData()
}
const emits = defineEmits<{
  (e: 'addBtnClick'): void
  (e: 'editBtnClick', data: any): void
}>()

const handleAddItem = () => {
  emits('addBtnClick')
}
const handleEdit = (item: any) => {
  emits('editBtnClick', item)
}
defineExpose({
  getPageData
})
</script>

<style lang="scss" scoped></style>
