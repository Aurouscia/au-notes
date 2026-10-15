//---------------------------------------
// (C) Fernando Lopez-Lezcano 2007

// 基于
//   Notes from the Meta Level (Rick Taube)
//   取自第 9 章的示例：
//     "Etudes, Op3: Iteration, Rows and Sets"
// 以及
//   David Cottle 的书

// 请查阅以下类的帮助文档：
//   Array
//   List


// 把音级编号(midi 音高)转换为音级(pitch class)

// 数组支持二元运算符，所以我们直接用 % (取模)
// 把 midi 音高转换为音级(音高中的
// 八度信息被剥离，只保留 0-11 的音级)

(
[55, 58, 62, 66, 69, 72, 76, 80, 83, 85, 86, 89] % 12;
)

// 我们也可以创建一个辅助函数，
// 赋值给一个恰当命名的全局变量，供程序中使用。
// 它对 List 和 Array 都有效

(
~list_to_pcs = {
	arg array;
	array % 12;
}
)

a = ~list_to_pcs.value([55, 58, 62, 66, 69, 72, 76, 80, 83, 85, 86, 89]);

// 音级归一化(normalize)
// 把所有音程重新定位，使第一个音级变为 0

a;
a[0];
b = (a - a[0]);
b = (a - a[0]) % 12;

(
~normalize_pcs = {
	arg row;
	(row - row[0]) % 12;
}
)

(
~normalize_pcs = {
	arg row;
	~list_to_pcs.value(row - row[0]);
}
)

~normalize_pcs.value(a);


// 基本的十二音变换

// 逆行(retrograde)
// 对集合直接使用 reverse 方法即可(见 [Array] 类的帮助文档)

b = a.reverse;

(
~retrograde_row = {
	arg row;
	row.reverse;
}
)

// 移调(transpose)

b = (a + 5) % 12;

(
~transpose_row = {
	arg transposition, row;
	(row + transposition) % 12;
}
)
(
~transpose_row = {
	arg transposition, row;
	~list_to_pcs.value(row + transposition);
}
)

b = ~transpose_row.value(a, 6);
b.postln;

// 倒影(invert)

b = 11 - a;

(
~invert_row = {
	arg row;
	11 - row;
}
)

b = ~invert_row.value(a);

// 逆行 + 倒影

b = ~retrograde_row.value(~invert_row.value(a));

// 我们可以用 Array 的一些方法轻松创建一个十二音序列
// 用这种简写记号创建一个包含全部 12 个音级的数组

a = (0..11);

// 然后可以用 scramble 创建全部 12 个音级的随机排列

b = a.scramble;

// 例如，创建一个包含十二条随机音列的列表
(
var a = (0..11);
var l = List.new;
12.do {
	l.add(a.scramble);
};
l.do( { arg item, i; item.postln; } );
l.size
)

// 由这些被操作的数组，我们可以轻松创建 pattern(模式)，
// 必要时还可以转换为 stream(流)

c = Pseq(a);
d = c.asStream;
d.next;


// Cottle 26.11. 序列主义(Serialism)

// Cottle 示例乐器

(
SynthDef("beep",
	{arg midinote = 60, amp = 0.9, dur = 1, pan = 0, legato = 0.8;
		Out.ar(0,
			Pan2.ar(
				SinOsc.ar(midinote.midicps, mul: amp)
				*
				EnvGen.kr(Env.perc(0, dur*legato), doneAction: 2),
				pan
			)
		)
	}).store;
)

SynthDescLib.global.read;

// Cottle 26.11. 序列主义

(
a = [0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11].scramble.postln;
o = [48, 60, 72];
Pbind(
	\instrument, "beep",
	\midinote, Prand(
		[ //P 原型, R 逆行, I 倒影, IR 逆行倒影
			Pseq(a) + o.choose,
			Pseq(a.reverse) + o.choose,
			Pseq(11 - a) + o.choose,
			Pseq((11 - a).reverse) + o.choose
		], inf),
	\dur, 0.2
).play;
)

