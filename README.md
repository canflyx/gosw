# 最新

所有版本升级至 2024.6 最新版本，go ver 1.22.2

- windows wails 编译
  wails build -webview2 embed

# 交换机批量运行工具

gui 采用 wails 工具，如果自行编译先安装好 wails 最新版本

https://wails.io/zh-Hans/docs/gettingstarted/installation/

## 完成功能

![界面](main.png)

自启动会生成 sqlite db 文件和 yaml 配置文件，默认为 ">",可以不写

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
    pre_cmd: # 预处理命令,为 [cmd,cmd_flag] 的数组,必须命令和标识一一对应
      - cmd: sys
        cmd_flag: "]"
      - cmd: user-interface vty 0 4
        cmd_flag: "]"
      - cmd: screen-length 0
        cmd_flag: "]"
      - cmd: quit
        cmd_flag: ""
      - cmd: quit
        cmd_flag: ">"
    user_cmd: [] #用于自定义批处理命令。
     access_cmd: dis mac-address #读取 mac-address命令
    core_cmd: dis arp    #读取arp与ip的对应表
    read_flag: ">"   #针对前两项的读取标志
    exit_cmd: #退出命令,一般与 PreCmd同时出现，运行数据不会被读取。
      - cmd: sys
        cmd_flag: "]"
      - cmd: user-interface vty 0 4
        cmd_flag: "]"
      - cmd: screen-length 50
        cmd_flag: "]"
      - cmd: quit
        cmd_flag: ""
      - cmd: quit
        cmd_flag: ""
      - cmd: quit
        cmd_flag: ""
        read_cmd: dis mac-add #需要读取终端执行的命令
        core_cmd: dis arp # 与 read_cmd 二选一，现阶段此为核心交换机需要执行的
        read_flag: "]"
        exit_cmds: #退出之前执行的命令
          - cmd: screen-length 0
            cmd_flag: "]"
```

## linux 编译环境

需要安装 vue 及 golang 的环境，需要安装 libwebkit2gtk-4.0-dev 包，不需要桌面编译。

- 安装 node，并下载依赖包
  http://nodejs.cn/

```shell
  wget https://nodejs.org/dist/v10.15.3/node-v10.15.3-linux-x64.tar.xz
  tar -xvf node-v10.15.3-linux-x64.tar.xz
  mv node... /usr/local/node
  vi /etc/profile
  export NODE_HOME=/usr/local/node
  export PATH=$NODE_HOME/BIN:$PATH

  cd frontend/
  npm install
```

- 安装 golang

```shell
 tar -C /usr/local -zxvf go1.20.3.linux-amd64.tar.gz
 vim /etc/profile

 export GOROOT=/usr/local/go
 export PATH=$PATH:$GOROOT/bin

 source /etc/profile
 go mod tidy

 wails build -upx
```
