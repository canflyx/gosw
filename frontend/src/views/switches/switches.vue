<template>
  <page-search
    :searchConfig="searchConfig"
    @queryBtnClick="handleQueryClick"
    @resetBtnClick="handleRestClick"
  ></page-search>
  <page-dialog
    :pageName="pageName"
    :dialogConfig="dialogConfig"
    ref="pageDialogRef"
    :defaultInfo="defaultInfo"
  ></page-dialog>
  <page-content
    :pageName="pageName"
    :contentConfig="contentConfig"
    @addBtnClick="handleAddClick"
    @editBtnClick="handleEditClick"
    :permission="permission"
    ref="pageContentRef"
  >
    <template #headHandler>
      <el-button
        type="primary"
        v-if="permission?.isCreate"
        @click="handleScanItem()"
        :icon="Search"
        >扫描交换机</el-button
      >
    </template>
    <template #UpdatedAt="scope">
      <span>{{ $filters.formatTime(scope.row.UpdatedAt) }}</span>
    </template>
    <template #iscore="scope">
      <span v-if="scope.row.iscore === 1">核心</span>
      <span v-else>接入</span>
    </template>
    <template #status="scope">
      <span v-if="scope.row.status === 1">禁用</span>
      <span v-else>正常</span>
    </template>
  </page-content>
</template>

<script setup lang="ts">
import {} from 'vue'
import PageDialog from '@/components/page-dialog'
import { dialogConfig } from './config/dialog.config'
import { usePageDialog } from '@/hooks/use-page-dialog'
import { usePageSearch } from '@/hooks/use-page-search'
import PageContent from '@/components/page-content'
import { contentConfig } from './config/content.config'
import PageSearch from '@/components/page-search'
import { searchConfig } from './config/search.config'
import { IPermission } from '@/base-ui/table/types'
import { Search } from '@element-plus/icons-vue'
import { mainStore } from '@/store'

const pageName = 'switches'
const permission: IPermission = {
  isCreate: true,
  isQuery: true,
  isUpdate: true,
  isDelete: true
}
const store = mainStore()
const handleScanItem = () => {
  store.scanPageDataAction()
}
const newCallback = () => {
  const passwordItem = dialogConfig.formItem.find(
    (item) => item.field === 'password'
  )
  passwordItem!.isHidden = false
}
const editCallback = () => {
  const passwordItem = dialogConfig.formItem.find(
    (item) => item.field === 'password'
  )
  passwordItem!.isHidden = false
}

// 抽离到 hook中

const [pageDialogRef, defaultInfo, handleAddClick, handleEditClick] =
  usePageDialog(newCallback, editCallback)
const [pageContentRef, handleRestClick, handleQueryClick] = usePageSearch()
</script>

<style lang="scss" scoped></style>
