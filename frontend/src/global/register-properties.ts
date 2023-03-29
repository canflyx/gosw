import { App } from 'vue'
import { formatUtcString } from '@/hooks/date-format'
export default function registerProperties(app: App) {
  // 注册全局属性，可以直接调用 $filters.formatTime(...)
  app.config.globalProperties.$filters = {
    formatTime(value: string) {
      return formatUtcString(value)
    }
  }
}
