# Customization
## Debug
Echo#Debug can be used to enable / disable debug mode. Debug mode sets the log level to DEBUG.
## Logging
The default format for logging is JSON, which can be changed by modifying the header.
## Log Header
Echo#Logger.SetHeader(string) can be used to set the header for the logger. Default value:
```go
{"time":"${time_rfc3339_nano}","level":"${level}","prefix":"${prefix}","file":"${short_file}","line":"${line}"}
```
