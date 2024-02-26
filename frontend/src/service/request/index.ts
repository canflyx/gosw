import axios from 'axios'
import type { AxiosInstance } from 'axios'
import { HFRequestInterceptors, HFRequestConfig } from './types'
import { type LoadingInstance } from 'element-plus/lib/components/loading/src/loading.js'
import { ElLoading } from 'element-plus'

const DEFAULT_LOADING = false
class HFRequest {
  // 使用构造函数，让每个对象创建的时候都保存在 instance 中
  instance: AxiosInstance
  interceptors?: HFRequestInterceptors
  loading?: LoadingInstance
  showLoading?: boolean

  constructor(config: HFRequestConfig) {
    // 创建axios实例
    this.instance = axios.create(config)
    // 保存基本信息
    this.showLoading = config.showLoading ?? DEFAULT_LOADING
    this.interceptors = config.interceptors
    // 从 config 中取出的拦截器是对应实例的拦截器
    // this.instance.interceptors.request.use(
    //   this.interceptors?.requestInterceptor,
    //   this.interceptors?.requestInterceptorCatch
    // )
    // this.instance.interceptors.response.use(
    //   this.interceptors?.responseInterceptor,
    //   this.interceptors?.responseInterceptorCatch
    // )
    // 添加公共拦截器
    this.instance.interceptors.request.use(
      (config) => {
        // console.log('再次拦截请求')
        if (this.showLoading) {
          this.loading = ElLoading.service({
            lock: true,
            text: '加载中。。。。',
            background: 'rgba(0, 0, 0, 0.5)',
            fullscreen: true
          })
        }
        return config
      },
      (err) => {
        return err
      }
    )
    this.instance.interceptors.response.use(
      (res) => {
        // 对数据进行预验证
        // console.log(res)
        this.loading?.close()
        // const data = res.data
        return res.data
        // if (data.returnCode !== 'SUCCESS') {
        //   console.log('请求失败')
        // } else {
        //   return res.data
        // }
      },
      (err) => {
        this.loading?.close()
        return err
      }
    )
  }
  // 使用已整合拦截器接口的接口，就可以接收拦截器参数了，返回是 Promise
  // 返回的类型应该是请求所决定的。如下，返回的数据是 datatype
  // interface DataType{
  //   data: any,
  //   returnCode: string,
  //   success: boolean
  // }
  // hyRequest.request<DataType>({
  //   url:'/home/multidata',
  //   method:'GET'
  // })
  // 此处 then 返回的 res 就是 datatype类型，那里面就会有 data,returnCode ...
  // .then((res) => { console.log(res.data);console.log(res.success)
  // })
  // T=any 为默认值
  request<T = any>(config: HFRequestConfig<T>): Promise<T> {
    // 外面调用需要返回数据
    return new Promise((resolve, reject) => {
      // 1.单个请求对请求config的处理
      console.log("new Promise:",config.method,config.url)
      if (config.interceptors?.requestInterceptor) {
        config = config.interceptors.requestInterceptor(config)
      }

      // 2.判断是否需要显示loading
      if (config.showLoading === false) {
        this.showLoading = config.showLoading
      }

      this.instance
        .request<any, T>(config)
        .then((res) => {
          // 1.单个请求对数据的处理
          if (config.interceptors?.responseInterceptor) {
            res = config.interceptors.responseInterceptor(res)
          }
          // 2.将showLoading设置true, 这样不会影响下一个请求
          this.showLoading = DEFAULT_LOADING

          // 3.将结果resolve返回出去
          resolve(res)
        })
        .catch((err) => {
          // 将showLoading设置true, 这样不会影响下一个请求
          this.showLoading = DEFAULT_LOADING
          reject(err)
          return err
        })
    })
  }
  // 封装，get 直接调用上面的 request
  get<T = any>(config: HFRequestConfig<T>): Promise<T> {
    return this.request<T>({ ...config, method: 'GET' })
  }
  post<T = any>(config: HFRequestConfig<T>): Promise<T> {
    return this.request<T>({ ...config, method: 'POST' })
  }
  delete<T = any>(config: HFRequestConfig<T>): Promise<T> {
    return this.request<T>({ ...config, method: 'DELETE' })
  }
  patch<T = any>(config: HFRequestConfig<T>): Promise<T> {
    return this.request<T>({ ...config, method: 'PATCH' })
  }
}

export default HFRequest
