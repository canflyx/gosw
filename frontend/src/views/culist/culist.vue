<template>
  <page-search
    :searchConfig="searchConfig"
    @queryBtnClick="handleQueryClick"
    @resetBtnClick="handleRestClick"
    :permission="permission"
  ></page-search>
  <page-content
    :pageName="pageName"
    :contentConfig="contentConfig"
    ref="pageContentRef"
    :permission="permission"
  >
    <template #log="scope" >
      <div  class="text-wrap"> <span>{{scope.row.log}}</span></div>
    </template>

    <template #UpdatedAt="scope">
      <span>{{ $filters.formatTime(scope.row.UpdatedAt) }}</span>
    </template>
  </page-content>
</template>

<script setup lang="ts">
import {} from 'vue'
import { usePageSearch } from '@/hooks/use-page-search'
import PageContent from '@/components/page-content'
import { contentConfig } from './config/content.config'
import PageSearch from '@/components/page-search'
import { searchConfig } from './config/search.config'
import { IPermission } from '@/base-ui/table/types'

const pageName = 'culog'

const permission: IPermission = {
  isCreate: false,
  isQuery: true,
  isUpdate: false,
  isDelete: false,
  queryTitle: '搜索'
}

// 抽离到 hook中

const [pageContentRef, handleRestClick, handleQueryClick] = usePageSearch()
</script>

<style lang="scss" scoped>
.text-wrap {
 /* white-space: pre-line	;
  text-align: left;
  width: 100%;  仅为示例，可以根据需要设置宽度 */
}
</style>
