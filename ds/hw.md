

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
## 1.3.9
infix 
## 1.3.10
eval postfix written
## 1.3.11 
Done
## 1.3.12

## 1.3.13
b c and d are impossible since a dequeue would mean that we remove the first item queued. therefore `n < N `Item where n is the smaller of numbers from 0-9. Since items are inserted in ascending order, we cannot break the ascsending queue. 

## 1.3.14
wont work with constant var size for arrays as the size is inherently part of the type in go.

## 1.3.15
## 1.3.16
## 1.3.17

# LinkedList Exercises

## 1.3.18
x.next = x.next.next 
repoints the next item to the next of the x.next
deletes the item from the list

## 1.3.19
traverse the links until the last node
reassign the next of x.next to be nil
```
x := s.first
if !s.isEmpty() {
    if s.size() == 1 { 
        first = nil
        break
        } else {
            for x != nil {
                if x.next.next == nil {
                    x.next.next = nil
                }
            }
        }
        s.size--
}
```
implemented in linkedlist/main.go
## 1.3.20
in linkedlist/main.go
## 1.3.21 
in linkedlist/main.go
## 1.3.22
inserts node t immediately after node x
## 1.3.23
we removed the link from node x to its next node
this results in t referring to itself as its next node 
## 1.3.24
in linkedlist/main.go
## 1.3.25 
in linkedlist/main.go
## 1.3.26
in linkedlist/main.go
## 1.3.27 
in linkedlist/main.go
## 1.3.28 
in linkedlist/main.go
## 1.3.29
in circular-linkedlist/main.go
## 1.3.30
