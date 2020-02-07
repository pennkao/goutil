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
 |- extnet
        |- ip.go      //核心文件
        |- ip_test.go //测试文件
        |- README.md     //说明文档
```
## 包名
- 如果与golang包冲突的,包名加ext
- 例如:golang包time 命名exttime

## 目录分类列表

| 功能 | 包名 |  备注 |
| :--- | :--- | :--- |
| net扩展包 |[extnet](extnet/README.md) | 网络扩展包 |
| http扩展包 |[exthttp](exthttp/README.md) | 网络扩展包 |
| html扩展包 |[exthtml](exthtml/README.md) | 网络扩展包 |
| hash扩展包 |[exthash](exthash/README.md) | 网络扩展包 |
| strings扩展包 |[extstrings](extstrings/README.md) | 网络扩展包 |
| time扩展包 |[exttime](exttime/README.md) | 网络扩展包 |
| extmap操作 |[extmap](extmap/README.md) | 网络扩展包 |
| 数值转换 | [conv](conv/README.md) | 操作数字等 |
| 加密 | [encrypt](encrypt/README.md) |  获取自定义时间格式等|
| 文件操作 | [file](file/README.md) |  获取文件目录,读取,写等|
| 分页操作 | [page](page/README.md) |  用于数据分页操作等|
| slice操作 |[slice](slice/README.md) | 网络扩展包 |
| 目录操作 | [pwdtools](pwdtools/pwdtools.go) | 获取目录 |


## 贡献来源
> 

### 部分来源名单




## 欢迎加入本团队
> 


