# 使用方法

1、 添加私有库
```shell script
go env -w GOPRIVATE=git.cifuwu.net
```

2、 安装
```shell script
go get github.com/lich-go/wand.git

# 或者可以指定`tag`对应的版本号
go get github.com/lich-go/wand.git@v1.0.0
```
3、 参考wiki

# 注意
go 包发布时版本号前需要加小写字母 v，后面版本号根据GUN规则进行规范。

由于 go 包默认会进行小版本升级，所以如果修改了API则使用大版本号进行替换。