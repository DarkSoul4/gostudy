# gostudy

## 一、环境安装

1. go下载：golang.google.cn/dl/
2. vscode安装：https://code.visualstudio.com/Download，原下载地址：	https://*`az764295.vo.msecnd.net`*/stable/054a9295330880ed74ceaedda236253b4f39a335/VSCodeUserSetup-x64-1.56.2.exe，替换为：https://***vscode.cdn.azure.cn***/stable/054a9295330880ed74ceaedda236253b4f39a335/VSCodeUserSetup- x64-1.56.2.exe
3. 新建go开发文件夹，环境变量配置：

![image-20210602164027292](C:\Users\wuj\AppData\Roaming\Typora\typora-user-images\image-20210602164027292.png)

![image-20210602164201464](C:\Users\wuj\AppData\Roaming\Typora\typora-user-images\image-20210602164201464.png)

![image-20210602164318403](C:\Users\wuj\AppData\Roaming\Typora\typora-user-images\image-20210602164318403.png)

​    4.vscode配置go环境包

​	   go env -w GOPROXY=https://goproxy.io,direct

​	   重启vscode，Ctrl+Shift+P，输入Go:Install/Update Tools，安装

## 二、helloworld

1. 在当前目录下执行go build xxx.go，生成xxx.exe
2. 当前目录下执行xxx.exe

