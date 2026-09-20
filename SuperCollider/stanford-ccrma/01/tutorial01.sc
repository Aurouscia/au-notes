/*
	***********************************************************************************
	Supercollider Basic Tutorial
	Based on Joshua Parmenter's Summer Course:
	http://www2.realizedsound.net:8080/josh/SuperCollider.html

    & Mark Polishook's tutorial
	http://akustik.hfbk.net/sc/SuperColliderHelp/Mark_Polishook_tutorial/
	(also availble in the help file system).

	& Cottle's Book

	Juan-Pablo Caceres, January 10, 2007
	Under Subversion Control System
	***********************************************************************************
*/


/*
Evalutating Code
=========================================================================
The list of the most usefull commands to memorize are:
	C-c C-c Evaluate the present line
	C-c C-x	Evaluate a selected line or region
	C-c C-s	Stop the main loop (i.e., shut off playback)
	C-c C-h	Request for Help

C means in emacs (on linux) the control key. So you need to press Control and then c
twice to evluatuate the present line for example.
You can find more in the SCLang Menu. You can use them from there or just memorize
them and use the shortcuts.

To evaluate a code, first make sure that your server (and Jack!!!)is running
(more on servers later on). For that, evaluate the following line (remember C-c C-c):
*/

Server.default.boot;


//Then you can evaluate the following:

{ SinOsc.ar([400, 401], 0, 0.1) * Saw.ar([11, 33], 1) * EnvGen.kr(Env.sine(10)) }.play

//and stop it if you want with C-c C-s


/*
Basics of the language (coments, code evalutation, and variables)
=========================================================================
Comments
--------
Anything that comes after a // in a line is a COMMENT and will be ignored by the language.
It will be colored in red, as the following text. You can (and have) to use them to help you
and others read and understand your code.

If you want to right more than just one line You can also use
/* YOUR COMMENTS HERE */ as we been doing in this tutorial.
*/

//For example this is a single line of comment


/*
And this is a multiline
bla bla
*/


/*
Evaluation
----------
SuperCollider is an interpreted language, which means that you need to tell
SuperCollider to evaluate any code (with the shortcuts in the begining of this
tutorial or with the SCLang menu).
For instance, if I want to execute the following, I can place the cursor
on that line and C-c C-c:
*/

"hello, world"; //in the interpreter (or post) window, the result will appear

/*(On the historical geek note, this is normally the first code you execute in
any programming language tutoral. I think it was first used as the first
example in the Kernighan C Programming language, and afterwards it bacame
a standard.)*/

3 + 3; //Matematical operations


/*You can also select many lines of code with the cursor, and execute all of
them at once (C-c C-x)*/
3 + 3;
6 * pi;
a = "The var 'a' is assigned this string";


/*
You can also wrap large chunks of code in parenthesis, double click one (any of them),
and the code will be selected (highlited). C-c C-x to execute it.

Notice that only the last object to be evaluated is returned to the interpreter. While
everything in this block is evaluated, only the last object is returned as the value
of the code that is executed.
*/

(
3 + 3;
(6 * pi);
a = "The var 'a' is assigned this string";
)


/*
Variables
---------
In SupperCollider, variable names can be anything but MUST begin with a lower
case letter (no numbers) and they must be contiguous.

Variables are declared (identified to the program) with the syntax var followed by a list of
variables separated by commas and terminated by a semicolon.

For example:
*/

var myVar1, imAvar;

/*
The scope of a variable is the function in which it is declared. A variable cannot be used
outside that function. You can declare global variables using the tilde (~). These do not
require the var syntax. All the letters are also by default global variables, so you don't need
the var syntax.
*/

~myGlobalVariable;
a; //this is also a global variable



/*
SuperCollider's help system
=========================================================================
SuperCollider has a fairly large help system built into the program.
To access the Main help page, simply press C-c C-h (and then enter).

Also, most Classes will have some sort of help file associated with it.
You can highlight (or just step on it) a Class, and hit C-c C-h (and enter).
*/

Object; //All classes are highlighted in green.


/*
You can also look at the class definition itself... you are able to see every
Objects class definition. Though some of these objects refer to UGens or language
primitives that are actually references to the programs underlying c-code, you
can often gather quite a bit of information from looking at the class def.
Highlight the Class, and use C-c ; (find references in the SCLang menu)
*/

Object;

/*
You can also see the complete subtree of a class (all the clases that inherit from it),
using the dumpClassSubtree message.

For example:
*/

Object.dumpClassSubtree; // all the classes
Filter.dumpClassSubtree; // a simplier class

/*
There is also a fancier way to navigate through the classes structure in emacs.
Type:
	M-x (M means meta in Emacs, and in linux is the ALT key)

and then:
	sclang-class-tree
and then type filer to see a simplier class
*/


