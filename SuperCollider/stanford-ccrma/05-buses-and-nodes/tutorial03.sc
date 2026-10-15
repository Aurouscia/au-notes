/*
	***********************************************************************************
	Supercollider Basic Tutorial (PART 3)

	Busses and Nodes.
	based partly on Cottle's book and Polishook Tutorial
	Read chapter 18 in Cottle books for a more detailed explnation on how to connect
	buses. You won't find there how to change the number of inputs and outputs and
	how jack works with supercollider

	Juan-Pablo Caceres, February, 2007
	Under Subversion Control System 
	***********************************************************************************
*/

s = Server.default.boot;


/*
This tutorial deals with Buses, Nodes and Groups. Buses are usefull con conect different
Synth from the client side of the language. Nodes and groups deal with the fact that order
matters in sound synthetisis when you conect different UGens and  Synths.
*/

/*
BUSES:
----------------------------------------------------------------------------------------
There are two kinds of busses in SC: audio and control.
By default, SuperCollider has 128 buses for audio signals and 4096 for control signals. 
The buses, which are items in an array, are what SuperCollider uses to represent audio and control rate data.

You inter-connect audio busses using Out.ar and In.ar.

Busses are common in mixers and are used to route several signals to or from a common
destination. They don't change the signal in any way, they just route it.

As explained earlier, Supercollider has by default:
128 audio buses
4096 control buses

You can confirm that (and change the defaults) with the following lines:
*/

o = Server.default.options;
o.numAudioBusChannels;
o.numControlBusChannels;

/*
Supercollider has by default in linux 8 inputs and 8 outputs (coming from jack, take a look at your jack connect).
 You can check this with the following:
*/

o.numInputBusChannels;
o.numOutputBusChannels;

//Using the same comands, you can chage those default, for example, if you need 20 outputs, and 10 inputs
//you do:

o.numInputBusChannels = 10;
o.numOutputBusChannels = 20;
//and you need to restart the server fot the changes to take effect:
s.reboot; //(and check you jack connect)

//Let's go back to the defaults;
s.quit;
o.numInputBusChannels = 8;
o.numOutputBusChannels = 8;
s.boot;


/*
Out is going to wirte to a bus, and In is going to read from a bus.
By default, the first 8 channels (ouputs) are numbered from 0 to 7, and the inputs from 8 to 15.
So, for example 0 sends to the left channel, and 1 to the right channel:
*/
{Out.ar(0, 0.1*SinOsc.ar)}.play;
{Out.ar(1, 0.1*SinOsc.ar)}.play;

/*
In the same way, In is going to read from a bus.
use a mic or any input conection in jack to test it:
*/

{In.ar([8,9])}.play;

/*
Note that by default the ouptut is beeing send to channels 0 and 1, so
the above line is the same as:
*/

{Out.ar([0,1], In.ar([8,9]))}.play;

/*
Check also the default comodity class AudioIn, which handles input (in Jack
in the case of linux) starting with input one.
For example, in the case of 8 channels inputs and outputs,
In.ar(8) is equivalent to AudioIn.ar(1)
*/


/*
The control busses have no default connections.

Let's look at the following patch to understand how buses work.
First, we look at a complete patch in the "normal way":
*/

//Entire patch
{SinOsc.ar(LFNoise0.kr([10, 15]) * 400 + 800, 0, 0.3)}.play;
//Just the SinOsc
{SinOsc.ar(800, 0, 0.3)}.play;
//Just the LFNoise0
{LFNoise0.ar([10, 15])}.play;


//Now, let do the same with buses:
{Out.kr(20, LFNoise0.kr([8, 11], 500, 1000))}.play;
{Out.ar(0, SinOsc.ar(In.kr(20, 2), 0, 0.3))}.play

//Note that you can start first the SinOsc and then the LFNoise0 (try that).


/*
Now let's see an example in which you can change dinamicale the control buses
that control the frequency in SinOsc
*/

