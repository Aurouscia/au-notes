# 网络
一些网连不了sqlserver（仅保留了常用端口，其他的全部禁了），要用手机热点

# 端口号
- sqlserver的默认端口号是1433
- 确保服务器运营商设置处可以访问（确保安全组设置了端口号开放，确保安全组里有这个实例）
- 确保系统防火墙能允许端口通过

## 设置端口号
- 用`sqlServer配置管理器`可设置服务监听的端口号
    - 找到`网络配置-<实例名>的协议`
    - 启用`TCP/IP`，打开其属性
    - 找到`IPALL`的`TCP端口`即可调整
- 如果不是1433，那么SMSS登录或者连接字符串的主机名就要填`<ip>,<端口号>`格式，中间逗号隔开
- 更换端口号有助于减少攻击
- 下面的方法可以检查服务实际使用的端口

```
In most cases, you connect to the Database Engine on another computer by using the TCP protocol. To get the TCP port of the instance, follow these steps:
Use SQL Server Management Studio on the computer running SQL Server and connect to the instance of SQL Server. In Object Explorer, expand Management, expand SQL Server Logs, and then double-click the current log.
In the Log File Viewer, select Filter on the toolbar. In the Message contains text box, type server is listening on, select Apply filter, and then select OK.
A message like Server is listening on [ 'any' ipv4 1433] should be listed.
This message indicates that the instance of SQL Server is listening on all IP addresses on this computer (for IP version 4) and TCP port 1433. (TCP port 1433 is usually the port that's used by the Database Engine or the default instance of SQL Server. Only one instance of SQL Server can use this port. If more than one instance of SQL Server is installed, some instances must use other port numbers.) Note down the port number used by the SQL Server instance that you're trying to connect to.
```


# NamedPipes
在新版sqlServer上，efcore默认使用的NamedPipe连接与sqlserver默认的名称对不上，这种时候应该手动指定tcp连接
- 在连接字符串的Server前加“tcp:”前缀即可，例如“tcp:localhost,1433”