# Chapter 0

Pragmatic philosophy is the idea that you should judge ideas, decisions, and actions by their practical consequences and whether they actually work in reality

One of the cornerstones of the pragmatic philosophy is the idea of taking responsibility for yourself and your actions in terms of your career advancement

```
Provide Options, Don’t Make Lame Excuses
```

```
When you find yourself saying, “I don’t know,” be sure to follow it up with “—but I’ll find out.” It’s a great way to admit what you don’t know, but then take responsibility
```

==Psychologists have done studies  that show hopelessness can be contagious==

```
Don’t Live with Broken Windows
```

Don’t leave “broken windows’’ (bad designs, wrong decisions, or poor code) unrepaired

"no broken windows"

Everyone will guard their own resources. Sometimes this is called “start-up fatigue.’

It’s time to bring out the stones. Work out what you can reasonably ask for. Develop it well. Once you’ve got it, show people, and let them marvel


```
People find it easier to join an ongoing success. Show them a glimpse of the future and you’ll get them to rally around
```


```
Remember the big picture
```

# Chapter 1

## Good - Enough Software
you can discipline yourself to write software that’s good enough—good enough for your users, for future maintainers, for your own peace of mind. You’ll find that you are more productive and your users are happier


```
Good enough:
All systems must meet their users’ requirements to be successful, and meet basic performance, privacy, and security standards. We are simply advocating that users be given an opportunity to participate in the process of deciding when what you’ve produced is good enough for their needs
```


```
Make Quality a Requirements Issue
```


##  Your Knowledge Portfolio


```
Your ability to learn new things is your most important strategic asset. But how do you learn how to learn, and how do you know what to learn?
```


##### Invest Regularly
The habit is as important as the sums, so plan to use a consistent time and place, away from interruptions. Learn regularly something new


##### Diversify
The more different things you know, the more valuable you are. As a baseline, you need to know the ins and outs of the particular technology you are working with currently. But don’t stop there. The face of computing changes rapidly

##### Manage risk
Technology exists along a spectrum from risky, potentially high-reward to low-risk, low-reward standards.

##### Buy low, sell high
Learning an emerging technology before it becomes popular can be just as hard as finding an undervalued stock, but the payoff can be just as rewarding.


Critically Analyze What You Read and Hear

## Communicate


"It’s not just what you’ve got, but also how you package it"

---
# Chapter 2

## Pragmatic Approach

### Essence of Good Design

why ?
> |Good Design Is Easier to Change Than Bad Design|

(design principle)

ETC Principle : Easier to change
also called single responsibility principle, isolating different systems so they are easier to change

Do it when you save a file. Do it when you write a test. Do it when you fix a bug.
There’s an implicit premise in ETC. It assumes that a person can tell which of many paths will be easier to change in the future
```
ultimate “easy to change” path: try to make what you write replaceable. That way, whatever happens in the future, this chunk of code won’t be a roadblock.
```

## DRY principle : Don’t Repeat Yourself


Every piece of knowledge must have a single, unambiguous, authoritative representation within a system.
The alternative is to have the same thing expressed in two or more places. If you change one, you have to remember to change the others, or, like the alien computers, your program will be brought to its knees by a contradiction.

DRY is about the duplication of knowledge, of intent. It’s about expressing the same thing in two different places, possibly in two totally different ways.

Not All Code Duplication Is Knowledge Duplication
Duplication in Documentation
DRY Violations in Data


overall ` MAKE IT EASY TO REUSE OR CHANGE`



# 10 Orthogonality

```
Two or more things are orthogonal if changes in one do not affect any of the others. In a well-designed system, the database code will be orthogonal to the user interface: you can change the interface without affecting the database, and swap databases without changing the interface.
```

Eliminate Effects Between Unrelated Things


"modular"
"component based"
"layered"
![[Pasted image 20261110111640.png]]




```
Decorator Pattern: adding functionality to things without changing them.
```


Keep your code decoupled

Write shy code—modules that don’t reveal anything unnecessary to other modules and that don’t rely on other modules’ implementations.If you need to change an object’s state, get the object to do it for you. This way your code remains isolated from the other code’s implementation and increases the chances that you’ll remain orthogonal.

Avoid global data

Every time your code references global data, it ties itself into the other components that share that data. Even globals that you intend only to read can lead to trouble (for example, if you suddenly need to change your code to be multithreaded). In general, your code is easier to understand and maintain if you explicitly pass any required context into your modules. In object-oriented applications, context is often passed as parameters to

Avoid similar functions

Often you’ll come across a set of functions that all look similar—maybe they share common code at the start and end, but each has a different central algorithm. Duplicate code is a symptom of structural problems. Have a look at the Strategy pattern in Design Patterns for a better implementation.


#### Testing it :

Pretty easier , test each components separately


## 11 Reversibility


```
But you said we’d use database XYZ! We are 85% done coding the project, we can’t change now!” the programmer protested. “Sorry, but our company decided to standardize on database PDQ instead—for all projects. It’s out of my hands. We’ll just have to recode. All of you will be working weekends until further notice
```

Changes don’t have to be that Draconian, or even that immediate. But as time goes by, and your project progresses, you may find yourself stuck in an untenable position

### Flexible Architecture

While many people try to keep their code flexible, you also need to think about maintaining flexibility in the areas of architecture, deployment, and vendor integration
What you can do is make it easy to change. Hide third-party APIs behind your own abstraction layers

No one knows what the future may hold, especially not us! So enable your code to rock-n-roll: to “rock on’’ when it can, to roll with the punches when it must.


## 12 . Tracer Bullets

























