# if + -e 示例

以下代码展示了“如果请求文件不存在则走代理”的设置，`-e`表示存在

```conf
location /
{
    if (!-e $request_filename) {
        proxy_pass http://$swoole_backend;
    }
}
```