(
//Start all the synths
SynthDef("LFN0Control",
{Out.kr(20, LFNoise0.kr([8, 11], 500, 1000))}).play(s);
SynthDef("SineControl",
{Out.kr(22, SinOsc.kr([3, 3.124], mul: 500, add: 1000))}).play(s);
SynthDef("MouseControl",
{Out.kr(24, MouseX.kr([100, 200], 1000))}).play(s);
a = SynthDef("PatchableSine", {arg busInNum = 0;
Out.ar(0, SinOsc.ar(In.kr(busInNum, 2), 0, 0.3))}).play(s);
)

a.set(\busInNum, 20); //set to LFNoise0
a.set(\busInNum, 22); //set to SineControl
a.set(\busInNum, 24); //set to MouseControl


/*
NODES:
----------------------------------------------------------------------------------------
We didn't care about the order of the buses because they where control buses.
When you work with audio signals, order is important and changes the final sonic result,
and the direction that the audio signal flows matters.

This is very important especially when using non-linear processing, but it's going to affect
all the connection in SC.

By default, the localhost and internal servers each boot 
with two predefined groups: the RootNode and the default_group 
(see their help files). To see this, start the localhost server 
and then evaluate
*/

s.queryAllNodes;

/*
Group(0) is the rootnode group and Group(1) is the default_group. 
Group(1) is indented to show that it branches from Group(0).
*/

//Let's look at the following example:

(
//The definitions can be any order
SynthDef("Saw", { arg filterBus = 16;
Out.ar(filterBus, LFSaw.ar([60, 90]))
}).send(s);
SynthDef("Filter", {arg filterBus = 16;
Out.ar(0, RLPF.ar(In.ar(filterBus, 2), LFNoise0.kr(12, 300, 350), 0.2))
}).send(s);
)
//won't work
Synth("Saw"); //source
Synth("Filter"); //filter
//works
Synth("Filter"); //filter
Synth("Saw"); //source
//or
Synth("Filter"); Synth("Saw");

s.queryAllNodes;

/*
The second case works because the Saw is at the HEAD, and the filter is at the TAIL.
This of signals flowing from HEAD to TAIL, in the following fashion:

	HEAD(saw)
	 |
	 |
	 |
	 v
	TAIL(tail)

So you can explicitly assing nodes to head or tail:
*/

Synth.tail(s,"Filter");
Synth.head(s,"Saw");

//now the order in which you excecute doesn't matter:
Synth.head(s,"Saw");
Synth.tail(s,"Filter");


//DYNAMIC BUS ALLOCATION
//-----------------------------------------------------------------------------
/*
So, we where discussing about the buses to connect Synth. You really want to 
be aware (and use) node numbers when you send audio or read audio from Jack.
In any other case, you don't want to be chosing "by hand" node numbers to connect
different modules.

SuperCollider has the class Node, that you can use to dinamically allocate note.
It's going to chose for you a number that is not used by any input or output, so
if later you increase the number of channels in Jack, you patch will continue 
to work without changing anynthing.

Read the Bus class help file for a lot of details.

All the following discussion applies to audio and control buses.
*/

//To allocate a bus (audio):
b = Bus.audio;
//get the bus number
b.index;
//free the bus
b.free;

//The following is going to allocate 2 channels
c = Bus.audio(s,2);
b = Bus.audio(s,2);
//You won't get sound here because the buses are not going to the output
{Out.ar(c.index, 0.1*Saw.ar)}.play;
{Out.ar(b.index, 0.1*SinOsc.ar([500, 1000]))}.play;

//Now you can read from the bus (BUT remember that the input HAS to be in the HEAD,
//so excecute the next lines in the following order (now you will hear):
{In.ar(c.index)}.play
{Out.ar(c.index, 0.1*Saw.ar)}.play;

s.queryAllNodes;


/*
to be continue with Groups and more on Nodes
*/

