# pairing-in-go-intro

## Mission Goal
We as a pairing group, need to write a function that returns n'th element of Fibonacci sequence.

## Example
```
Fibanocci(0) -> 0
Fibanocci(1) -> 1
Fibanocci(2) -> 1
Fibanocci(3) -> 2
Fibanocci(4) -> 3
Fibanocci(5) -> 5
Fibanocci(6) -> 8
Fibanocci(7) -> 13
Fibanocci(8) -> 21
```

## PingPong TDD Approach
1. Write a test (describing an aspect of the program)
2. Run the test (which should fail because the program lacks that feature)
3. Write just enough code (the simplest possible) to make the test pass
4. Refactor the code and run the test
5. Hand-over to the next person and repeat