# zod
又能“名义上”类型正确，又能“实际上”检查对象

```js
import { z } from 'zod';

// 1. 一次性声明 schema + TypeScript 类型
const UserSchema = z.object({
  id:    z.number(),
  name:  z.string(),
  email: z.string().email()
});
type User = z.infer<typeof UserSchema>;   // 编译期类型

// 2. 运行时校验任意数据
const raw: any = await fetch('/api/user').then(r => r.json());
const user = UserSchema.parse(raw);       // 不合格直接抛异常
// user 现在就是 User 类型，安全使用
```