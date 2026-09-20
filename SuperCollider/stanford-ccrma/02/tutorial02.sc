/*
	***********************************************************************************
	Supercollider Basic Tutorial (PART 2)
	Based mainly on the help files at Streams/
	Streams-Patterns-Events1.html
	Streams-Patterns-Events2.html
	Streams-Patterns-Events3.html
	Streams-Patterns-Events4.html
	Streams-Patterns-Events5.html
	Streams-Patterns-Events6.html
	Streams-Patterns-Events7.html

	Juan-Pablo Caceres, January 17, 2007
	Under Subversion Control System
	***********************************************************************************
*/
s = Server.default.boot;

/*
This tutorial is an introduction to Patterns and how to use them in order to make music.
We will continue with synthesis in the next classes.
I strongly recommend that you go through the help files, they explain everything in detail.

Take also a look at the help files of Array
*/

/*
Streams
=========================================================================
A stream, how the name describe, is a secuence of values.
In SuperCollider, any object that responds to the messages:
- next
- reset
is a stream.

The next value in the sequence is obtained by sending the message next to the stream object.
The sequence can be restarted from the beginning by sending the message reset to the stream
object. A stream can be of finite or infinite length. When a finite length stream has
reached the end, it returns nil.

By default every object is a stream is supercollider, it just returns itself.
This is not usful for our purposes, so I'm going to skip the details.

FuncStream creates a function that responds to the next and reset message.
We create here a function that selects a random value from an array.
The stream in this case is an infinite random sequence.
*/

a = FuncStream({ [1, 2, 3, 4].choose });
a.next;
a.next;
a.next;

/*
Another useful class is Routine
Routines are functions that can return in the middle and then resume where
they left off when called again.
To better understand the difference, let's look at the following example

In the first case (FuncStream) the function always returns 1 with next.
This is because the function is being called over and over again.
*/

a = FuncStream({1});
a.next;
a.next;

/*
With Routine, the yield message returns a value from the Routine. The next time theRoutine
is called it begins by returning from the yield and continues from that point.
Take a look at the Stream help file for more details
*/

(
x = Routine({
	1.yield;
	2.yield;
	3.yield;
});
)

x.next;
x.next;
x.next;
x.next;

/* Other stuff to look on streams (on the help files)
- Math operations on Streams
- Filtering operations on streams (collect, select, reject messages)
*/


/*
Patterns
=========================================================================
The best way to understand a pattern is like a specification to create something (in our case, streams)
This way you can create multiple streams with a single specification, like using a cookie cutter.

A Pattern is any object that responds to the asStream message by creating a Stream.

Pfunc returns a FuncStream. You may not see why Pfunc is usefull now (why not just use FuncStream instead?).
You will see later that Pfunc can be used with other pattern specifications BEFORE creating the stream.
Remember that Pfucn doesn't return a stream, it returns a Pattern.
*/

Pattern
a = Pfunc({ [1, 2, 3, 4].choose });
b = a.asStream; // make a stream from the pattern
5.do({ b.next.postln; }); // print 5 values from the stream


/*
Prout is similar but returns a Routine
*/
(
a = Prout({
3.do({ arg i; 3.rand.yield; })
});
)
// make two streams from the pattern
b = a.asStream;
c = a.asStream;
4.do({ b.next.postln; }); // print 4 values from first stream
4.do({ c.next.postln; }); // print 4 values from second stream




/*
ListPatterns
=========================================================================
ListPatterns are Patterns that iterate over arrays of objects in some fashion.
All ListPatterns have in common the instance variables list and repeats.
The list variable is some Array to be iterated over. The repeats variable is
some measure of the number of times to do something, whose meaning
varies from subclass to subclass. The default value for repeats is 1.

Pseq is a Pattern that cycles over a list of values. The repeats variable gives
the number of times to repeat the entire list.
*/

a = Pseq([1, 2, 3], 2); // repeat twice
b = a.asStream;
7.do({ b.next.postln; });


/*
Take also a look at
- Pser
- Prand
- Pxrand
- Pshuf
*/

/*
Nesting Patterns
=========================================================================
If a Pattern encounters another Pattern in its list, it embeds that pattern in its output.
That is, it creates a stream on that pattern and iterates that pattern until it ends before moving on.
For example here is one pattern nested in another.
*/


(
a = Pseq([
	Pseq([1, 2], 2),
	Pseq([3, 4], 2),
	Pseq([5, 6], 2)
], 3);
)
b = a.asStream;
17.do({ b.next.postln; });



/*
Value Patterns, Event Patterns and Pbind
=========================================================================
The patterns discussed in parts 2 and 3 are known as "value patterns" because their streams
return a single value for each call to next. Here we introduce "event patterns" which once turned
into streams, return an Event for each call to next.

The class Pbind provides a bridge between value patterns and event patterns. It binds symbols in
each event to values obtained from a pattern. Pbind takes arguments in pairs, the first of a pair
being a Symbol and the second being a value Pattern. Any object can act as a Pattern,
so you can use constants as the pattern ( see \amp in the example below ).

The Pbind stream returns nil whenever the first one of its streams ends.
*/


a = Pbind( \freq, Pseq([440,880], inf) );
b = a.asStream;
b.next(Event.default);
b.reset;

/*
The nice thing about Pbind is that you can play the pattern.
There are a lot of stuff that happens under the scenes that you don't need to worry
when you play the pattern. It will also call the default values for the symbols that
are needed in you synth and not specified (see below how to see all the env variables)

(When calling Pattern-play an EventStreamPlayer is automatically generated which handles
scheduling as well as passing the protoEvent into the event stream.)
*/

a.play;

//You can play for example with a different duration
(
Pbind( \freq, Pseq([440,880], inf),
\dur, 0.1).play;
)

//To see the list of all the default Environment variables
Event.default.parent(_.postcs);

/*
We will see a complete example now on how to use all this stuff with
Steve Reich's Piano Phase
*/


/*
To be continued...
*/
