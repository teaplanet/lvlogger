`lvlogger` (Level Logger) is zap wrapper for like a `log` .

## How to use
```go
import "github.com/teaplanet/lvlogger/log"
```

The `lvlogger` has the same methods as `zap` . 

```go
log.Debug(args ...interface{})
log.Debugf(template string, args ...interface{})
log.Debugw(msg string, keysAndValues ...interface{})
...
log.Warn(...)
log.Info(...)
log.Error(...)
log.Fatal(...)
```

## Environment
Switch between the `zap.NewDevelopment` and `zap.NewProduction` methods by setting the environment variable `LVLOGGER_ENV` .

If `LVLOGGER_ENV` starts with `P` or `p`, you are in Production mode. 
Other than that, it is Development mode.