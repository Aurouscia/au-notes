// Michael McNabb's fm bell:

(
SynthDef(\fmbell, {
    // dur：音符持续时间
    // freq：基频（音高）
    // amplitude：音量大小
    // index：FM 调制指数（控制泛音的丰富程度/亮度）

    arg dur, freq, amplitude, index,
        ampenv = #[0, 1, 0.6, 0.3, 0.15, 0.1, 0],      // 电平（带默认值）
        amptimes = #[0.1, 9.9, 15, 25, 40, 10],        // 电平时间（数组长度是电平-1）
        indexenv = #[1, 1.1, 0.75, 0.5, 0.2],
        indtimes = #[2, 23, 50, 25];

    // mod1-4：存储 4 个调制振荡器（Modulator）的信号
    // fmind1-4：存储 4 个调制指数相关的频率值
    // ampf, indf：分别存储振幅包络和调制指数包络的信号
    // out：最终输出的音频信号

    var mod1, mod2, mod3, mod4;
    var fmind1, fmind2, fmind3, fmind4;
    var ampf, indf, out;
    
    // 调制器
    // 创建了 4 个正弦波振荡器，频率是基频的倍数：
    // mod1：2 倍频（八度）
    // mod2：约根号2倍频（非整数比，产生不协和的金属感）
    // mod3：2.82 倍频（接近 2 倍根号2，与 mod2 配合）
    // mod4：2.4 倍频
    // 这些信号将用于调制载波（Carrier）的频率

    mod1 = SinOsc.ar(freq * 2);
    mod2 = SinOsc.ar(freq * 1.41);
    mod3 = SinOsc.ar(freq * 2.82);
    mod4 = SinOsc.ar(freq * 2.4);

    // 调制指数
    // 这部分计算了每个调制器对应的“频率偏移量”或“调制深度”。
    // fmind1：与基频成正比（32×freq ），高频音会有更宽的频偏。
    // fmind2：随着频率升高而略微降低（8−freq/50 ）。
    // fmind3：基于 fmind2 进一步调整，带有复杂的比例系数。
    // fmind4：随频率升高而显著降低（20−freq/20 ），防止高频声音过于刺耳。
    fmind1 = freq * 32.0;
    fmind2 = 4.0 * (8.0 - (freq / 50.0));
    fmind3 = fmind2 * 0.705 * (1.4 - (freq / 250.0));
    fmind4 = 32.0 * (20 - (freq / 20));

    // 使用 indexenv 和 indtimes 构建 index 包络
    indf = EnvGen.kr(
        Env.new(indexenv, indtimes, 'linear'),
        levelScale: index,
        timeScale: dur / 100,
        doneAction: 2
    );

    // 使用 ampenv 和 amptimes 构建音量包络
    ampf = EnvGen.kr(
        Env.new(ampenv, amptimes, 'linear'),
        levelScale: amplitude,
        timeScale: dur / 100,
        doneAction: 2
    );

    // 声音合成
    out = ampf * (
        SinOsc.ar(freq + (mod1 * indf * fmind1))
        + (0.15 * (SinOsc.ar(freq + (indf * ((mod2 * fmind2) + (mod3 * fmind3))))))
        + (0.15 * (SinOsc.ar(freq * 2.4 + (indf * fmind4 * mod4))))
    );

    // ! 2 是“复制为两份的数组”的意思，即填满左右两个声道，避免只有左耳有声音
    Out.ar(0, out ! 2);
}).send(s);
)

// 使用默认包络（和原来一样）
Synth(\fmbell, [\dur, 5.0, \freq, 233.046, \amplitude, 0.028, \index, 0.75]);

// 自定义振幅包络：更尖锐的起音，更快的衰减
Synth(\fmbell, [
    \dur, 2.0,
    \freq, 440,
    \amplitude, 0.1,
    \index, 1.5,
    \ampenv, #[0, 1, 0.8, 0],           // 电平：快速起音，立即衰减
    \amptimes, #[0.01, 0.5, 1.49]       // 时间：极短起音，短衰减
]);

// 自定义调制指数包络：让声音逐渐变亮而不是变暗
Synth(\fmbell, [
    \dur, 3.0,
    \freq, 220,
    \amplitude, 0.05,
    \index, 1.0,
    \indexenv, #[0.2, 1.2, 0.2],        // 电平：从暗到亮再到暗
    \indtimes, #[30, 40, 30]            // 时间分布
]);

Synth(\fmbell, [\dur, 5.000, \freq, 233.046, \amplitude, 0.028, \index, 0.750]);
Synth(\fmbell, [\dur, 2.000, \freq, 205.641, \amplitude, 0.19, \index, 0.650]);
Synth(\fmbell, [\dur, 5.000, \freq, 207.152, \amplitude, 0.17, \index, 0.750]);
Synth(\fmbell, [\dur, 7.000, \freq, 205.641, \amplitude, 0.010, \index, 0.650]);
Synth(\fmbell, [\dur, 0.500, \freq, 880, \amplitude, 0.060, \index, 0.500]);
Synth(\fmbell, [\dur, 6.500, \freq, 293.66, \amplitude, 0.1, \index, 0.500]);
Synth(\fmbell, [\dur, 7.000, \freq, 146.5, \amplitude, 0.1, \index, 1.000]);
Synth(\fmbell, [\dur, 6.000, \freq, 110, \amplitude, 0.1, \index, 1.000]);
Synth(\fmbell, [\dur, 10.00, \freq, 73.415, \amplitude, 0.028, \index, 0.500]);
Synth(\fmbell, [\dur, 4.000, \freq, 698.46, \amplitude, 0.068, \index, 0.500]);