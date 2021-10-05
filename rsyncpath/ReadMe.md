## 概要
同一应用多主机部署很常见，但是次数会出现用户上传文件无法同步的问题。为此引入网络存储文件系统又会出现其他一系列的问题。如果可以选择minio，当然此处利用go实现。实现多主机对统一目录自动同步。


## 使用方法
* 编译
	通过buildXXX方式即可编译相应平台的可执行文件
* 配置
	配置使用json方式，简单明了
	```json
	{
	    "addr":"0.0.0.0:9090",			//配置监地址和端口
	    "path":"c:/var/upload/wwww/",	//文件存储路径
	    "rysncAddr":[					//同步地址
	        "http://localhost:9091/rsync",
          "http://localhost:9092/rsync",
	    ]
	}
	```
* 启动
```base
UploadRysnc -conf conf/server1.conf > run.log

```

## 运行
下面是运行部分运行日志
```log
2019/03/17 10:40:00 Server is starting:0.0.0.0:9090
2019/03/17 10:40:00 Server UploadPath:c:/var/upload/wwww/
2019/03/17 10:40:00 Server Rysnc Addr:http://localhost:9091/rsync
2019/03/17 10:40:10 [::1]:49743 uploadfile [server1.conf][server1.conf] > c:/var/upload/wwww/banner/LwhSfU1nh6w.conf
2019/03/17 10:40:31 Clientrsyncfile Error http://localhost:9091/rsync c:/var/upload/wwww/banner/LwhSfU1nh6w.conf 
```