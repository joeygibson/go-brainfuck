[
    This program prints "Hello World!" and a newline to the screen, its
    length is 106 active command characters. [It is not the shortest.]

    This loop is an "initial comment loop", a simple way of adding a comment
    to a BF program such that you don't have to worry about any command
    characters. Any ".", ",", "+", "-", "<" and ">" characters are simply
    ignored, the "[" and "]" characters just have to be balanced. This
    loop and the commands it contains are ignored because the current cell
    defaults to a value of 0; the 0 value causes this loop to be skipped.
]

++++++++        // set c0 to 8
[
    >++++       // Add 4 to c1; this will always set c1 to 4
    [           //  as the cell will be cleared by the loop
        >++     // Add 2 to c2
        >+++    // Add 3 to c3
        >+++    // Add 3 to c4
        >+      // Add 1 to c5
        <<<<-   // Decrement loop counter in c1 
    ]           // Loop till c1 is 0; number of iterations is 4
    >+          // Add 1 to c2
    >+          // Add 1 to c3
    >-          // Subtract 1 from c4
    >>+         // Add 1 to c6
    [<]         // Move back to first cell with 0; this will be c1
                //  which was cleared by the previous loop
    <-          // Decrement loop counter in c0
]               // Loop until c0 is 0; number of iterations is 8

// This is the result of the previous loop
// Cell #:      c0   c1  c2  c3  c4  c5  c6
// Contents:    0    0   72  104 88  32  8
// Pointer:     ^

>>.             // Cell #2 has value 72 which is 'H'
>---.           // Subtract 3 from Cell #3 to get 101 which is 'e'
+++++++..+++.   // Likewise for 'llo' from Cell #3
>>.             // Cell #5 is 32 for the space
<-.             // Subtract 1 from Cell #4 for 87 to give a 'W'
<.              // Cell #3 was set to 'o' from the end of 'Hello'
+++.------.
--------.       // Cell #3 for 'rl' and 'd'
>>+.            // Add 1 to Cell #5 gives us an exclamation point
>++.            // And finally a newline from Cell #6         

