# bus

SC 中有两种总线，音频和控制（ar 和 kr）

查看当前 SC 实例有多少个总线（不同版本中这个值可能不同）：

```sc
o = Server.default.options;
o.numAudioBusChannels;
o.numControlBusChannels;
```

总线是数组中的元素，SC 用它们来表示音频和控制信号

可以用 Out.ar 和 In.ar 连接音频总线

使用如下代码查看音频输入和输出共有几个总线：

```sc
o = Server.default.options;
o.numInputBusChannels;
o.numOutputBusChannels;
```

一般的家用电脑（配置好了聚合设备的 mac）应该是各2个

## out

如果各自2个，那么前两个是 out（0、1），后两个是 in（2、3）

以下代码会分别向左耳和右耳播放正弦波

```sc
{Out.ar(0, 0.1*SinOsc.ar)}.play;
{Out.ar(1, 0.1*SinOsc.ar)}.play;
```

## in

以下代码会把两个麦克风分别透传到两个喇叭（默认情况下可能只有 0）

```sc
{Out.ar([0,1], In.ar([2,3]))}.play;
```

## 控制信号

例子：

```sc
// 完整示例：创建一个左10、右15的 LFNoise0，把 [-1,1] 映射到 [400,1200]，作为 SinOsc 的参数，形成左右分别随机的音频
{SinOsc.ar(LFNoise0.kr([10, 15]) * 400 + 800, 0, 0.3)}.play;

// 只播放稳定的正弦波
{SinOsc.ar(800, 0, 0.3)}.play;

// 只播放 LFNoise0，形成左右声道独立的“噗噗”阶跃声
{LFNoise0.ar([10, 15])}.play;
```

改为使用 buses：

```sc
// 用 Out.kr 指定向 20 号控制总线发送 LFNoise0
{Out.kr(20, LFNoise0.kr([8, 11], 500, 1000))}.play;

// 用 In.kr 从 20 号控制总线接收 LFNoise0，并用于 SinOsc 参数
{Out.ar(0, SinOsc.ar(In.kr(20, 2), 0, 0.3))}.play
```

动态控制使用哪条控制信号：

```sc
(
// 启动所有 synths
SynthDef("LFN0Control",
{Out.kr(20, LFNoise0.kr([8, 11], 500, 1000))}).play(s);
SynthDef("SineControl",
{Out.kr(22, SinOsc.kr([3, 3.124], mul: 500, add: 1000))}).play(s);
SynthDef("MouseControl",
{Out.kr(24, MouseX.kr([100, 200], 1000))}).play(s);

// 此时 a 尚未连接，不会发出声音
// 后续为 a 的 \busInNum 设置有信号的 bus 后，才会发出声音
a = SynthDef("PatchableSine", {arg busInNum = 0;
Out.ar(0, SinOsc.ar(In.kr(busInNum, 2), 0, 0.3))}).play(s);
)

a.set(\busInNum, 20); // 设为 LFNoise0
a.set(\busInNum, 22); // 切换为 SineControl
a.set(\busInNum, 24); // 切换为 MouseControl
```