# location

location 通过请求路径来分配流量

## 问题记录

同时存在两个 location

- `location /fickit {...}`
- `location ~* \.(png|jpg|jpeg|gif|ico)$`

但 `/fickit/abc.png` 却去了第二个，这是为什么？

## 匹配

1. 前缀匹配（普通 location）
    先进行前缀字符串匹配，找到**最长**匹配的那个，并暂时记住它。
2. 正则匹配（~ 或 ~*）
    如果存在正则 location（~ 区分大小写，~* 不区分），按配置文件中的书写顺序依次测试。一旦某个正则匹配成功，就立即使用它，并**放弃**之前记住的前缀匹配结果。
3. 精确匹配（=）和优先前缀（^~）
    - =：精确匹配，优先级最高，找到立即返回
    - ^~：如果前缀匹配是最长的，且前面加了 ^~，则跳过正则测试，直接使用这个前缀匹配

## 例子

```nginx
server {
    listen 80;
    server_name example.com;

    # 1. 精确匹配（=）优先级最高
    location = / {
        root /var/www/exact;
        index index.html;
    }

    # 2. 优先前缀匹配（^~），匹配成功则跳过正则
    location ^~ /static/ {
        root /var/www/static;
        expires 30d;
    }

    # 3. 普通前缀匹配
    location /api {
        proxy_pass http://backend;
    }

    # 4. 区分大小写的正则匹配（~）
    location ~ \.php$ {
        fastcgi_pass 127.0.0.1:9000;
        fastcgi_index index.php;
    }

    # 5. 不区分大小写的正则匹配（~*）
    location ~* \.(png|jpg|jpeg|gif|ico)$ {
        root /var/www/images;
        expires 30d;
    }

    # 6. 通用前缀匹配（兜底）
    location / {
        root /var/www/html;
        try_files $uri $uri/ =404;
    }
}
```

### 各 location 说明

| 类型 | 语法 | 作用 |
|------|------|------|
| 精确匹配 | `= /` | 仅匹配根路径 `/`，优先级最高 |
| 优先前缀 | `^~ /static/` | 前缀匹配成功后**跳过正则**，直接生效 |
| 普通前缀 | `/api` | 前缀匹配，最终可能被正则覆盖 |
| 正则（区分大小写） | `~ \.php$` | 按书写顺序匹配，成功即停止 |
| 正则（不区分大小写） | `~* \\.(png|jpg|...)$` | 同上，但不区分大小写 |
| 通用前缀 | `/` | 兜底，当其他均不匹配时生效 |

### 匹配流程示例

假设请求路径为 `/static/logo.PNG`：

1. **精确匹配**：检查 `= /`，不匹配
2. **前缀匹配**：`^~ /static/` 和 `/` 都匹配，记住最长的 `^~ /static/`
3. **检查是否有 `^~`**：有，因此**跳过所有正则**，直接使用 `^~ /static/`
4. 最终由 `^~ /static/` 处理，返回 `/var/www/static/logo.PNG`

> 注意：如果 `^~` 换成普通前缀 `/static/`，则会继续测试正则。`~* \.(png|jpg|...)$` 不区分大小写可以匹配 `.PNG`，且按配置顺序在 `~ \.php$` 之后，因此会命中该正则 location，去 `/var/www/images` 目录查找。

> 点：如果正则中含有`.`，前面必须添加反斜杠（否则会被当做正则的“任意字符”符号）