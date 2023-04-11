# 交换机批量运行工具

gui 采用 wails 工具，如果自行编译先安装好 wails 最新版本

https://wails.io/zh-Hans/docs/gettingstarted/installation/

## 完成功能

![界面](main.png)

自启动会生成 sqlite db 文件和 yaml 配置文件

- 交换机扫描功能
  测试 huawei/h3c 无问题。其它品牌可以通过 配置文件中一一对应。
- 配置文件

```yaml
app:
  name: go-switch
  encrypt: ""
  http:
    host: 127.0.0.1
    port: "8055" #前端指定端口为 8055,改了必须修改前端代码
    enable_ssl: false
    cert_file: ""
    key_file: ""
log:
  level: debug
  path_dir: logs #log 存放目录，默认 ./logs
  format: json # 日志格式，可选 json text
  to: file # 日志方式，可选 file,stdout
sql_type:
  path: ./app.db #sqlite 保存文件，默认./app.db
telnet_cmds: #telnet 配置文件
  - brand: default # 默认default,与交换机录入品牌对应,同一品牌不同命令可以使用不同的 brand 值
    user_flag: "name:" #登陆telnet时要求输入用户名的显示标识
    password_flag: "ssword:" #登陆telnet时要求输入密码的显示标识
    login_flag: ">" #登陆成功后显示标识
    enable_cmd: sys #进入配置模式
    enable_flag: "]" #进入配置模式后的显示标识
    cmds: # 其它命令,为 [cmd,cmd_flag] 的数组,必须命令和标识一一对应
      - cmd: user-interface vty 0 4
        cmd_flag: "]"
      - cmd: screen-length 0
        cmd_flag: "]"
    read_cmd: dis mac-add #需要读取终端执行的命令
    core_cmd: dis arp # 与 read_cmd 二选一，现阶段此为核心交换机需要执行的
    read_flag: "]"
    exit_cmds: #退出之前执行的命令
      - cmd: screen-length 0
        cmd_flag: "]"
```
