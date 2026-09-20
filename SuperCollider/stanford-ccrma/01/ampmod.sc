/*
	***********************************************************************************
	ampmod.sc
	A very simple amplitude modulation instrument using two sinusoids.
	
	Reference: Computer Music Turorial, Figure 6.8
	Juan-Pablo Caceres, January 09, 2007
	Under Subversion Control System 
	***********************************************************************************
*/

s = Server.default.boot;

(
SynthDef("ampmod",
	{
		// Define arguments and default
		arg out=0, gate = 1.0, amp = 0.1,
		carrier_freq = 440.0, carrier_amp = 1.0,
		mod_freq = 1.0, mod_amp = 1.0;

		var modulator, ampmod;
		//Modulation Oscilator
		modulator = SinOsc.ar(
			freq: mod_freq, 
			//make the sinusoid unipolar between 0 and 1
			mul: 0.5*mod_amp,
			add: 1
		);
		//Carrier Oscilator with Ring Multiplication
		ampmod = SinOsc.ar(
			freq: carrier_freq, 
			mul: EnvGen.kr(Env.adsr, gate, doneAction:2) * amp * carrier_amp * modulator
		);

		Out.ar(out, ampmod);
	}).load(s);
)


// Let's play someting
//-------------------
~ma = Synth("ampmod");  //play the default
~ma.set(\mod_freq, 2);  //increase the tremollo
~ma.set(\mod_freq, 50); //spectrum
~ma.set(\gate, 0);      //stop and free

// Use the mouse to control mod_freq
//---------------------------------
(
SynthDef("Mousex", { arg out=0; 
	Out.kr(out, MouseX.kr(0, 250)) 
}).send(s); // this will send it to the server without saving it.
)

(
//Do the Bus Mappings
~busx = Bus.control(s); // Obtain a control bus dinamically
Synth("Mousex", [\out, ~busx.index]);

//play it
~ma = Synth("ampmod");
~ma.map(\mod_freq,~busx.index);
)

//free when done playing with the mouse
~ma.set(\gate, 0);
