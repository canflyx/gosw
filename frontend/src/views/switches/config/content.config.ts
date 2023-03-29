import { IContent } from '@/base-ui/table/types'

export const contentConfig: IContent = {
  propList: [
    { prop: 'ip', label: 'IP地址', minWidth: '55' },
    { prop: 'user', label: '用户名', minWidth: '60' },
    {
      prop: 'UpdateAt',
      label: '更新时间',
      minWidth: '70',
      slotName: 'UpdatedAt'
    },
    {
      prop: 'iscore',
      label: '类型',
      minWidth: '30',
      slotName: 'iscore'
    },
    { prop: 'status', label: '状态', minWidth: '30', slotName: 'status' },
    { label: '操作', minWidth: '80', slotName: 'handle' }
  ],
  showIndexColum: true,
  selectColum: true
}
