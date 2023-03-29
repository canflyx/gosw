export interface IContent {
  title?: string
  propList: any[]
  showIndexColum?: boolean
  selectColum?: boolean
  childrenProps?: object
  showFooter?: boolean
}
export interface IPermission {
  isCreate: boolean
  isUpdate: boolean
  isDelete: boolean
  isQuery: boolean
}
