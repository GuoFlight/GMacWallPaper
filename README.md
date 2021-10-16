# 作者信息

京城郭少

# 关于项目

> 适用于Mac，在连接不同显示器时，会切换不同的壁纸。

# 配置文件

* ```config.toml```是配置文件。
* 默认会显示```[default]```指定的壁纸。
* 连接了```[special]```指定的显示器后，会显示```[special]```指定的壁纸。
* ```path```可以指定文件；也可以指定目录，这时候会1小时设置1次此目录的图片。


# Example

运行一次程序：

```shell
go run mian.go
#或
go run main.go -c ./config.toml
```

添加计划任务：

```shell
#需要注意的是，在mac上使用计划任务，需要给cron命令添加"全磁盘访问权限"。
crontab -e    #添加计划任务
```