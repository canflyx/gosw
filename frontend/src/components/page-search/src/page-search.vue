<template>
  <div class="page-search">
    <h-form v-bind="props.searchConfig" v-model="formData">
      <template #header>
        <span v-if="props.searchConfig.title">{{
          props.searchConfig.title
        }}</span>
        <span v-else> 高级检索</span>
      </template>
      <template #other>
        <slot name="other"></slot>
      </template>
      <template #footer>
        <div class="handle-btns">
          <slot name="otherHandler"></slot>
          <el-button
            type="primary"
            @click="handleQueryClick"
            v-if="permission?.isQuery"
            :icon="Search"
            plain
          >
            <span v-if="permission.queryTitle">{{
              permission.queryTitle
            }}</span>
            <span v-else>搜索</span></el-button
          >
          <el-button
            type="danger"
            :icon="RefreshRight"
            @click="handleResetClick"
            plain
            >重置</el-button
          >
        </div>
      </template>
    </h-form>
  </div>
</template>

<script setup lang="ts">
import { ref, PropType } from 'vue'
import { Search, RefreshRight } from '@element-plus/icons-vue'
import type { IPermission } from '@/base-ui/table/types'

import HForm from '@/base-ui/form'

const props = defineProps({
  searchConfig: {
    type: Object,
    required: true
  },
  permission: { type: Object as PropType<IPermission> }
})
// 搜索由 formItem 决定，可以直接读取 field字段 用作搜索
const formItems = props.searchConfig?.formItem ?? []
const formOriginData: any = {}
for (const item of formItems) {
  formOriginData[item.field] = ''
}

const formData = ref(formOriginData)
const emits = defineEmits<{
  (e: 'resetBtnClick'): void
  (e: 'queryBtnClick', data: any): void
}>()
// reset 使用 :model-value 和 @update:modelValue 来双向绑定
const handleResetClick = () => {
  formData.value = formOriginData
  emits('resetBtnClick')
}
const handleQueryClick = () => {
  emits('queryBtnClick', formData.value)
}
</script>

<style lang="scss" scoped>
.handle-btns {
  display: flex;
  justify-content: right;
  padding: 10px 10px 10px 0px;
}
</style>
