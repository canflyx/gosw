// 统一出口
import HFRequest from './request/'
import { BASE_URL, TIME_OUT } from './request/config'

const hfRequest = new HFRequest({
  baseURL: BASE_URL,
  timeout: TIME_OUT
})
export default hfRequest
