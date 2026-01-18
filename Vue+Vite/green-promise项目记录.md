1. 周期性自动操作（从服务器同步）和触发操作（点击修改）的互斥锁，不能让它们同时被处理

2. 花里胡哨的transition要考虑软键盘弹出/缩回和目标浏览器（微信内置）是否支持
    - 微信浏览器中背景图片transition第一次需要“热身”运动，1ms切过去切回来

3. 后端：Deleted标识忘记判断（Existing()）

4. 后端：杜绝ExecuteUpdate与SaveChanges混用