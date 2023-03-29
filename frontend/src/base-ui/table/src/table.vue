<template>
  <div class="hy-table">
    <div class="header">
      <slot name="header">
        <div class="title">{{ title }}</div>
        <div class="handle">
          <slot name="headHandler"></slot>
        </div>
      </slot>
    </div>
    <el-table
      :data="listData"
      border
      class="show_table"
      style="width: 100%"
      @selection-change="handleSelectionChange"
      v-bind="childrenProps"
    >
      <el-table-column v-if="selectColum" type="selection" width="55" />
      <el-table-column
        v-if="showIndexColum"
        label="序号"
        align="center"
        type="index"
        width="60"
      ></el-table-column>

      <template v-for="item in propList" :key="item.prop">
        <el-table-column v-bind="item" align="center" show-overflow-tooltip>
          <!-- 此处默认是有一个 #default的插槽 scope.row 取行，item.prop 相当于 name-->
          <template #default="scope">
            <!-- 为了不把 slot 写死，在propList中增加一个 slotName来定义哪些是需要格式化值的。
          通过 slot 属性，将 row行传回到上一层     -->
            <slot :name="item.slotName" :row="scope.row">
              {{ scope.row[item.prop] }}</slot
            >
          </template>
        </el-table-column>
      </template>
    </el-table>
    <div class="footer" v-if="showFooter">
      <slot name="footer">
        <el-pagination
          :current-page="page.currentPage"
          :page-size="page.pageSize"
          :page-sizes="[10, 20, 50, 100]"
          layout="total, sizes, prev, pager, next, jumper"
          :total="listCount"
          @size-change="handleSizeChange"
          @current-change="handleCurrentChange"
        />
      </slot>
    </div>
  </div>
</template>

<script setup lang="ts">
import { defineEmits, PropType } from 'vue'

const props = defineProps({
  title: {
    type: String,
    default: ''
  },
  listCount: {
    type: Number,
    default: 0
  },
  listData: {
    type: Array,
    required: true
  },
  propList: {
    type: Array as PropType<any[]>,
    required: true
  },
  page: {
    type: Object,
    default: () => ({ currentPage: 0, pageSize: 10 })
  },
  showIndexColum: {
    type: Boolean,
    default: false
  },
  selectColum: {
    type: Boolean,
    default: false
  },
  childrenProps: {
    type: Object,
    default: () => ({})
  },
  showFooter: {
    type: Boolean,
    default: true
  }
})

const emits = defineEmits<{
  (e: 'selectionChange', data: any): void
  (e: 'update:page', data: any): void
}>()
const handleSizeChange = (pageSize: number) => {
  emits('update:page', { ...props.page, pageSize })
}
const handleCurrentChange = (currentPage: number) => {
  emits('update:page', { ...props.page, currentPage })
}
const handleSelectionChange = (value: object) => {
  emits('selectionChange', value)
}
</script>

<style lang="scss" scoped>
.title {
  font-size: 20px;
  font-weight: 700;
  margin-bottom: 30px;
}
.footer {
  margin-top: 15px;
  display: flex;
  justify-content: right;
}
.handle {
  display: flex;
  justify-content: right;
}
.show_table {
  position: relative;
  width: 100%;
  height: 450px;
  overflow: auto;
}
</style>
