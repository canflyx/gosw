import { IContent } from '@/base-ui/table/types'

export const contentConfig: IContent = {
  propList: [
    { prop: 'arp_ip', label: 'IP地址', minWidth: '40' },
    { prop: 'mac_address', label: 'MAC地址', minWidth: '50' },
    { prop: 'switch_ip', label: '交换机', minWidth: '40' },
    {
      prop: 'port',
      label: '交换机端口',
      minWidth: '40'
    },
    {
      prop: 'UpdatedAt',
      label: '更新时间',
      minWidth: '60',
      slotName: 'UpdatedAt'
    }
  ],
  showIndexColum: true
}
