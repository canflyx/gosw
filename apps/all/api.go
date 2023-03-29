package all

import (
	// 注册所有HTTP服务模块, 暴露给框架HTTP服务器加载
	_ "github.com/canflyx/gosw/apps/maclist/api"
	_ "github.com/canflyx/gosw/apps/switches/api"
)
