# ExecuteDelete报错

## 报错：
ExecuteDelete如果报错：`There is no method 'ExecuteDelete' on type 'Microsoft.EntityFrameworkCore.RelationalQueryableExtensions' that matches the specified arguments`

## 原因：
很有可能是因为其前面的IQueryable语句转换失败了

## 这样写会失败：
public string GetKey(){return "topbar";}
where(x=>x.key==GetKey())
因为efcore没法把GetKey函数转换为sql语句

## 这样写会成功：
var str = "topbar"
where(x=>x.key==str)
efcore可以把字符串变量转换为sql语句的一部分