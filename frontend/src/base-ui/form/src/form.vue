<template>
  <div class="h-form">
    <div class="header">
      <slot name="header" />
    </div>
    <el-form
      ref="ruleFormRef"
      :label-width="labelWidth"
      :model="props.modelValue"
      label-position="right"
    >
      <el-row>
        <template v-for="item in formItem" :key="item.label">
          <el-col v-bind="colLayout">
            <div class="el-form-item">
              <el-form-item
                v-if="!item.isHidden"
                :label="item.label"
                :rules="item.rules"
                :style="itemStyle"
                :prop="item.field"
              >
                <template
                  v-if="
                    item.type === 'input' ||
                    item.type === 'password' ||
                    item.type === 'textarea'
                  "
                >
                  <el-input
                    :placeholder="item.placeholder"
                    v-bind="item.otherOptions"
                    :type="item.type"
                    :show-password="item.type === 'password'"
                    :model-value="modelValue[`${item.field}`]"
                    @update:modelValue="handleValueChange($event, item.field)"
                  />
                </template>
                <template v-else-if="item.type === 'text'">
                  <el-text type="info">{{ item.text }}</el-text>
                </template>
                <template v-else-if="item.type === 'switch'">
                  <el-switch
                    style="
                      --el-switch-on-color: #ff4949;
                      --el-switch-off-color: #13ce66;
                    "
                    :active-text="item.activeText"
                    :inactive-text="item.inactiveText"
                    :active-value="item.activeValue"
                    :inactive-value="item.inactiveValue"
                    :model-value="modelValue[`${item.field}`]"
                    @update:modelValue="handleValueChange($event, item.field)"
                  />
                </template>
                <template v-else-if="item.type === 'radio'">
                  <el-radio-group
                    :model-value="modelValue[`${item.field}`]"
                    @update:modelValue="handleValueChange($event, item.field)"
                  >
                    <el-radio-button
                      v-for="opts in item.options"
                      :key="opts.label"
                      :label="opts.label"
                      >{{ opts.value }}</el-radio-button
                    >
                  </el-radio-group>
                </template>
                <template v-else-if="item.type === 'select'">
                  <el-select
                    filterable
                    allow-create
                    default-first-option
                    :placeholder="item.placeholder"
                    :model-value="modelValue[`${item.field}`]"
                    @update:modelValue="handleValueChange($event, item.field)"
                  >
                    <el-option
                      v-for="opts in item.options"
                      :key="opts.value"
                      :value="opts.value"
                      :label="opts.label"
                      :model-value="modelValue[`${item.field}`]"
                      @update:modelValue="handleValueChange($event, item.field)"
                    />
                  </el-select>
                </template>
              </el-form-item>
            </div>
          </el-col>
        </template>
      </el-row>
      <div class="footer">
        <slot name="footer" />
      </div>
    </el-form>
  </div>
</template>

<script setup lang="ts">
import { ref, PropType } from 'vue'
import type { FormInstance } from 'element-plus'
import { IFormItem } from '../types/index'
// const props = withDefaults(
//   defineProps<{
//     formItem: IFormItem[]
//     labelWidth?: string
//     colLayout?: object
//     itemStyle?: object
//     modelValue: object
//   }>(),
//   {
//     labelWidth: '100px',
//     itemStyle: () => ({
//       padding: '10px 40px'
//     }),
//     colLayout: () => ({
//       xl: 6, // >1920px 4个
//       lg: 8,
//       md: 12,
//       sm: 24,
//       xs: 24
//     })
//   }
// )

// for (const item of props.dialogConfig?.formItem) {
//       formData.value[item.field] = newValue[item.field]
//     }
const props = defineProps({
  modelValue: {
    type: Object,
    required: true
  },
  formItem: {
    type: Array as PropType<IFormItem[]>,
    default: () => []
  },
  labelWidth: {
    type: String,
    default: 'auto'
  },
  itemStyle: {
    type: Object,
    default: () => ({ padding: '10px 40px' })
  },
  colLayout: {
    type: Object,
    default: () => ({
      xl: 6, // >1920px 4个
      lg: 8,
      md: 12,
      sm: 24,
      xs: 24
    })
  }
})
const ruleFormRef = ref<FormInstance>()
const emits = defineEmits<{
  (e: 'update:modelValue', data: any): void
}>()
const handleValueChange = (value: any, field: string) => {
  emits('update:modelValue', { ...props.modelValue, [field]: value })
}

const validate = async (callback: any) => {
  // 这个ruleFormRef是子组件内部el-form 的ref="ruleFormRef"
  ruleFormRef.value?.validate((valid) => {
    callback(valid)
  })
}
defineExpose({ validate })
</script>

<style lang="scss" scoped>
.header {
  display: flex;
  justify-content: center;
  margin-bottom: 10px;
}
.h-form {
  padding: 10px 10px;
}
.el-form-item {
  margin-bottom: 5px;
}
</style>
