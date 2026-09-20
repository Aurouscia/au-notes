/*
	***********************************************************************************
	Piano Phase by Steve Reich (using routines)

	A complete reference on the theretical aspects of the phase
	construction can be found at:
	Heinrich K. Taube
	"Notes from the Metalevel: Introduction to Algorithmic Music Composition"
	Routledge; 1 edition (November 24, 2004),
	Pages 130-136.

	The score can be found in :
	Stefan Kostka & Dorothy Payne,
	"Tonal Harmony",
	McGraw-Hill Humanities/Social Sciences/Languages; 5 edition (June 10, 2003),
	Page 542.

	Created by Juan-Pablo Caceres, January 4, 2007
	Under Subversion Control System
	***********************************************************************************
*/

s = Server.default.boot;

//Very simple Piano SythDefs for be use.
//Piano1 is on the left channel
//Piano2 is on the right channel
//----------------------------------------------------------
(
SynthDef("pianophase1",
	{ arg out=0, freq=440, sustain=0.05, gate;
		var env;
		env = EnvGen.kr(Env.perc(0.01, sustain, 0.2), doneAction:2);
		Out.ar(out, SinOsc.ar(freq, 0, env))
	}).load(s);

SynthDef("pianophase2",
	{ arg out=1, freq=440, sustain=0.05, gate=1;
		var env;
		env = EnvGen.kr(Env.perc(0.01, sustain, 0.2), doneAction:2);
		Out.ar(out, SinOsc.ar(freq, 0, env))
	}).load(s);
)


//Define a funtion to convert tempo from BPM to sec.
//--------------------------------------------------
(
~bpm2sec = {arg tbpm;
	var tsec;
	tsec = 60/tbpm;
};
)

//Parameters and Pattern calculations
//----------------------------------
(
~tempoBPM = 72; //global tempo (from score)

//tempo in seconds for each pitch (1/8 notes)
~tempo = ~bpm2sec.value(~tempoBPM)/6;

//Pitches for one measure (pattern)
~notes1 = Pseq(#[64, 66, 71, 73, 74, 66, 64, 73, 71, 66, 74, 73].midicps, inf).asStream;
~notes2 = Pseq(#[64, 66, 71, 73, 74, 66, 64, 73, 71, 66, 74, 73].midicps, inf).asStream;

//Now lets make the calcultations for piano 2
//Piano 1 will always play at a stady state, and Piano 2 will phase
//We decide how long the phasing is going to be by deciding:
//p = number of pulses the phasing takes
//n = number of events to play in that time
//In this example we will asume that the phasing occurrs in one measure
p = 12; n = 13;
~phasetime = p/n;
~nphase = 5; //<-- this is the number of measure to play in phase

//We now create the patter of tempo for Piano 2
~p2tempoPatter = Pseq([
	Pseq(Array.fill(12,1), ~nphase),//<-- number of measure to play in phase
	Pseq(Array.fill(n,~phasetime))]//<-- this is the phase array
	,inf).asStream;
)

//Play the Piece
(
Routine({
		loop({
			Synth( "pianophase1", [\out, 0,	\freq,  ~notes1.next]);
			~tempo.wait;
		})
	}).play;

Routine({
		loop({
			Synth( "pianophase2", [\out, 1,	\freq,  ~notes2.next]);
			(~tempo*~p2tempoPatter).next.wait;
		})
	}).play;
)

//Reset all Streams before starting again
(
~notes1.reset;
~notes2.reset;
~p2tempoPatter.reset;
)



