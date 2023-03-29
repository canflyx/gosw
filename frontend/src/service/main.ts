import hfRequest from './index'

// post 请求服务器端资源
export function getPageListData(url: string, queryInfo: any) {
  return hfRequest.post({
    url: url,
    data: queryInfo
  })
}
// url /pageName/id
export function deletePageDate(url: string) {
  return hfRequest.delete({
    url: url
  })
}
export function createPageDate(url: string, createInfo: any) {
  return hfRequest.post({
    url: url,
    data: createInfo
  })
}
// url /pageName/id
export function updatePageDate(url: string, updateInfo: any) {
  return hfRequest.patch({
    url: url,
    data: updateInfo
  })
}
export function scanPageData(url: string, list: any) {
  return hfRequest.post({
    url: url,
    data: list
  })
}
