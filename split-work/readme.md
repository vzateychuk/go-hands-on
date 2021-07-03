Equally Splitting the Work between Routines
---

In this exercise, we will see how we can perform our sum of numbers in a predefined
number of routines for them to gather the result at the end. Essentially, we want to
create a function that adds numbers and receives the numbers from a channel. When
no more numbers are received by the function, we will send the sum to the main
function through the channel