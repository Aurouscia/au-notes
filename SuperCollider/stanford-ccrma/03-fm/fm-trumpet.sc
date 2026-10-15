// FM 小号

// FM(频率调制)小号的实现
// 基于 Dodge 一书第 129 页，改编自 Dexter Morrill 的设计

(
SynthDef(\fmtrumpet, {arg dur = 1, amp = -3, fundfreq = 440, index = 1, halfsine = 1;
	var env, amp1, amp2, formant, dev1, dev2, rand, vibamp, vibosc, porta, vib, ac1env,
		ac2env, modenv, car1 ,car2, mod, mod2;
	env = EnvGen.kr(Env([0.01, 1, 1, 0.01], [0.01, dur - 0.02, 0.01], \exp), doneAction: 2);
	amp1 = amp.dbamp * 0.7; // 第一载波(基频载波)的振幅
	amp2 = amp1 * 0.2; // 第二载波(共振峰载波)的振幅。
					// 正是这个第二载波用来模拟共振峰(formant)

// 下面的公式是计算频率的标准方法：
// 它能让频率尽可能接近目标共振峰频率，
// 同时又是基频的整数倍谐波分量。小号最重要的
// 共振峰大约在 1500 Hz。因此我们希望在 1500 Hz
// 附近增加能量，以模拟共振峰，从而合成出
// 类似小号的声音。为此我们使用第二个载波
// 振荡器，其频率对应于共振峰。小号是
// 谐波类乐器，所以共振峰也必须是谐波。
// 例如，假设基频是 440 Hz。不难看出
// 1500 Hz 并不是 440 Hz 的谐波分量(1500 不是 440 的整数
// 倍)。把 440 代入下面的公式：

//	 formant =  (1500/440).round(1) * 440

// 得到的 formant(共振峰频率)结果是 1320。"round" 函数：

//	 a.round(b)

// 返回把 a 舍入到最接近的 b 的整数倍后的值。所以 (1500/440).round(1)
// 会返回 3 而不是 3.409。
// 这样做是因为我们要用一个整数去乘基频，
// 使得到的频率成为谐波分量。
// 然后我们把基频乘以该整数，就得到
// 该谐波分量的频率(以 Hz 为单位)，如上面的 1320 Hz。如果
// 你换成自己的基频和共振峰数值代入计算，
// 看看结果，会帮助你更好地理解这个过程！

	formant = (1500 / fundfreq).round(1) * fundfreq;

// 变量 "dev1" 是调制振荡器的偏移量(即调制幅度)。
// index 是在乐谱(score)中设定的一个类似调制指数的量，
// 用来让声音更明亮或更暗淡。

	dev1 = (fundfreq * 3) * index;

// 变量 dev2 用来在调制信号送入第二载波之前
// 对其进行缩放。因为我们不希望在第二载波(共振峰载波)
// 周围产生那么多边带，所以偏移量应当更小。

	dev2 = 0.666 * index;

// 颤音(VIBRATO)

// 使用带插值的 LFNoise1 产生的随机分量

	rand = LFNoise1.ar(15, 0.007);

// 颤音分量的幅度。这里从下方代码分配的、内存中存储的一个
// 半正弦形状的 buffer(缓冲区)中读取。为了
// 确保整个音符期间只读一遍，它的读取频率
// 设为音符时长的倒数。

	vibamp = Osc.kr(halfsine, dur.reciprocal, 0, 0.007);

// 颤音分量的主振荡器

	vibosc = SinOsc.ar(5, 0, vibamp);

// 音高滑音(portamento)函数

	porta = EnvGen.kr(Env([0, 0.03, 0.02], [0.06, dur - 0.06], \lin));

// 把这三个分量相乘，并加上 1，
// 得到一个以 1 为中心的频率倍乘系数

	vib	= (1 + rand) * (1 + vibosc) * (1 + porta);

// 包络(ENVELOPES)

	ac1env = EnvGen.kr(Env([0.001, 1, 0.8, 0.001], [0.1, dur - 0.25, 0.15], \exp));
	ac2env = EnvGen.kr(Env([0.001, 1, 0.8, 0.001], [0.1, dur - 0.45, 0.3], \exp));
	modenv = EnvGen.kr(Env([0.001, 1, 0.8, 0.001], [0.1, dur - 0.11, 0.01], \exp));

// 调制器(MODULATOR)

// 载波与调制的比 C:M 是 1:1，所以 fm = fc * 颤音系数

	mod = SinOsc.ar(fundfreq * vib, 0, dev1 * modenv);

// 载波1(CARRIER1)

// fc*vib 给出基频，然后
// 加上调制信号

	car1 = SinOsc.ar((fundfreq * vib) + mod, 0, ac1env * amp1);

// 在把调制信号施加到共振峰载波之前，
// 先按比例缩小调制偏移量，从而
// 产生更少的边带

	mod2 = mod * dev2;

// 共振峰载波的 C:M 比与
// 基频调制器的不同，但由于
// 共振峰频率与 fc 成谐波关系，所得的 共振峰:M
// 比同样是谐波关系，而在约分后的比例中
// M 仍为 1 —— 因此产生的频谱成分
// 完全相同。

// 载波2(CARRIER2)

	car2 = SinOsc.ar((formant * vib) + mod2, 0, amp2 * ac2env);

	Out.ar(0, (car1 + car2) * env ! 2);
	}).load(s);
)



s = Server.internal.boot;

// 分配一个存放半正弦波形的 buffer(0 上升到 1 再回落到 0)，
// 供上面 Osc.kr 读取，作为颤音的幅度包络
~halfsine = Buffer.loadCollection(s, (0..1023).collect({ |i| sin(pi * i / 1023) }));

Synth(\fmtrumpet, [\dur, 0.5, \amp, -3, \fundfreq, 440, \index, 1, \halfsine, ~halfsine.bufnum]);

Synth(\fmtrumpet, [\dur, 1, \amp, -3, \fundfreq, 522, \index, 1, \halfsine, ~halfsine.bufnum]);

Synth(\fmtrumpet, [\dur, 1, \amp, -3, \fundfreq, 440, \index, 1.5, \halfsine, ~halfsine.bufnum]);

Synth(\fmtrumpet, [\dur, 2.5, \amp, -3, \fundfreq, 522, \index, 1.5, \halfsine, ~halfsine.bufnum]);

