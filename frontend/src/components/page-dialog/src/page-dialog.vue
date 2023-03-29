<template>
  <el-dialog v-model="dialogVisible" :title="title()" center destroy-on-close>
    <h-form v-bind="props.dialogConfig" v-model="formData" ref="content">
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="dialogVisible = false">取消</el-button>
          <el-button type="primary" @click="handleConfirmClick">
            提交
          </el-button>
        </span>
      </template>
    </h-form>
  </el-dialog>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue'
import HForm from '@/base-ui/form'
import { mainStore } from '@/store'
// import { IFormItem } from '@/base-ui/form/types'

const props = defineProps({
  dialogConfig: {
    type: Object,
    required: true
  },
  defaultInfo: {
    type: Object,
    default: () => {}
  },
  otherInfo: { type: Object, default: () => {} },
  pageName: { type: String, required: true }
})
const formData = ref<any>({})
const dialogVisible = ref(false)

watch(
  () => props.defaultInfo,
  (newValue) => {
    for (const item of props.dialogConfig?.formItem) {
      formData.value[item.field] = newValue[item.field]
    }
  }
)
const title = () => {
  if (Object.keys(props.defaultInfo).length) {
    return '编辑'
  } else {
    return '新建'
  }
}
const store = mainStore()
const content = ref<InstanceType<typeof HForm>>()

const handleConfirmClick = () => {
  content.value?.validate(async (valid: any) => {
    if (valid) {
      dialogVisible.value = false
      if (Object.keys(props.defaultInfo).length) {
        store.updatePageDataAction({
          pageName: props.pageName,
          id: props.defaultInfo.ID,
          data: { ...formData.value, ...props.otherInfo }
        })
      } else {
        store.createPageDataAction({
          pageName: props.pageName,
          data: { ...formData.value, ...props.otherInfo }
        })
      }
    }
  })
}
defineExpose({
  dialogVisible
})
</script>

<style lang="scss" scoped>
.dialog-footer {
  display: flex;
  justify-content: center;
}
</style>
