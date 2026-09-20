# Stream

在 SC 中，stream 是一系列的值，任意满足`next`和`reset`的对象都算 stream（约定）

next 可以读取下一个值，而 reset 可以把读取头挪到初始位置

如果读完了，会返回 nil

示例：无限的随机序列（FuncStream）

```scd
a = FuncStream({ [1, 2, 3, 4].choose });
a.next; // 能无限返回 1-4 的随机数
```

## Routine

带 yield 关键字的即为 Routine（可记录自身状态，一种便于下次 next 继续的 Stream）

例子：

```scd
a = FuncStream({1});
a.next; // 1
a.next; // 1

(
x = Routine({ 
	1.yield;
	2.yield;
	3.yield; 
});
)

x.next; // 1
x.next; // 2
x.next; // 3
x.next; // nil（没有了）
```

# Pattern

Pattern 是 Stream 的模板（specification）

所有能接受`asStream`生成 Stream 的对象都是 Pattern

以下介绍一些常用内置 Pattern 构造函数

## Pfunc

（Pattern-Function）

根据函数创建 Pattern

```scd
a = Pfunc({ [1, 2, 3, 4].choose });
b = a.asStream; // make a stream from the pattern
5.do({ b.next.postln; }); // print 5 values from the stream
```

## Prout

（Pattern-Routine）

根据函数创建 Pattern，但这个 Pattern 产生的 Stream 是 Routine

```scd
/*
Prout is similar but returns a Routine
*/
(
a = Prout({ 
3.do({ arg i; 3.rand.yield; }) 
});
)
// make two streams from the pattern
b = a.asStream;
c = a.asStream;
4.do({ b.next.postln; }); // print 4 values from first stream（应该只有前三个有值）
4.do({ c.next.postln; }); // print 4 values from second stream
```

如果 stream 耗尽了需要重新播放，用同一个 Pattern 实例重新 asStream 即可

## PSeq

（Pattern-Sequence）

`ListPatterns` 是以某种方式遍历数组的 Patterns，最代表性的是 Pseq

```scd
a = Pseq([1, 2, 3], 2); // repeat twice
b = a.asStream;
7.do({ b.next.postln; });
```

其他 ListPatterns：

- Pser
- Prand
- Pxrand
- Pshuf

# EventPattern

之前介绍的 Pattern 都属于 ValuePattern，形成的 Stream 是每次返回单个值的

还有另一种 Pattern 名为 EventPattern，形成的 Stream 每次返回一个事件（Event）


- Value Stream
	- 来源：Pseq, Prand, Pwhite 等
	- 原始值
	- 只能手动遍历
- Event Stream
	- 来源：Pbind, Pmono, Pdef 等
	- 完整事件
	- 由于知道 \\dur，所以可以由调度器自动播放

## Pbind

Pbind 是两种 Pattern 的桥梁，它的参数是成对的 Symbol 和 Pattern，生成一个 EventStream（可直接 play 播放声音）

```scd
Pbind(
	\degree, Pseq([0, 2, 4], inf)
	\dur, Pseq([0.25, 0.5], inf)
).play
```

上述例子会形成事件流：

- degree 0，dur 0.25
- degree 2，dur 0.5
- degree 4，dur 0.25
- ... 以此类推

由于知道每个事件的持续时间和和音调，所以可以播放出来

