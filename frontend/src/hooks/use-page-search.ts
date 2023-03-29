import { ref } from 'vue'
import PageContent from '@/components/page-content'
export function usePageSearch(): any {
  const pageContentRef = ref<InstanceType<typeof PageContent>>()
  const handleRestClick = () => {
    pageContentRef.value?.getPageData()
  }
  const handleQueryClick = (queryInfo: any) => {
    pageContentRef.value?.getPageData(queryInfo)
  }
  return [pageContentRef, handleRestClick, handleQueryClick]
}
