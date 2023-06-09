import type { AxiosRequestConfig, AxiosResponse } from 'axios'

export interface HFRequestInterceptors<T = AxiosResponse> {
  requestInterceptor?: (config: AxiosRequestConfig) => AxiosRequestConfig
  requestInterceptorCatch?: (error: any) => any
  responseInterceptor?: (res: T) => T
  responseInterceptorCatch?: (error: any) => any
}
export interface HFRequestConfig<T = AxiosResponse> extends AxiosRequestConfig {
  interceptors?: HFRequestInterceptors<T>
  showLoading?: boolean
}
