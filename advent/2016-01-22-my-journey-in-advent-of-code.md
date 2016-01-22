---
title: My Journey in advent of code
author: hannibal
layout: post
date: 2016-01-22
url: /2016/01/22/my-journey-in-advent-of-code
categories:
  - Go
  - Programming
  - Performance
---

Hello folks.

I wanted to share with you my tale of working through the problems with [Advent Of Code](http://adventofcode.com).

It is a nice tale and there are a few things I learned from it, especially in Go, since I used that solve all of the problems. So, let's get started.

Solving the problems
====================

The most important lesson I learned while doing these exercises was, how to solve these problems. A couple of them were simple enough to not have to over think it, but most of them got very tricky. I could have gone with a brute force attempt, but as we see later, that wasn't always a very good solution. And people who used that, actually just got lucky finding their solutions.

The eight stages of breaking down a problem according to this book [Thinking Like a Programmer](http://www.amazon.co.uk/Think-Like-Programmer-Introduction-Creative/dp/1593274246/ref=sr_1_1?s=books&ie=UTF8&qid=1453449575&sr=1-1&keywords=thinking+like+a+programmer) are the following:

* Have a plan
* Rephrase
* Divide
* Start with what you know
* Reduce
* Analogies
* Experiment
* Don't get frustrated

Have a plan and understanding your goal
------------------------------------

This is simple. Always have a plan of what you would like to do, and how to start. This will help you massively along the way to not to loose sight of what your goal is actually. For example, look at [Day 24](http://adventofcode.com/day/24). At first, it looks like a permutational puzzle, but if you understand the solution we are looking for, you realize that there is an easier way of finding it. Since you only want the packages which consists of the fewest item counts, you would only care about the largest numbers because those will be the fewest which still give you the desired package weight. Suddenly the problem gets easier because you don't have to worry about the other groups any longer.

Rephrase
--------

Rephrasing the problem with your own words can help in understanding it better. Or even better, try explaining it to somebody else. If you cannot rephrase it, you didn't understand it in the first place.

Divide
------

If the problem seems daunting because it's massive, just divide it into smaller chunks. This is something that we usually do with large problems, but it's more subtle than that. If you face a problem which seems complex, just factor out parts of it until you got a problem which you do understand. Even if you have to butcher the original puzzle problem. It doesn't matter. Adding complexity later is easier than adding complexity in its infancy.

Start with what you know && Finding analogies
---------------------------------------------

This one speaks for itself. If you know parts of the problem, because you know analogy for it, or you faced something similar before, or exactly that, start with that.

Reduce
------

If the problem seems too complex, remove complexity. Start with a smaller set. Preferably something testable (I'll come back to that later). Remove constraints, or add them as desired. A constraint makes it harder to solve the puzzle? Remove it, and try solving it without. After that, the solution will give you insight into the problem and you can add that constraint back in.

Consider [Day 11](http://adventofcode.com/day/11). I had fun with this one. In order to easy it up a little, I first, removed the constraint of doing the increment with letters. I did it with numbers. I also removed the constraint of doing it within the confines of a limited length array. After I got that I'll use modulo to make the numbers wrap around, it was way more easy to apply it to characters. And after a little fidgeting this came to life:

~~~go
passwd[i] -= 'a'
passwd[i] = (passwd[i] + 1) % (('z' - 'a') + 1)
passwd[i] += 'a'
~~~

The -,+ 'a' is needed so that it's dealing with ascii code from 0 - 'z'. This basically makes it so that when I reach the end of the alphabet it will wrap around and start from 'a' again.

Experiment
----------

This led to more solutions than I care to admit. Basically just start experimenting with solutions which are in your head. There is a chance, that what you come up with, will be the solution. This goes very well with the principle of *Make it work*, *Make it right*, *Make it fast*. Just have something working first, and than you can make it work properly after. It's always better to have *something* rather than nothing.

And last but not least...

Don't get frustrated
--------------------

This is something I cannot say strongly enough. Seriously. **DO NOT GET FRUSTRATED**. Most of the problems were designed to be harder. Unless you work as a programmer professionally for several years now, or this is a field of interest for you, you will spend a day hacking around on a problem and trying to find a solution which is adequate. In these times, you will get frustrated and think you are too stupid for this, this couldn't be more far from the truth! You might need some inspiration, you might need some time away from the screen, it helps if you draw out the problem in a piece of paper, or just think about it without seeing it for a while. Just take a break, eat something, watch a comedy and get back to it later with a fresh view.

Technical Gotchas
=================

So after the general problem solving side of things, I learned many things about Go, and about the tidbits of this language.

Byte Slices
-----------

I already knew that []byte is more performant and that Go optimizes on them more, but not to this extent. As in my previous blog posts I discovered that using them can really make a huge difference. Go even has a library called ```bytes``` which has helper functions similar to that of ```strings``` to help you out in these situations. Go optimizes on map recalls as well when you cast to string from []byte and use that as a map key like this: myMap[string(data)].

Brute Force or Looping
----------------------

Most of the times you could get away with looping or trying to brute force out a solution. But there were times, where you really had to huddle down and think the problem through. Because simply looping, either took too long, or didn't come up with a good answer. That's why I rather always start with: 'How could I solve this without looping?'. This will get you into the right mindset. Or thinking: 'How could I solve this without taking each and every combination into account?'. These questions will help you to think about the problem without loops. Or only if you REALLY must use one.

Doing this will get you into the right way of thinking. I know that in advent of code there is a Leaderboard and you could get on it if you were fast. But most of the times having a fast solution is far from having the right solution.

Structs are Awesome
-------------------

I like using structs. They are a very lightweight way of defining objects, structures which stick together. For example in the [Day 6](http://adventofcode.com/day/6) Light puzzle, or even [Day 3](http://adventofcode.com/day/3) Traveling santa example, a struct which stuck x,y locations together and made it a map key, it was trivial to make my gif out of it with SVG ->

![Traveling Santa](https://github.com/Skarlso/goprojects/blob/master/advent/day3/day1.gif)

Go is Simple to Read
--------------------

[opinion] I like Go because of its simplicity. You don't see stuff in Go most of the times, where you need to look three times to understand what the heck is going on. I like filter, reduce, map and syntactic sugar, but they make for a very poor reading experience. Go, in that way, choose not to incorporate these paradigms and I find that refreshing. [/opinion]

Testing
-------

TDD is something we all should know by now and care about. When I'm doing puzzles, or finger exercises, I tend to not write tests. But on a more complex puzzle, or a task, I always start with a test. Especially if you are given samples for a puzzle which work. That's a gold mine. You can tweak your algorithm using those samples until they work and then simply apply a larger sample size.

Tests will also help you with breaking down a problem and identifying parts which you already know.

For example [Day 13](http://adventofcode.com/day/13). Optimal Seating arrangements. Or the similar [Day 9](http://adventofcode.com/day/9). Which was calculating shortest route distance. Or the password one, Day 11 which I showed before. In these cases, tests helped me make the core of the algorithm solid. Calculating connections, or the odd regex here and there, which was making sure that the password was validated properly.

Tests will also help you to be able to move on after you found your solution. When I was done with the first iteration of passwords which was still using strings, I went on to optimize it, to use []byte. The tests helped me to know that the code was still working as expected after the refactoring.


Closing words
=============

All in all it was a massive amount of fun doing these exercises and I'm thankful to the creator for making it. And I did enjoy the story behind the exercises as well. I think this site stood out because it had a fun factor. For simple exercises there are a lot of other sites -like Project Euler, or Sphere Judge Online-, which just plainly present you a problem and that's it. It's still fun, but it can also became boring very fast. Don't forget the fun factor which makes you plow on and go into a blind frenzy that you cannot quit until it's done. That's the fun part.

Thank you for reading!
Have a nice day.
Gergely.