// 使用我们自己定义的函数...
(
a = [0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11].scramble.postln;
o = [48, 60, 72];
Pbind(
	\instrument, "beep",
	\midinote, Prand(
		[ //P 原型, R 逆行, I 倒影, IR 逆行倒影
			Pseq(a) + o.choose,
			Pseq(~retrograde_row.value(a)) + o.choose,
			Pseq(~invert_row.value(a)) + o.choose,
			Pseq(~retrograde_row.value(~invert_row.value(a))) + o.choose
		], inf),
	\dur, 0.2
).play;
)

// 加入时值和力度(amplitude)模式
(
a = [0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11].scramble.postln;
r = [0.1, 0.1, 1.0, 0.2, 0.3, 0.166, 0.166];
o = [48, 60, 72];
Pbind(
	\instrument, "beep",
	\midinote, Prand(
		[ //P 原型, R 逆行, I 倒影, IR 逆行倒影
			Pseq(a) + o.choose,
			Pseq(a.reverse) + o.choose,
			Pseq(12 - a) + o.choose,
			Pseq((12 - a).reverse) + o.choose
		], inf),
	\dur, Pseq([Prand([0.1, 0.2, 0.5, 1.0], 7),
		Prand([Pseq(r), Pseq(r.reverse)], 1)], inf),
	\amp, Prand([
		Pseries(0.1, 0.1, 5), // 渐强
		Pseries(0.9, -0.1, 6), // 渐弱
		Prand([0.1, 0.3, 0.5, 0.7], 5)
	], inf)
).play;
)

(
// 当然，也可以三个声部同时进行。
// 如果 seed(随机种子)是不为 0 的任意数字，就使用该种子。
// 如果为 0，则随机挑选一个种子并打印出来。把它记下来，
// 就可以重复某一次演奏的结果。
var seed = 0;
if(seed !=0, {thisThread.randSeed = seed},
	{thisThread.randSeed = Date.seed.postln});
a = [0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11].scramble.postln;
r = [0.1, 0.1, 1.0, 0.2, 0.3, 0.166, 0.166];
o = [48, 60, 72];
Pbind(
	\instrument, "beep",
	\pan, -1,
	\midinote, Prand(
		[ //P 原型, R 逆行, I 倒影, IR 逆行倒影
			Pseq(a) + o.choose,
			Pseq(a.reverse) + o.choose,
			Pseq(12 - a) + o.choose,
			Pseq((12 - a).reverse) + o.choose
		], inf),
	\dur, Pseq([Prand([0.1, 0.2, 0.5, 1.0], 7),
		Prand([Pseq(r), Pseq(r.reverse)], 1)], inf),
	\amp, Prand([
		Pseries(0.1, 0.1, 5), // 渐强
		Pseries(0.9, -0.1, 6), // 渐弱
		Prand([0.1, 0.3, 0.5, 0.7], 5)
	], inf)
).play;
Pbind(
	\instrument, "beep",
	\pan, 0,
	\midinote, Prand(
		[ //P 原型, R 逆行, I 倒影, IR 逆行倒影
			Pseq(a) + o.choose,
			Pseq(a.reverse) + o.choose,
			Pseq(12 - a) + o.choose,
			Pseq((12 - a).reverse) + o.choose
		], inf),
	\dur, Pseq([Prand([0.1, 0.2, 0.5, 1.0], 7),
		Prand([Pseq(r), Pseq(r.reverse)], 1)], inf),
	\amp, Prand([
		Pseries(0.1, 0.1, 5), // 渐强
		Pseries(0.9, -0.1, 6), // 渐弱
		Prand([0.1, 0.3, 0.5, 0.7], 5)
	], inf)
).play;
Pbind(
	\instrument, "beep",
	\pan, 1,
	\midinote, Prand(
		[ //P 原型, R 逆行, I 倒影, IR 逆行倒影
			Pseq(a) + o.choose,
			Pseq(a.reverse) + o.choose,
			Pseq(12 - a) + o.choose,
			Pseq((12 - a).reverse) + o.choose
		], inf),
	\dur, Pseq([Prand([0.1, 0.2, 0.5, 1.0], 7),
		Prand([Pseq(r), Pseq(r.reverse)], 1)], inf),
	\amp, Prand([
		Pseries(0.1, 0.1, 5), // 渐强
		Pseries(0.9, -0.1, 6), // 渐弱
		Prand([0.1, 0.3, 0.5, 0.7], 5)
	], inf)
).play;
)
