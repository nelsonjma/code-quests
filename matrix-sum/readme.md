<div id="top"></div>


## matrix Sum

<b>url:</b> https://www.reddit.com/r/dailyprogrammer/comments/oirb5v/20210712_challenge_398_difficult_matrix_sum/

Find the minimum such sum when selecting 20 elements (one from each row and column) of this 20x20 matrix. The answer is a 10-digit number whose digits sum to 35.

There's no strict runtime requirement, but you must actually run your program all the way through to completion and get the right answer in order to qualify as a solution: a program that will eventually give you the answer is not sufficient.

123456789   752880530   826085747  576968456   721429729 <br>
173957326   1031077599  407299684  67656429    96549194 <br>
1048156299  663035648   604085049  1017819398  325233271 <br>
942914780   664359365   770319362  52838563    720059384 <br>
472459921   662187582   163882767  987977812   394465693 <br>

If you select 5 elements from this matrix such that no two elements come from the same row or column, what is the smallest possible sum? The answer in this case is 1099762961 (123456789 + 96549194 + 663035648 + 52838563 + 163882767).

for 5x5 python code can run, for the challenge 20x20 this does not work, we require faster code.

permutations:

5x5 => 120
20x20 => 2,432,902,008,176,640,000