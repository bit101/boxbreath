# boxbreath

`boxbreath` is a command line app to guide you through a box breathing session.

What is "box breathing" you ask? It's a type of breath work exercise, you might call it mindfulness. It's just a short, easy structured breathing exercise. You breath in for a certain count, hold it for that count, breath out for the same count and hold it again for that count. Repeat that for some number of repetitions or target time.

Box breathing can supposedly releive stress, help you focus, etc. Some even say it can reduce blood pressure and have other physical health benefits. I won't make any wild claims, but it seems to help me a bit.

I looked at various mobile apps, and unsurprisingly most want to charge you something like a $10 monthly subscription fee. Ridiculous. Honestly you don't even need this app. Just do what I just described above. It's easy.

But I do like the idea of having something there to keep track of the beats and reps as I tend to lose focus... while trying to improve focus. So I made a command line app! It's free!

## Installation

The easiest way if you have Go installed it just typing `go install github.com/bit101/boxbreath@latest`. You're done.

Alternately, check out this repo, `go build` it and put the executable whereever it's convenient for you.

## Use

Simple use, just type:

`boxbreath`

or

`path/to/boxbreath`

if you build it yourself.

This will start a session immediately, so get ready to breath in. By default it has you breathing and holding to a count of four, and walks you through eight reps.

To change the count, use the command line flag `-count`.

To change how many reps to do, use the flag `-reps`.

For exmaple, to do a session with ten reps and a count of five, type:

`boxbreath -reps 10 -count 5`
