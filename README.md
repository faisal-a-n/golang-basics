# golang-basics
Repository contains the journey of learning Go Programming language

## Variables
Variables have two visibility scopes (NO private scope) and they're set by the variable name. When the name starts with an uppercase it is exposed in the package and when lowercase it is local.

### Global Variables
Global variables need a full declaration explicitly defining the variable type.

### Constants
Constants can be grouped in paranthesis. __iota__ is an incrementing counter in a specified constant block. One of the practical implementation is user roles and permissions using bitmasking.

## Arrays and Slices
Arrays are fixed collections that are passed by value. Slices on the other hand are pointers that point to the underlying array declaration. Using methods __cap()__ and __len()__ to check capacity and length.

### How does slice capacity increment?
Making some tests show how a 0-size struct will increment capacity by just 1 element, and int or string will duplicate on each growth, and a 3-byte struct "roughly" doubles on each growth. Check in [playground](https://go.dev/play/p/OKtCFskbp2t).

### The append() function
Keep in mind that a slice is a structure of 3 fields.
- a pointer to the underlying array
- length of the slice
- capacity of the slice
append() function may either modify its argument in-place or return a copy of its argument with an additional entry, depending on the size and capacity of its input. append() function creates a new slice, if the length the slice is greater than the length of the array pointed by the slice.

## Maps and Structs
Maps are pass by reference just like slice. Structs are pass by value even though they retain the reference to map or slice inside them. Struct attributes can have tags enclosed with backticks which are checked by reflect libray or any other validation library

### Struct methods
Struct having methods should be defined on a global scope

### Inheritance in go = Embedding?
Go doesn't support inheritance but has it's composition of embedding. A parent struct is directly embedded onto the child struct without any named parameter

## Control Flow (Defer, Panic & Recovery)
### Defer
When used on a statement it is called after the function completely executes but just before returning the control. It works in LIFO way. __defer__ keyword is used before the statement. Usually used to close resources. The statement state doesn't change at the time of execution.

### Panic
Panic is just a thrown exception. Panic is executed after defer statements.

### Recover
Recover is used to continue program execution and is more like catch method. If recover is called outside the deferred function it will not stop a panicking sequence.

## Interfaces
Interfaces are implicitly defined and can be casted to the type implementing them. Empty interfaces can be used as a plain object to hold any data with type switch.

## Go Routines
Go routines are abstract threads which bind themselves to the OS threads to get the work done by taking rounds. Routines have a race condition. Add __-race__ flag while running the program to check memory manipulations.

## Channels
Need more references...