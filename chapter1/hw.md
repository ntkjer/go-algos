##1.3.1 
implemented in stacks/fixedcap but this doesnt really make sense for the following reasons:
    * we dont use fixed array type as internal collection because go does not support using a var int and expects constant. This means we cant variably grow the size of the array and makes 100% sense in the context of go. Slices were made for this reason but its a bummer cause slices arent comparable!
    * since we set N as the len(entries) it will always be at capacity
## 1.3.2
it
## 1.3.3
b. The last push is on 9 with two items on the stack (1 -> 0). The next to pop should be 1 before 0 but returned vals have it as 0 before 1.
## 1.3.4

