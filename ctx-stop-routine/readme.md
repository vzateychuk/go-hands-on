# Context stops a goroutine

Here we start a Goroutine with an infinite loop counting from zero until  we decide 
to stop it. We will make use of the context to notify the routine to stop and a 
sleeping function to make sure we know how many iterations we do