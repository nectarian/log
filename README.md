# log

logger module based on uber zap

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

package `log` will read `logger.toml` file in current folder , create and initialize zap logger. default config will be used if `logger.toml` does not exsist.

## log.Config property

* **Level** : Main logger level, Can be one of : `debug`, `info`, `warn`, `error`, `dpanic`, `panic`, `fatal`. Case insensitive, default : `INFO`
* **Output.File** :  Whether to output to fil? default : `false`
* **Output.FilePath** : File path for output to file, default : `/tmp/$APP_NAME.log`
* **Output.MaxSize** : Log file max size, Unit : **MB**, default : `10`
* **Output.MaxBackups** : How many copies can be stored in the archive, default : `30`
* **Output.MaxAge** : How many days the archive can be kept at most, default : `7`
* **Output.Compress** : Compressed the archive or not , default : `false`
* **Output.Console** : Whether to output to the console? `Fix` function will change this value as `true` if File is `false`, default : `true`
* **Context.Format** : log format, can be one of : `json`, `console`. default : `console`
* **Context.Fileds** : additional fields, type : `map[string]string`, default : `nil`
* **Context.Caller** : whether to record the caller file name and line number, default : `true`
