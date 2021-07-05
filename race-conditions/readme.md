# Race conditions example

A goroutine starts appending strings to an array. At the same time, the main goroutine is
attempting to print out all the items in the array. So, both goroutines are accessing
the same resource at the same time, which is a race condition.