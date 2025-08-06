# 配置远程服务器

1. 确保远程服务器上有IIS8.0+
2. 服务器管理器-添加角色和功能-服务器角色-iis-管理工具-管理服务，安装
3. cmd运行：
    ```
    net start msdepsvc
    net start wmsvc
    ```
    启用两个服务
4. 在 https://www.microsoft.com/zh-cn/download/confirmation.aspx?id=43717 下载webDeploy_amd64_zhCN.msi，然后在服务器上双击安装，**安装所有可选的功能**
    - 注意要先添加第二步的角色，webdeploy才有完整的安装选项，不然不全
    - 注意要先配置好iis，再安装整个webdeploy（否则先卸载再安装），不然iis配置文件不齐全
5. 创建用户组MSDepSvcUsers，创建一个新用户放到这个组里
    - 或者使用iis自己的账户：[https://www.cnblogs.com/longbky/p/11884449.html]
6. 给C:\Windows\System32\inetsrv\config添加EveryOne的读取权限
7. 确保8172端口在windows防火墙和服务商防火墙都开着

# 填写vs的发布配置文件
- 服务器填ip（有时候需要http://，有时候不要，玄学）
- 网站名填IIS内的网站名
- 用户名和密码填刚才注册的
- 目标url填网站的网址

# 用了一段时间后出现问题
无法执行此操作。请与服务器管理员联系，检查授权和委派设置。  
打开计算机管理-用户和组-找到用户WDeployAdmin和WDeployConfigWriter-属性-密码永不过期
[https://xoyozo.net/Blog/Details/WDeployAdmin-WDeployConfigWriter]