// Digital Filters
// (c) 2007 Fernando Lopez-Lezcano

// we will be looking at different digital filters with examples
// using the correspongin SuperCollider UGens.

// for comparison purposes here's unfiltered white noise:

{ WhiteNoise.ar([0.5, 0.5]) }.play;

{ PinkNoise.ar([0.5, 0.5]) }.play;

// we start with the simplest possible digital filters using
// a delay of just one sample

// OneZero
// A one zero filter. Implements the formula :
// out(i) = ((1 - abs(coef)) * in(i)) + (coef * in(i-1))


{ OneZero.ar(WhiteNoise.ar([0.5, 0.5]), 0.5) }.play

{ OneZero.ar(WhiteNoise.ar([0.5, 0.5]), Line.kr(-0.5, 0.5, 10)) }.play

{ OneZero.ar(WhiteNoise.ar([0.5, 0.5]), MouseX.kr(-0.5, 0.5)) }.play


// OnePole
// A one pole filter. Implements the formula :
// out(i) = ((1 - abs(coef)) * in(i)) + (coef * out(i-1))

{ OnePole.ar(WhiteNoise.ar([0.5, 0.5]), MouseX.kr(-0.99, 0.99)) }.play

// or we could use FOS (first order filter section) and manipulate
// all coefficients without any restriction (ie: convert it to a
// OnePole, OneZero, etc, etc)


// TwoPole
// A two pole filter. 

{ TwoPole.ar(WhiteNoise.ar(0.005), MouseX.kr(800,8000,1), 0.95) }.play

{ TwoPole.ar(WhiteNoise.ar([0.05, 0.05]), MouseX.kr(800,8000,1), MouseY.kr(0.4, 0.999)) }.play


// RLPF
// a resonant low pass filter

{ RLPF.ar(WhiteNoise.ar([0.5, 0.5]), MouseX.kr(800,8000,1), 0.95) }.play

// ... see also RHPF
// ... see also BPF (2nd order Butterworth bandpass filter)
// ... see also BRF (2nd order Butterworth bandreject filter)

// or we could use SOS (standard second order filter section - biquad)
// and manipulate all coefficients directly. Problem: the coefficients
// do not translate directly into perceptually related controls
// for example (same as TwoPole):

(
{
	var rho, theta, b1, b2;
	theta = MouseX.kr(0.2pi, pi);
	rho = MouseY.kr(0.6, 0.99);
	b1 = 2.0 * rho * cos(theta);
	b2 = rho.squared.neg;
	SOS.ar(LFSaw.ar(200, 0, 0.1), 1.0, 0.0, 0.0, b1, b2)
}.play
)

// Resonz, a two pole, two zero tunable resonator
// A two pole resonant filter with zeroes at z = +/- 1. Based on
// K. Steiglitz,  "A Note on Constant-Gain Digital Resonators," 
// Computer Music Journal, vol 18, no. 4, pp. 8-10, Winter 1994.

// modulate frequency
{ Resonz.ar(WhiteNoise.ar([0.5, 0.5]), XLine.kr(1000,8000,10), 0.05) }.play

{ Resonz.ar(WhiteNoise.ar([0.5, 0.5]), MouseX.kr(500,10000,1), 0.05) }.play

// modulate bandwidth
{ Resonz.ar(WhiteNoise.ar([0.5, 0.5]), 2000, XLine.kr(1, 0.001, 8)) }.play

{ Resonz.ar(WhiteNoise.ar([0.5, 0.5]), 2000, MouseX.kr(1, 0.001, 1)) }.play

// both
{ Resonz.ar(WhiteNoise.ar([0.5, 0.5]), MouseX.kr(500,10000,1), MouseY.kr(1, 0.001, 1)) }.play

// ... see also Ringz (ringing filter, Q specified as decay time)
// ... see also DynKlank (dynamic bank of resonators)
// ... see also Formlet (FOF-like filter)
//                        --> Formlet examples...

// visit:
// /usr/share/SuperCollider/Help/UGens/Filters
// for help files on many more SC UGens available for filtering...

// From additional SC libraries we pull in more stuff:


// MoogVCF - a classic
//
// MoogVCF is a port of the CSound opcode moogvcf written by Josep
// Comajuncosas and Hans Mikelson.  This UGen is a digital emulation of
// the Moog diode ladder filter configuration.  It is based loosely on
// the paper "Analyzing the Moog VCF with Considerations for Digital
// Implemnetation" by Stilson and Smith (CCRMA).
//
// 4 pole, 24 db/oct low pass filter with resonance

// see Tim's page here:
// http://ccrma.stanford.edu/~stilti/papers/Welcome.html
// and his paper here:
// http://ccrma.stanford.edu/~stilti/papers/moogvcf.pdf


{ MoogVCF.ar(Pulse.ar([40,121],[0.3,0.7]),SinOsc.kr(LFNoise0.kr(0.42).range(0.001,2.2)).range(30,4200),	0.83)}.play;

{ MoogVCF.ar(WhiteNoise.ar([0.5, 0.5]), MouseX.kr(100, 10000), MouseY.kr(0.3, 0.99)) }.play


// See also the BEQsuite collection of preconfigured biquad digital filters...
// info and examples here:
//   /usr/share/SuperCollider/Help/BEQSuiteHelp/

// BLowPass (12db/oct low pass)

(
z = {
	BLowPass.ar(
		WhiteNoise.ar([0.5, 0.5]),
		MouseX.kr(10,20000,\exponential), // cutoff freq.
		MouseY.kr(0.0, 1.0, \linear), // rq
		0.5); // mul
}.play)
z.release;

// BBandPass (band pass)

(
z = {
BBandPass.ar(
	WhiteNoise.ar(0.5),
	MouseX.kr(20,20000,\exponential),
	MouseY.kr(0.0, 10.0, \linear), // bw
	0.5); // mul
}.play)
z.release;

// BLowShelf 

(
z = { // toy around with boost/cut
BLowShelf.ar(
	WhiteNoise.ar(0.5),
	MouseX.kr(40,6000,\exponential),
	1.0,  // rs
	MouseY.kr(24.0, -24.0, \linear),
	0.5); // mul
}.play)
z.release;

(
z = { // toy around with rs
BLowShelf.ar(
	WhiteNoise.ar(0.5),
	MouseX.kr(20,6000,\exponential),
	MouseY.kr(0.1, 1.0, \linear),  // rs
	6,
	0.5); // mul
}.play)
z.release;
