##1.3.1 
implemented in stacks/fixedcap but this doesnt really make sense for the following reasons:
    * we dont use fixed array type as internal collection because go does not support using a var int and expects constant. This means we cant variably grow the size of the array and makes 100% sense in the context of go. Slices were made for this reason but its a bummer cause slices arent comparable!
    * since we set N as the len(entries) it will always be at capacity
## 1.3.2
it
## 1.3.3
b. The last push is on 9 with two items on the stack (1 -> 0). The next to pop should be 1 before 0 but returned vals have it as 0 before 1.
## 1.3.4
implemented as isbalanced
## 1.3.5
binary respresentation when N = 50
it checks for powers of two and pushes on stack to calculate binary representation of base 10 numbers
## 1.3.6
* Dequeueing every item in q and pushing it on a stack preservese the order of the queue.
* enqueueing every single item from the stack as its popped means we append popped items that were in order to the back of the queue. This will *reverse* the original q input order. 
## 1.3.7
implement peek in the stack interface
```
stacks
├── dynamic
│   ├── main.go
│   └── tobe.txt
└── fixedcap
    ├── main.go
        └── tobe.txt
```
## 1.3.8
it was the - the best - of times - - - it was - the - -
>>it

