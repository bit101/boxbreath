# boxbreath

`boxbreath` is a command line app to guide you through one of two types of breathwork sessions - box breathing and 4-7-8 breathing.

## Box Breathing

What is "box breathing" you ask? It's a type of breath work exercise, you might call it mindfulness. It's just a short, easy structured breathing exercise. You breath in for a certain count (usually 4 or 5), hold it for that count, breath out for the same count and hold it again for that count. Repeat that for some number of repetitions or target time.

Box breathing can supposedly relieve stress, help you focus, etc. Some even say it can reduce blood pressure and have other physical health benefits. I won't make any wild claims, but it seems to help me a bit.

## 4-7-8 Breathing

4-7-8 breathing is a slightly different type of breathwork, also very simple though. In this one, you breath in for four seconds, hold it for seven seconds and release it over eight seconds. And repeat.

## Background

I looked at various mobile apps that do this kind of thing, and unsurprisingly most want to charge you something like a $10 monthly subscription fee. And have all kinds of ridiculous social networking and gamification features. Not at all what I wanted in a stress relief app. Honestly you don't even need this app. Just do what I just described above. It's easy.

But I do like the idea of having something there to keep track of the beats and reps as I tend to lose focus... while trying to improve focus. So I made a command line app! It's free! It's open source.

## Installation

The easiest way if you have Go installed it just typing `go install github.com/bit101/boxbreath@latest`. You're done.

Alternately, check out this repo, `go build` it and put the executable whereever it's convenient for you.

## Use

Simple use, just type:

`boxbreath`

or

`path/to/boxbreath`

if you build it yourself.

When you launch the app, it will let you choose which type of session you want to do (box or 4-7-8) and let you specify how many breaths to do (and for box breathing, the count of each segment).

It will then give you an estimate of how long the session will take. Hit enter to begin.

If you know which session you want, you can type `boxbreath -box` or `boxbreath -478` to jump right to that type of session.
