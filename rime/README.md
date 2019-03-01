# Rime五笔091版
用go转换码表自定义五笔091版

## clone项目 并转换码表()
```shell
# clone project
go get git@github.com:mozyy/tools.git

# to dir
cd $GOPATH/src/github.com/mozyy/tools/rime

# get help
go run ./ -h
# start run convert codeTabel
go run ./
```
## 然后重新部署小狼亳

## build
```shell
# build
go build  -o ./convertcodetable.exe ./
```

|参数  |说明           |
|---  |---            |
|-h   |帮助           |
|-np  |不进行码表处理  |
|-nc  |不进行复制操作  |