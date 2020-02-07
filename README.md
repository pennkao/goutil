# goutil
> go常用工具包

# 安装
```
go get -u github.com/pennkao/goutil
```
//todo ip

## 目录说明
- 一个类别一个文件夹,文件夹里必含一个测试文件
- 可以使用go test测试
- 如array

```
 |- array
        |- array.go      //核心文件
        |- array_test.go //测试文件
        |- readme.md     //说明文档
```
## 包名
- 如果与golang包冲突的,包名加ext
- 例如:golang包time 命名exttime

## 目录分类列表

| 功能 | 包名 |  备注 |
| :--- | :--- | :--- |
| 数值转换 | [conv](conv/readme.md) | 操作数字等 |
| 加密 | [encrypt](encrypt/readme.md) |  获取自定义时间格式等|
| 文件操作 | [file](file) |  获取文件目录,读取,写等|
| 分页操作 | [page](page) |  用于数据分页操作等|
| 字符串扩展 | [extstrings](extstrings) | 字符串操作|
| 切片扩展 | [slice](slice) |  linux相关等|
| 常用hash函数 | [hash](hash/README.md) |  string、byte、file 的hash值 包括md5 sha1 sha256 sha512 |
| http包扩展 | [curl](curl/curl.go) |  curl get ,post 请求 |
| 目录操作 | [pwdtools](pwdtools/pwdtools.go) | 获取目录 |


## 贡献来源
> 

### 部分来源名单




## 欢迎加入本团队
> 


