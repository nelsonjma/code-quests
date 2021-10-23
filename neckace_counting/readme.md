<div id="top"></div>


## Necklace counting

<b>url:</b> https://www.reddit.com/r/dailyprogrammer/comments/g1xrun/20200415_challenge_384_intermediate_necklace/

Given k and n, find the number of distinct k-ary necklaces of length n. That is, the size of the largest set of k-ary necklaces of length n such that no two of them are equal to each other. You do not need to actually generate the necklaces, just count them.

For example, there are 24 distinct 3-ary necklaces of length 4, so necklaces(3, 4) is 24. Here they are:

AAAA  BBBB  CCCC <br>
AAAB  BBBC  CCCA <br>
AAAC  BBBA  CCCB <br>
AABB  BBCC  CCAA <br>
ABAB  BCBC  CACA <br>
AABC  BBCA  CCAB <br>
AACB  BBAC  CCBA <br>
ABAC  BCBA  CACB <br>

A more efficient way is with the formula:

```necklaces(k, n) = 1/n * Sum of (phi(a) k^b) ```

for all positive integers a,b such that a*b = n.

For example, the ways to factor 10 into two positive integers are 1x10, 2x5, 5x2, and 10x1, so:

```
necklaces(3, 10)
  = 1/10 (phi(1) 3^10 + phi(2) 3^5 + phi(5) 3^2 + phi(10) 3^1)
  = 1/10 (1 * 59049 + 1 * 243 + 4 * 9 + 4 * 3)
  = 5934
```