/*
Understanding Clases and Object with a simple example
=========================================================================
The following is just a simple example on classes and objects. A very good introduction
on the theorical aspects of Classes and Objects, can be found on Chapeter 2 of:
The Art and Science of Smalltalk - Simon Lewis (avialable online)

It's a very good and SHORT introduction (higly recommended).


The Date Class shows us quite a bit about the way we can use it. In the class definition,
first, we see a list of instance variables. These are variables that the Class uses for
every instance of this Object. Next, we see a list of functions that you can ask an Class
to do. If the method begins with an *, this signifies a creation method for creating a new
object. The rest are instance methods that give an instance of a method some functionality.
The '.' is the message operator, it sends a message to an Object
*/

// Create a new instance of Date, and assign it to the variable 'a'. We send the getDate creation
// method to the Date Class
a = Date.getDate;

//now 'a' is our Date object, and we can send messages to it;
a.dayStamp;
a.secStamp;

/*
You can get values of variables.
In the class definition, if a variable is pre-pended with a '>', (variables are in the begining of
the code after 'var') we can also 'get' that variables value.
*/
a.year;

// if a variable is pre-pended with a '>', we can also set that variable to a new value
a.year = 2009;


/*
Networks and client/server
=========================================================================
SuperCollider 3 uses a client/server model to operate across a network.
What this means is that users write client programs that ask a server to
do something, that is, they request service. Such requests can occur locally
on one computer or they can be distributed remotely among two or more computers.
Whether the computers are in the same room or separated across the world makes
no difference as long as they're connected on a network.

Client programs in SuperCollider typically specify synthesis definition
(how a particular sound will be made) and synthesis scheduling
(when a particular sound will be made). In turn, a SuperCollider server
(or servers) synthesizes audio according to client instructions.

To summarize, clients request; servers respond.

For now we will be using the default server, so you just have to boot it like this.
DONT FORGET TO START JACK BEFORE YOU START THE SERVER
*/

Server.default.boot;


/*
Basic Elements (taken from Cottle page 62-63)
=========================================================================
Objects:
	Actually, everything in SC is an object, but for the next few
	chapters we will focus on those that are upper case words, numbers,
	quoted text: Synth, SinOsc, LFNoise, EnvGen, 5, 6.7, "string", etc.

Messages:
	These usually follow objects separated by a dot (.): play,
	ar, max, rand, midicps.

Arguments:
	A list of items separated by commas, enclosed in parentheses,
	following a message:
	(1, 2, 4), ("string", 45, myVar), ({function}, 0.1, [1, 2, 3])

Variables:
	User defined names of memory locations, beginning with lower case,
	but often words strung together with caps in the middle: pitchClass,
	nextCount, freqArray, myVar, etc.

Enclosures:
	Matching parentheses, brackets, and braces, that define functions,
	arguments lists, and arrays: (0, 100, 555), {arg next; next = 100},
	[100, 200, 500]

Functions:
	Anything enclosed in braces: {arg next; next = 100}, {possibly many
	pages of code}, {SinOsc.ar(60.midicps)}

Arrays:
	A list of items separated by commas, enclosed in brackets: [1, 2,
	3], ["C", "D"], [a, b, c]

Expression:
	An expression is punctuated by a semicolon. They delineate the order
	in which lines of code are executed.

Ugens:
	Unit generators are a combination of object, message, and argument
	list that result in some type of output. Ugens are connected
	together to create a patch: LFNoise0.ar(0.5), SinOsc.ar(440, 0,
	0.4), Sequencer.ar(s, r, j)

You will become familiar with all this in the future examples and in your own
coding
*/


/*
Unit Generators (UGen)
=========================================================================
SuperCollider follows the Unit Generator paradigm. A Unit Generator
(UGens from now on) is a a low level building block that does some
specific function.

These primitive building blocks are for example tone generators, filters,
spatializer modules, and so on. These blocks can be connected together, in
a static or dinamical way. That means that you can use all the algorithmic
capabilities of the language to create and destroy connections.

All UGens have a set of inputs and outputs. The number of inputs and outputs
depends on the specific function of the UGen.


Let's use a very simple subtractive synthesis example to explain how to connect
UGens to build more complex synthesizers. In subtractive synthesis, you connect a
tone generator (sine, saw, triangle...) to a filter in order to shape the sound.

Tone  Generator -----> Filter -----> Audio Output

Ugen Class have two important methods that you need to remember .ar and .kr.
	.ar stands for Audio Rate
	.kr stands for Control Rate (one sample generated for each block of 64 audio samples, by default)

*/


/*
THE OSCILATOR
-------------
Let's first run our Tone Generator:
*/
{Saw.ar}.play; //sawtooth oscilator

