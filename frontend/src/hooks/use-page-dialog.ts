import { ref } from 'vue'
import PageDialog from '@/components/page-dialog'

type CallbackFn = (item?: any) => void
export function usePageDialog(newCb?: CallbackFn, editCb?: CallbackFn): any {
  const pageDialogRef = ref<InstanceType<typeof PageDialog>>()
  const defaultInfo = ref({})
  const handleAddClick = () => {
    defaultInfo.value = {}
    if (pageDialogRef.value) {
      pageDialogRef.value.dialogVisible = true
    }
    // 有值就执行
    newCb && newCb()
  }
  const handleEditClick = (item: any) => {
    defaultInfo.value = { ...item }
    if (pageDialogRef.value) {
      pageDialogRef.value.dialogVisible = true
    }
    editCb && editCb(item)
  }
  return [pageDialogRef, defaultInfo, handleAddClick, handleEditClick]
}
