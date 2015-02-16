##  What's you talkin 'bout?

Communication is difficult in concurrent applications because of state, and the fact that the majority of developers learn to
develop not worrying about sharing state.

### When writing a single path of execution, who cares if a method changes the value of that variable?

However if that variable is now shared between two separate areas of execution, how do you know it will be in the correct
state that the code expects it to be?
