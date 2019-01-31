# Rime五笔091版
用go转换码表自定义五笔091版

## clone项目 并转换码表()
```shell
# clone
git clone git@github.com:mozyy/rime.git $USERPROFILE/AppData/Roaming/Rime

# to rime
cd $USERPROFILE/AppData/Roaming/Rime

# run convert codeTabel
go run ./
```

```shell
# build
go build  -o ./convertcodetable.exe ./
```