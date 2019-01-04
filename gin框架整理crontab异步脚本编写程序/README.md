
# crontabx下面加入main执行二进制文件，进程管理工具用python supervisor
    go build main2.go

# supervisor 配置
    [program:crontab]
    command=go run /data/gopath/src/hytx_sync/crontab/crontab.go
    autostart=true
    autorestart=true
    startsecs=10

