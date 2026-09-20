/*
	***********************************************************************************
	ringmod.sc
	A very simple ring modulation instrument using two sinusoids.

	Reference: Computer Music Turorial, Figure 6.2
	Juan-Pablo Caceres, December 12, 2006
	Under Subversion Control System
	***********************************************************************************
*/

s = Server.default.boot;

//SynthDef (Definition of the instrument)
//--------------------------------------
(
SynthDef("ringmod",
	{
		// Define arguments and default
		arg out=0, gate = 1.0, amp = 0.4,
		carrier_freq = 440.0, carrier_amp = 1.0,
		mod_freq = 1.0, mod_amp = 1.0;

		var modulator, ringmod;
		// Modulation Oscilator
		modulator = SinOsc.ar(
			freq: mod_freq,
			mul: mod_amp
		);
		// Carrier Oscilator with Ring Modulation
		ringmod = SinOsc.ar(
			freq: carrier_freq,
			mul: EnvGen.kr(Env.adsr, gate, doneAction:2) * amp * carrier_amp * modulator
		);

		Out.ar(out, ringmod);
	}).load(s);
)

// Let's play someting

s.scope

~rm = Synth("ringmod");  // default arguments
~rm.set(\mod_freq, 2);   // increase the tremolo speed
~rm.set(\mod_freq, 120); // spectrum
~rm.set(\gate, 0);       //stop and free

// Use the mouse to control mod_freq
//---------------------------------

// we could incorporate mouse control directly into the synthdef
// but here we route a the mouse through a bus, so that we can
// dynamically control any parameter without having to change
// our synthdef

// you could also start jaaa, a jack based spectrum analyser and connect
// the outputs of supercollider to it (use qjackctl's "Connect"
// button to bring up the connection patchbay), you will be able
// to see the spectrum changes in realtime

(
SynthDef("Mousex", { arg out=0;
	Out.kr(out, MouseX.kr(0, 550))
}).send(s); // this will send it to the server without saving it.
)

(
// Do the Bus Mappings
// first obtain a control bus dinamically
~busx = Bus.control(s);
// and then send the output of the mouse to it
Synth("Mousex", [\out, ~busx.index]);

// start the ring modulator instrument
~rm = Synth("ringmod");
// map the bus output to the mod_freq parameter
~rm.map(\mod_freq,~busx.index);
)

// free when done playing with the mouse
~rm.set(\gate, 0);
