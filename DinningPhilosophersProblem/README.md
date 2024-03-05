# Dinning Philosopher's Problem Using sync.Mutex and Channels

Here is my submission for the Week 4 assignment in the "Concurrency in Go" specialization course on Coursera.https://www.coursera.org/learn/golang-concurrency/home/welcome

## Problem Statement provided by Coursera

There should be 5 philosophers sharing chopsticks, with one chopstick between each adjacent pair of philosophers.
Each philosopher should eat only 3 times (not in an infinite loop as we did in lecture)
The philosophers pick up the chopsticks in any order, not lowest-numbered first (which we did in lecture).
In order to eat, a philosopher must get permission from a host which executes in its own goroutine.
The host allows no more than 2 philosophers to eat concurrently.
Each philosopher is numbered, 1 through 5.
When a philosopher starts eating (after it has obtained necessary locks) it prints “starting to eat ” on a line by itself, where is the number of the philosopher.
When a philosopher finishes eating (before it has released its locks) it prints “finishing eating ” on a line by itself, where is the number of the philosopher.

## My Solution
I am using sync.mutex which make sures the multiple goroutines cannot excecute blocks of code concurrently.Mutexes provide a way to protect critical sections of code, ensuring that only one goroutine can access the shared resource at a time. In our scenario forks/chopsticks must be a mutex, as they are shared with philosphers and thus must help manage them effectively.

Action of eating would be goroutine which needs shared chopsticks.

There might be Dijkstra's solution of picking least numbered chopstick. It was a brilliant idea, but we were asked not to use that, rather to have a host which allows 2 philosophers to have food at a time.

I request host with channel of type bool, with buffer size of 2 so we can only have 2 request at a time.

I also added done channel which is of type struct so i wont be using memory for just signalling when all philosophers have finished eating three times.

Please refer the code and provide pointers.

## Thank you