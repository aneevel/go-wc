# Summary

Implementation of the WC Unix Command Line tool in Golang.

For more details, see the wonderfully laid out project task description on [Coding Challenges](https://codingchallenges.fyi/challenges/challenge-wc).

Includes a test file.

# Features 

- Byte Count Flag - Default option, simply provide -c to get the byte count of a file
- Line Count Flag - provide -l to get the number of lines in a file
- Word Count Flag - provide -w to get the number of words in a file
- Character Count Flag - provide -m to get the number of chars in a file
- No Flags - equivalent to -l -w -c

# Usage 

To build, simply run;

`go build`

To run, invoke the binary with a filename (or provide input via stdin) and then provide the flags you want;

`./ccwc test.txt -l`

With stdin piped in;

`cat test.txt | ./ccwc`

# Lessons learned 

I got a great amount of experience actually reading Golang package documentation and using a number of fairly low level packages.
If I were to do this project again, I would change the interface for the individual functions to instead take in a byte slice; this 
represents what I want them to do more effectively, and would unify the way I handle a file vs. stdin. It should be the responsibility
of the main (calling) function to provide a byte slice for each function, while the individual functions get to be a bit more generic 
as a result.
