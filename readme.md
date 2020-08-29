# 对[bind.Wrap](https://github.com/lun-zhang/zlutils/tree/master/bind) 的接口解析swagger.json

先将本项目编译`go build`成二进制执行文件`zlswagger`

执行命令:
```shell script
./zlswagger \
-dir="改成你的main.go所在的目录" \
-title="示例标题" \
-desc="示例描述"
```
会在执行目录下生成一个swagger.json文件，粘贴到 https://editor.swagger.io 就能看到接口文档了