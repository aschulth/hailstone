# hailstone

The tool calculates the hailstone sequence for a given positive number `n`. The
hailstone sequence is computed as follows: if `n` is even, the next number is
`n/2` else it is `3n+1`. If the sequence reaches `1` the program stops as it
would otherwise enter an infinte loop `1 4 2 1`. It has been conjectured by
Lothar Collatz and others that every positive integer eventually falls through
to `1` (cf. [Collatz conjecture](https://en.wikipedia.org/wiki/Collatz_conjecture)).

DISCLAIMERS:
  - The script is not intended for research but as an exercise in Go.
  - Beware of large numbers! They will cause infinite loops.

## Usage
```
Usage: hailstone [-l|--length] <positive_number>

Calculate the hailstone sequence for a given positive number.

Options:
  --help        print this help page.
  -l|--length   print only sequence length.

Examples:
  Calculate hailstone sequence for 7.
  $ hailstone 7
  // Output: 7 22 11 34 17 52 26 13 40 20 10 5 16 8 4 2 1

  Calculate the length of the hailstone sequence for 7.
  $ hailstone -l 7
  // Output: 17

Error codes:
  2 Divergent sequence. Congratulations to your Fields medal!
```
