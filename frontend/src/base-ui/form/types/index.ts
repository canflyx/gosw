type IFormType =
  | 'input'
  | 'password'
  | 'select'
  | 'datepicker'
  | 'textarea'
  | 'text'
  | 'switch'
  | 'radio'

export interface IFormItem {
  field: string
  type: IFormType
  label?: string
  rules?: any[]
  placeholder?: string
  options?: any[]
  otherOptions?: any
  switch?: any[]
  isHidden?: boolean
  activeValue?: number
  inactiveValue?: number
  activeText?: string
  inactiveText?: string
  text?: string
}
export interface IForm {
  formItem: IFormItem[]
  labelWidth?: string
  colLayout?: any
  itemStyle?: any
}
