# golang-basics
Repository contains the journey of learning Go Programming language

## Variables
Variables have two visibility scopes (NO private scope) and they're set by the variable name. When the name starts with an uppercase it is exposed in the package and when lowercase it is local.

### Global Variables
Global variables need a full declaration explicitly defining the variable type.

### Constants
Constants can be grouped in a paranthesis. iota is an incrementing counter in a specified constant block. One of the practical implementation is user roles and permissions using bitmasking.

## Arrays and Slices
Arrays are fixed collections that are passed by value. Slices on the other hand are pointers that point to the underlying array declaration. Using methods __cap()__ and __len()__ to check capacity and length.

### How does slice capacity increment?
Making some tests show how a 0-size struct will increment capacity by just 1 element, and int or string will duplicate on each growth, and a 3-byte struct "roughly" doubles on each growth. Check in [playground](https://go.dev/play/p/OKtCFskbp2t).