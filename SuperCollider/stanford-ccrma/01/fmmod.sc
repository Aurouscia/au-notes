/*
	***********************************************************************************
	fmmod.sc
	A very simple FM modulation instrument using two sinusoids.
	
	Reference: Computer Music Tutorial, Figure 6.9
	Juan-Pablo Caceres, January 09, 2007
	Under Subversion Control System 
	***********************************************************************************
*/

s = Server.default.boot;

//SynthDef (Definition of the instrument)
//--------------------------------------
(
SynthDef("fmmod",
	{
		// Define arguments and default
		arg out=0, gate = 1.0, amp = 0.1,
		carrier_freq = 440.0, carrier_amp = 1.0,
		mod_freq = 1.0, mod_amp = 1.0;

		var modulator, fmmod;
		// Modulation Oscilator
		modulator = SinOsc.ar(
			freq: mod_freq, 
			mul: mod_amp
		);
		// Carrier Oscilator with Frequency Multiplication
		fmmod = SinOsc.ar(
			freq: carrier_freq + modulator, 
			mul: EnvGen.kr(Env.adsr, gate, doneAction:2) * amp
		);

		Out.ar(out, fmmod);
	}).load(s);
)


// Let's play someting
//-------------------
~fm = Synth("fmmod",[\mod_freq, 1.0, \mod_amp, 50.0]); //play the default
~fm.set(\mod_freq, 2);   //increase the tremollo
~fm.set(\mod_freq, 120); //spectrum
~fm.set(\gate, 0);       //stop and free

// Use the mouse to control mod_freq and modulation index
//------------------------------------------------------
(
SynthDef("Mousex", { arg out=0; 
	Out.kr(out, MouseX.kr(0, 250)) 
}).send(s);
SynthDef("Mousey", { arg out=0; 
	Out.kr(out, MouseY.kr(0, 300)) 
}).send(s); 
)

(
//Do the Bus Mappings
~busx = Bus.control(s);//Obtain a control bus dinamically
~busy = Bus.control(s);//Obtain a control bus dinamically
Synth("Mousex", [\out, ~busx.index]);
Synth("Mousey", [\out, ~busy.index]);

//play it
~fm = Synth("fmmod",[\mod_freq, 1.0, \mod_amp, 50.0]);
~fm.map(\mod_freq, ~busx.index); //moseX controls Mod. Freq.
~fm.map(\mod_amp, ~busy.index); //mouseY controls Mod. Amp.
)

//free when done playing with the mouse
~fm.set(\gate, 0);
