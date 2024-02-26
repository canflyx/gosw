<template>
  <page-search
    :searchConfig="searchConfig"
    :permission="permission"
    @queryBtnClick="handleQueryClick"
    @resetBtnClick="handleRestClick"
  >
    <template #other>
      <el-card class="tree-card">
        <template #header>
          <el-input v-model="filterText" placeholder="筛选" />
        </template>
        <div class="tree">
          <el-tree
            ref="treeRef"
            class="filter-tree"
            :data="store.swList"
            node-key="id"
            default-expand-all
            :filter-node-method="filterNode"
            show-checkbox
            @check-change="handleCheckChange"
          />
        </div>
      </el-card>
      <el-card class="tag-card">
        <template #header>
          <span>已选择列表</span>
        </template>
        <el-tag
          v-for="(tag, index) in tags"
          :key="index"
          style="margin: 10px 10px"
          closable
          @close="handleClose(tag)"
        >
          {{ tag.label }}
        </el-tag>
      </el-card>
      <div class="jsonedit">
        <Vue3JsonEditor
          v-model="json"
          :show-btns="false"
          :expandedOnStart="true"
          mode="code"
          @json-change="onJsonChange"
        />
      </div>
    </template>
  </page-search>
</template>

<script setup lang="ts">
import { ref, watch, onMounted } from 'vue'
import PageSearch from '@/components/page-search'
import { searchConfig } from './config/search.config'
import { mainStore } from '@/store'
import { ElTree, ElMessage } from 'element-plus'
import { TreeNodeData } from 'element-plus/es/components/tree/src/tree.type'
import { IPermission } from '@/base-ui/table/types'
import { Vue3JsonEditor } from 'vue3-json-editor'

const permission: IPermission = {
  isCreate: true,
  isQuery: true,
  isUpdate: true,
  isDelete: true,
  queryTitle: '扫描交换机'
}

// 树型处理
interface Tree {
  id: number
  label: string
  children?: Tree[]
}

const filterText = ref('')
const treeRef = ref<InstanceType<typeof ElTree>>()
const tags = ref<Tree[]>([])
const handleCheckChange = () => {
  //   去掉 children 根结点
  const tag = treeRef.value!.getCheckedNodes() as Tree[]
  tags.value = tag.filter((item) => item.children == null)
}
// 监听过滤条件
watch(filterText, (val) => {
  treeRef.value!.filter(val)
})
// 监听 tags 变化
watch(tags, () => {
  store.tags = tags.value
})
// tags 关闭事件，去掉 tree 选择
const handleClose = (value: any) => {
  treeRef.value!.setChecked(value.id, false, false)
}
// 树型过滤
const filterNode = (value: string, data: TreeNodeData) => {
  if (!value) return true
  return data.label.includes(value)
}

const store = mainStore()
onMounted(() => {
  store.getInitial()
})

const handleQueryClick = () => {
  // console.log(JSON.parse(queryInfo))
  if (tags.value.length < 1) {
    ElMessage({ showClose: true, message: '没有选择交换机', type: 'error' })
    return
  }
  store.scanSwDataAction()
}
const handleRestClick = () => {
  treeRef.value!.setCheckedKeys([], false)
  filterText.value = ''
}

// json 编辑器处理
const json = ref([{ cmd: 'show users', flag: ']' }])

const onJsonChange = (value: []) => {
  if (!Array.isArray(value) || value.length < 1) return
  store.readCmd = value
}
</script>

<style lang="scss" scoped>
.tag-card {
  width: 40%;
}
.tree-card {
  height: 400px;
  width: 40%;
  overflow-y: auto;
}
.jsonedit {
  width: 81%;
}

.tree {
  display: flex;
  justify-content: left;
}
</style>
