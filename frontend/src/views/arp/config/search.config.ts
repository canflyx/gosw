import { type IForm } from '@/base-ui/form/types'
export const searchConfig: IForm = {
  formItem: [
    {
      field: 'arp_ip',
      label: 'IP地址：',
      type: 'input',
      placeholder: '搜索接入IP'
    },
    {
      field: 'mac_address',
      label: 'MAC地址：',
      type: 'input',
      placeholder: '搜索MAC地址'
    },
    {
      field: 'switch_ip',
      label: '交换机IP：',
      type: 'input',
      placeholder: '搜索交换机IP'
    }
  ],
  colLayout: { span: 8 },
  itemStyle: { width: '80%' }
}
