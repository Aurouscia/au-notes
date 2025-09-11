# 入口类
java使用`@`符号作为annotation  
`@SpringBootApplication`用来标注一个类是入口类，其`public static void main(String[] args)`方法会被用来启动程序  
它是个组合注解（composite annotation），完整形式其实有三个：
- `@Configuration`: 标注该class为应用上下文（Application context）的bean定义源
- `@EnableAutoConfiguration`: 根据classpath以及其他设置，classpath会作为开关的效果，自动注册一些bean
- `@ComponentScan`: 扫描项目，寻找@Component（包括@Service、@Controller、@Repository等）让应用之后能找到

# pom.xml
maven的“项目文件”，全称project object model  
pom.xml 是 Maven 项目的“简历 + 说明书 + 配线表”  
它用 XML 格式声明项目坐标、依赖、插件、构建规则等  
Maven 读它来完成编译、打包、测试、发布等所有生命周期任务。

# REST控制器类
- `@RestController` 标注一个类是rest控制器
- `@RequestMapping(method=GET)` 标注一个方法可以处理http请求，以下是简写形式
    - `@GetMapping("/greet")` 
    - `@PostMapping("/greet")`
- `@RequestParam` 为一个方法参数添加额外信息（请求中的名称、默认值）
```java
@RestController
public class GreetingController {
	private static final String template = "Hello, %s!";
	private final AtomicLong counter = new AtomicLong();
	@GetMapping("/greeting")
	public Greeting greeting(@RequestParam(defaultValue = "World") String name) {
		return new Greeting(counter.incrementAndGet(), String.format(template, name));
	}
}
```

# bean
bean就是服务，会被Spring的Ioc容器按需构造
## 定义bean
- `@Bean`必须加在**方法**上，该方法返回一个bean对象
- 上述方法必须在某种类中，该类有`@Configuration`标注（或其元注解`@SpringBootConfiguration`）
    或：有`@Component`（`@Service`、`@Controller`、`@Repository`）
```java
@Configuration
public class SomeClass {
    @Bean
    @Scope("prototype")   //设置其生命周期
    public Foo prototypeFoo() {
        return new Foo();
    }
}
```
## scope的值
- singleton	    单例（默认值）
- prototype 	每次获取都创建新实例
- request	    每个 HTTP 请求一个实例（Web 环境）
- session   	每个 HTTP Session 一个实例（Web 环境）
- application	整个 ServletContext 一个实例（Web 环境）
## bean的生命周期
这个得单开一个文件  
bean的生命周期有一大堆可以配置和钩住的东西