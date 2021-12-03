# hailstone
```
Usage:
  hailstone --help
  hailstone [-l|--length] <positive_number>

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
  2 Number too large.
  3 Divergent sequence! // not implemented
```
