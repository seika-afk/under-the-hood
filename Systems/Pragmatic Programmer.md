Started : 06 November 2026
Ended : 
# Chapter 0

Pragmatic philosophy is the idea that you should judge ideas, decisions, and actions by their practical consequences and whether they actually work in reality

One of the cornerstones of the pragmatic philosophy is the idea of taking responsibility for yourself and your actions in terms of your career advancement

```
Provide Options, Don’t Make Lame Excuses
```

```
When you find yourself saying, “I don’t know,” be sure to follow it up with “—but I’ll find out.” It’s a great way to admit what you don’t know, but then take responsibility
```

==🟢hopelessness can be contagious==

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

That same principle applies to projects, particularly when you’re building something that hasn’t been built before. We use the term tracer bullet development to visually illustrate the need for immediate feedback under actual conditions with a moving goal.

Like the gunners, you’re trying to hit a target in the dark. Because your users have never seen a system like this before, their requirements may be vague. Because you may be using algorithms, techniques, languages, or libraries you aren’t familiar with, you face a large number of unknowns. And because projects take time to complete, you can pretty much guarantee the environment you’re working in will change before you’re done.


```
Look for the important requirements, the ones that define the system. Look for the areas where you have doubts, and where you see the biggest risks. Then prioritize your development so that these are the first areas you code.
```


```
Tracer development is an approach where you develop software by first making the execution flow visible and understandable 

like having stated function : input and output
files imported etc
```

Tracer code is code written specifically to observe or record the execution of a program.

tracer code example:

define everything without logic yet


(for logging)
```
def calculate_total(price, quantity):
    print("[TRACE] Entering calculate_total")
    print(f"[TRACE] price={price}, quantity={quantity}")

    result = price * quantity

    print(f"[TRACE] Returning {result}")
    return result


def apply_discount(total):
    print("[TRACE] Entering apply_discount")

    result = total * 0.9

    print(f"[TRACE] Returning {result}")
    return result


def checkout(price, quantity):
    print("[TRACE] Entering checkout")

    total = calculate_total(price, quantity)
    final = apply_discount(total)

    print("[TRACE] Checkout completed")
    return final


print(checkout(100, 2))
```

```
[TRACE] Entering checkout
[TRACE] Entering calculate_total
[TRACE] price=100, quantity=2
[TRACE] Returning 200
[TRACE] Entering apply_discount
[TRACE] Returning 180.0
[TRACE] Checkout completed
180.0
```

![[Pasted image 20261114094732.png]]


# Prototypes


Like the car makers, we can target a prototype to test one or more specific aspects of a project.

### Estimating

![[Pasted image 20261115100323.png]]

o, if after doing all the necessary work, you decide that a project will take 125 working days (25 weeks), you might want to deliver an estimate of “about six months.”


Before you get too committed to model building, cast around for someone who’s been in a similar situation in the past. See how their problem got solved. It’s unlikely you’ll ever find an exact match, but you’d be surprised how many times you can successfully draw on others’ experiences.


```
“How long will it take to paint the house?”

“Well, if everything goes right, and this paint has the coverage they claim, it might be as few as 10 hours. But that’s unlikely: I’d guess a more realistic figure is closer to 18 hours. And, of course, if the weather turns bad, that could push it out to 30 or more.”

That’s how people estimate in the real world. Not with a
 single number (unless you force them to give you one) but with a range of scenarios

```


![[Pasted image 20261115101150.png]]

---

# Chapter 3 


As Pragmatic Programmers, our base material isn’t wood or iron, it’s knowledge. We gather requirements as knowledge
We gather requirements as knowledge, and then express that knowledge in our designs, implementations, tests, and documents

#### Debugging

- check failing testcases
- read the damn error
- logging tracing
- process of elimination
```
- Is the problem being reported a direct result of the underlying bug, or merely a symptom?
    
- Is the bug really in the framework you’re using? Is it in the OS? Or is it in your code?
    
- If you explained this problem in detail to a coworker, what would you say?
    
- If the suspect code passes its unit tests, are the tests complete enough? What happens if you run the tests with this data?
    
- Do the conditions that caused this bug exist anywhere else in the system? Are there other bugs still in the larval stage, just waiting to hatch?
```

---
### VIDEO TUTORIAL ON DEBUGGING
debugging  is getting away from delusion of what you think your code does vs what it actually does

- you cant change everything
	- so you gotta prioritize ,what things can fail on priority
```
readability
formating
```

-- dont think of how feature can  be useful, but rather how it can go wrong

`you should not be thinking about the syntax but the problem you are trying to solve`

typing is easy - debugging is hard


opensource could help , finding the issue in the current codebase type shit
-> making a algo type for myself 

```
investigative technique

- narrow down the bug 
	  - at what point after that error begins(see ins and outs of all folders)
- use a visual debugger
-or
- print functions
- validation functions(validates data)
- tests
- lanugage doesnt know your constraints
- make code deterministic
```

