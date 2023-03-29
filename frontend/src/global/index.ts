import { App } from 'vue'
import registerProperties from './register-properties'

// 调用注册全局函数
export function globalRegister(app: App): void {
  app.use(registerProperties)
}
