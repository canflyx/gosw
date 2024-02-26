import { IContent } from '@/base-ui/table/types'

export const contentConfig: IContent = {
  propList: [
    { prop: 'switch_ip', label: '地址', minWidth: '20' },
    { prop: 'log', label: '日志', minWidth: '100',slotName: 'log' },
    {
      prop: 'UpdatedAt',
      label: '更新时间',
      minWidth: '30',
      slotName: 'UpdatedAt'
    }
  ],
  showIndexColum: true
}
