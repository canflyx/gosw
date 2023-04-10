import { IContent } from '@/base-ui/table/types'

export const contentConfig: IContent = {
  propList: [
    { prop: 'switch_ip', label: 'IP地址', minWidth: '55' },
    { prop: 'log', label: '用户名', minWidth: '60', slotName: 'log' },
    {
      prop: 'UpdatedAt',
      label: '更新时间',
      minWidth: '70',
      slotName: 'UpdatedAt'
    }
  ],
  showIndexColum: true
}
