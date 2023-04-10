import { type IForm } from '@/base-ui/form/types'
export const searchConfig: IForm = {
  formItem: [
    {
      field: 'switch_ip',
      label: '交换机IP：',
      type: 'input',
      placeholder: '输入交换机IP'
    }
  ],
  colLayout: { span: 16 },
  itemStyle: { width: '80%' }
}
