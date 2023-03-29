import { IForm } from '@/base-ui/form/types'
export const dialogConfig: IForm = {
  formItem: [
    {
      field: 'ip',
      label: '交换机地址:',
      type: 'textarea',
      placeholder: '多个IP用";"分隔,连续IP最后一位用"-"',
      rules: [{ required: true, message: 'IP地址不能为空', trigger: 'blur' }]
    },

    {
      field: 'user',
      label: '用户名:',
      type: 'input',
      placeholder: '管理用户名',
      rules: [
        {
          required: true,
          message: '用户名不能为空',
          trigger: 'blur'
        }
      ]
    },
    {
      field: 'password',
      label: '密码:',
      type: 'password',
      placeholder: '管理密码',
      isHidden: false,
      rules: [{ required: true, message: '管理密码不能为空', trigger: 'blur' }]
    },
    {
      field: 'brand',
      label: '品牌:',
      type: 'select',
      placeholder: '自定义直接输入',
      options: [
        { label: '华为', value: 'huawei' },
        { label: '华三', value: 'h3c' },
        { label: '锐捷', value: 'ruijie' },
        { label: '思科', value: 'cisco' },
        { label: '其它', value: 'default' }
      ]
    },
    {
      field: 'iscore',
      label: '类型:',
      type: 'radio',
      placeholder: '交换机状态',
      options: [
        { label: 0, value: '接入' },
        { label: 1, value: '核心' }
      ]
      // options: [
      //   { label: '接入', value: 1 },
      //   { label: '核心', value: 0 }
      // ],
      // rules: [
      //   { required: true, message: '交换机类型必须选择', trigger: 'blur' }
      // ]
    },
    {
      field: 'status',
      label: '状态:',
      type: 'switch',
      activeValue: 1,
      inactiveValue: 0,
      activeText: '禁用',
      inactiveText: '正常'
    },
    {
      field: 'note',
      label: '备注:',
      type: 'input',
      placeholder: '输入备注'
    }
  ],
  labelWidth: '100px',
  colLayout: {
    span: 24
  },
  itemStyle: { padding: '10px 40px', width: '80%' }
}
