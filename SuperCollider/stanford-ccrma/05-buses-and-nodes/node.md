# node

```sc
(
//定义顺序任意
SynthDef("Saw", { arg filterBus = 16;
Out.ar(filterBus, LFSaw.ar([60, 90]))
}).send(s);

SynthDef("Filter", {arg filterBus = 16;
Out.ar(0, RLPF.ar(In.ar(filterBus, 2), LFNoise0.kr(12, 300, 350), 0.2))
}).send(s);
)

// 不会有声音
Synth("Saw"); // 源（写入 16 号 bus）
Synth("Filter"); // filter（读取 16 号 bus）

// cmd+.

// 会有声音
Synth("Filter"); //filter
Synth("Saw"); //source

// cmd+.

// 或者写成：
Synth("Filter"); Synth("Saw");

// 查看当前节点
s.queryAllNodes;
```

顺序不同造成声音有无的原因：

后一种写法中，Saw（后创建的 Node）在 HEAD，而 Filter 在 TAIL

audio信号的**数据会从 HEAD 流向 TAIL**，所以这种顺序管用，反过来不管用

使用 Synth 下的 head 或 tail 方法，手动指定顺序：

```sc
Synth.tail(s,"Filter");
Synth.head(s,"Saw");

// 调用顺序无所谓，head 和 tail 中的名字是对的就行

Synth.head(s,"Saw");
Synth.tail(s,"Filter");
```

或者，创建时指定插入到哪里：

```sc
(
x = Synth("Saw");
y = Synth("Filter", addAction: \addAfter, target: x);
y;
)
```

## 自动分配 bus

```sc
// 自动分配空闲的 audio bus
c = Bus.audio(s,2);

// 按顺序执行（先创建的 node 在 TAIL，后创建的 node 在 HEAD）
// 数据由 Out.ar（写入 bus）流向 In.ar（读取 bus）
{In.ar(c.index) ! 2}.play
{Out.ar(c.index, 0.1*Saw.ar)}.play;
```