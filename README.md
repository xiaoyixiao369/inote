# iNote是一款开源，免费，简洁的单页博客
> iNote是在[beego(golang语言)](http://beego.me/ "beego(golang语言)")&[bootstrap](http://getbootstrap.com/ "bootstrap")等开源项目基础之上开发的

## 功能简介
- 前后端完全分离
- 响应式布局
- 内嵌[markdown](https://pandao.github.io/editor.md/ "markdown")编辑器
- URL支持hash(#)+id 文章导航
- 支持更换首页banner大图背景
- 文章功能
- 文章内容照片墙预览
- 文章标签功能
- 文章留言功能
- Web后台管理


## Linux环境编译安装（OSX,Win环境类似）
1. 安装GO      
   参考[install golang](http://golang.org/doc/install#tarball "install golang")
2. 安装mysql      
   参考[install mysql](http://dev.mysql.com/doc/refman/5.6/en/installing.html "install mysql")
3. 安装beego, bee工具(可选), mysql驱动        
               go get github.com/astaxie/beego    
               go get github.com/beego/bee      
               go get github.com/go-sql-driver/mysql        
	      
4. 安装iNote           
               go get github.com/igordonshaw/inote     
	   
5. 新建数据库inode并导入初始化脚本($GOPATH/src/github/igordonshaw/inote/dbinit/inote.sql)      
6. 按照实际情况修改iNote配置文件中的程序运行模式、监听端口及数据库参数        
                ###################### 程序基本配置 ############################
    
		# 程序运行实例名称
		appname = inote

		# 程序运行模式  dev:开发模式  prod:产品模式
		runmode = dev

		# 程序运行监听端口
		httpport = 8080
		
		# MYSQL地址
		dbhost = localhost

		# MYSQL端口
		dbport = 3306

		# MYSQL用户名
		dbuser = root

		# MYSQL密码
		dbpassword = root

		# MYSQL数据库名称
		dbname = inote

7. 编译iNote       
		cd $GOPATH/src/github/igordonshaw/inote     
		go build
		
8. 运行iNote(nohup模式)          
    		nohup ./inote &
		
9. 访问iNote      
		首页：localhost:8080  
		后台登录：localhost:8080/login      
                默认密码：admin

	<img src="https://raw.githubusercontent.com/igordonshaw/inote/master/screenshot/21A9C0EB-30AB-4512-96C3-4FCC754F9E80.png" width="200" height="350"/>

	[My iNote demo](http://120.55.100.241 "DEMO")
    