/*
When you run this line, the oscillator Saw is called with its default parameters.
If you go to the help file of Saw, you'll find that the oscillator has 3 parameters.

Saw.ar(freq, mul, add);

The defaults are 440 for kfreq, 1.0 for mul, and 0 for add (you can see these in the source code).
Of course you can specify different values:
*/

{Saw.ar(freq:200, mul:0.5)}.play;


/*
NOTE: As a matter of good style, we recomend that you always specify the argument with its name,
like in the line above. This way the code is much more readable, and you can also use the
order that you prefer (not the default one). This way you can also just specify just the
argument that you need to change.
*/

/*
THE FILTER
----------
Now let's use a simple Low Pass Filter to shape our sound. We look at the help of LPF
LPF.ar(in, freq, mul, add).
Do in the input parameter, we include our Saw.
*/

(
{LPF.ar(in: Saw.ar(freq:200, mul:0.5),
	freq:100)}.play;
)


/*
To make the code more readable, you can assign the saw to a variable.
*/
~mySaw = {Saw.ar(freq:600, mul:0.8)};
(
{LPF.ar(in: ~mySaw,
	freq:500)}.play;
)


/*
SynthDefs (Synthetizer Definitions)
(taken Parmenter & Cottle)
=========================================================================
Use the SynthDef class to build the engines for synths that will run on the server.
The engines, which can be saved to disk and reused, are analogous to presets on
commercial hardware and software synthesizers.

When notated as code in client programs, the engines have two essential parts:
	-a name
	-a function.

In the jargon of SuperCollider, the function is called a ugenGraphFunc.


Template
--------
Here's a template for a synthdef showing that it consists of a name and a ugenGraphFunc
*/

SynthDef(
"aSynthDef", // the 1st argument is the name
{ .... i am a ugenGraphFunc ... } // the 2nd argument is the ugenGraphFunc
)

/*
To make the template functional
	1. put code into the ugenGraphFunc
	2. send a .load message to the synthdef
*/

/*
Let's see how it works with and without SynthDef
*/

//Without Synthdef
(
{
	var rate = 12, att = 0, decay = 5.0, offset = 400;
	var env, out, pan;
	pan = LFNoise1.kr(1/3);
	env = EnvGen.kr(Env.perc(att, decay));
	out = Pan2.ar(
		Blip.ar(LFNoise0.ar(rate, min(100, offset), offset),
			(env)*12 + 1, 0.3),
		pan)*env;
	out
}.play
)

//With SynthDef (naming it) and arguments
(
SynthDef("SH",
	{
		arg rate = 12, att = 0, decay = 5.0, offset = 400;
		var env, out, pan;
		pan = LFNoise1.kr(1/3);
		env = EnvGen.kr(Env.perc(att, decay), doneAction: 2); //free when done
		out = Pan2.ar(
			Blip.ar(LFNoise0.ar(rate, min(100, offset), offset),
				(env)*12 + 1, 0.3),
			pan)*env;
		Out.ar(0, out) //Output bus now explicit
	}).play
)


/*
With SynthDef a few things that were previously done automatically need to be explicit.
	- assigning an output bus using Out.ar. (More on busses later.)
	- stopping the patch when it is done playing.
We use the envelope to generate the event, but CPU and Ugens remained the same.
That is because even though the envelope terminated the sound we hear, the patch
is still running. There are a number of possible results when the envelope terminates.
The default is to do nothing. Changing it to doneAction: 2 signals the server to deallocate
and free up the processing power used for that ugen when it is done.

Now that the code is loaded into the server we can run several copies of the instrument with
control values (for the arguments) using set. The arguments are placed in an array with the
syntax symbol, value: [\rate, 10, \offset, 200]. Run these lines singly but one right after the
other. (SC makes this easy because when you run one line the cursor moves to the next line.
So you can just keep pressing Ctrl-c Ctrl-c.)
*/

//Multiple nodes of SH
//Three separate nodes of "SH" with different arugments.
//Run these three lines in sequence then stop.
a = Synth("SH", [\rate, 10, \offset, 200]);
b = Synth("SH", [\offset, 400, \att, 3.0, \decay, 0]);
c = Synth("SH", [\rate, 30, \offset, 2000]);

//Let them die out or run these lines to stop them.
a.free;
b.free;
c.free;

//Changing a parameter of an existing node. Run these lines in sequence.
a = Synth("SH", [\rate, 23, \offset, 30, \decay, 20]);
a.set(\offset, 1000)
a.set(\offset, 300)
a.set(\offset, 800)
a.free;

//Two nodes with arguments
a = Synth("SH", [\rate, 7, \offset, 200, \decay, 20]);
b = Synth("SH", [\rate, 23, \offset, 1200, \decay, 20]);

a.set(\offset, 40)
b.set(\offset, 1000)
a.set(\offset, 800)
b.set(\offset, 600)
a.set(\offset, 1200)
b.set(\offset, 50)
a.free; b.free;



/*
To be Continued...
*/
