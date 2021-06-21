# log

基于 uber zap 的日志模块

## USAGE

```go
package main

import (
	"github.com/nectarian/log"
	"go.uber.org/zap"
)

func main() {
	log.Debug("Go", zap.Int("id", 1))
	log.Info("Go", zap.Int("id", 2))
	log.Warn("Go", zap.Int("id", 3))
	log.Error("Go", zap.Int("id", 4))
	log.Fatal("Go", zap.Int("id", 5))
	// log.DPanic("Go")
	// log.Panic("Go")
}

```

package `log` 会读取当前目录的 `logger.toml` 文件创建并初始化 zap 日志记录器，如果 `logger.toml` 文件不存在，则会使用默认配置。

## log.Config 属性说明

* **Level** : 日志级别,可以是: `debug`, `info`, `warn`, `error`, `dpanic`, `panic`, `fatal`。大小写不敏感,默认值为:`INFO`
* **Output.File** : 是否输出到文件？ 默认值为: `false`
* **Output.FilePath** : 输出文件路径，默认值为: `/tmp/$APP_NAME.log`
* **Output.MaxSize** : 日志文件最大尺寸，单位: **MB**, 默认值为: `10`
* **Output.MaxBackups** : 保存多少个归档，默认值为: `30`
* **Output.MaxAge** : 归档保存多少天，默认值为: `7`
* **Output.Compress** : 是否压缩归档，默认值为: `false`
* **Output.Console** : 是否输出到控制台？ 如果 `Output.File` 属性设置为 `false` ，调用 `Fix` 函数会强制将此值设置为`true`，默认值为: `true` 
* **Context.Format** : 日志格式，可以是: `json`, `console`，默认值为: `console`
* **Context.Fileds** : 附加字段，类型为： `map[string]string`，默认值为: `nil`
* **Context.Caller** : 是否输出调用者，默认值为: `true`
