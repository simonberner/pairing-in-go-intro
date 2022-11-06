# pairing-in-go-intro

## Excercise Goal
Everyone gets their hand on the keyboard to get familiar with the gitpod environment.

## Mission Goal
We as a pairing group, need to write a function that takes an integer `n` as an argument
and **prints** and **returns**
fizz: if dividable by 3
buzz: if dividable by 5
fizz buzz: if dividable by 3 and 5
n: otherwise

## Example
```
1 -> "1"
2 -> "2"
3 -> "fizz"
4 -> "4"
5 -> "buzz"
...
15 -> "fizz buzz"
```

## PingPong TDD Approach
1. Write a test (describing an aspect of the program)
2. Run the test (which should fail because the program lacks that feature)
3. Hand-over to the next person
3. Write just enough code (the simplest possible) to make the test pass
4. Refactor the code and run the test
5. Hand-over to the next person and repeat until it conforms to our mission goal