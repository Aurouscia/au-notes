/*
***********************************************************************************
ringmod.sc
A very simple ring modulation instrument using two sinusoids.

Reference: Computer Music Turorial, Figure 6.2
Juan-Pablo Caceres/Fernando Lopez-Lezcano, January 2007
Under Subversion Control System
***********************************************************************************
*/

s = Server.default.boot;

// this is a slightly simpler version, we just add the modulation
// directly by using the "mul:" parameter of the carrier oscillator

(
SynthDef("ringmod2",
	{
		arg out=0, gate = 1.0, amp = 0.4,
		carrier_freq = 440.0, carrier_amp = 1.0,
		mod_freq = 1.0, mod_amp = 1.0;

		Out.ar(out, SinOsc.ar(
			freq: carrier_freq,
			mul: SinOsc.ar(freq: mod_freq, mul: mod_amp) *
			     EnvGen.kr(Env.adsr, gate, doneAction:2) *
			     amp * carrier_amp)
		);
	}).load(s);
)


// get a note started with the default parameters...

~rm = Synth("ringmod2");

// and then change parameters in the running synth

~rm.set(\mod_freq, 2);
~rm.set(\mod_freq, 20);
~rm.set(\mod_freq, 80);

// finally stop and free the synth

~rm.set(\gate, 0);

// we can also override parameters when we create the synth
// just supply an array with pairs of parameter name and value

~rm = Synth("ringmod", [\carrier_freq, 450, \mod_freq, 110]);

// create an instrument that outputs mouse events to a bus

(
SynthDef("Mousex", { arg out=0;
	Out.kr(out, MouseX.kr(0, 50))
}).send(s);
)

(
// start a mousex synth and send the output to a bus that
// is dynamically allocated

~busx = Bus.control(s);
Synth("Mousex", [\out, ~busx.index]);

// play it

~rm = Synth("ringmod2");
~rm.map(\mod_freq,~busx.index);
)

// free when done playing with the mouse

~rm.set(\gate, 0);
