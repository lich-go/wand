# 使用方法

1、 添加私有库
```shell
# 私有库将不通过代理访问
# go env -w GOPRIVATE=github.com
```

2、 安装
```shell
go get github.com/lich-go/wand

# 或者可以指定`tag`对应的版本号
go get github.com/lich-go/wand@v1.0.0
```
3、 参考wiki

# 注意
go 包发布时版本号前需要加小写字母 v，后面版本号根据GUN规则进行规范。

由于 go 包默认会进行小版本升级，所以如果修改了API则使用大版本号进行替换。
