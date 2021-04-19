## Tic-Tac-Toe game implementation using the pygame library in python.

This game was created in the style of Al Sweigart in his book, [Making Games with Python and Pygame](http://inventwithpython.com/makinggames.pdf). It helped me get comfortable with concepts in the pygame module, working with UI/UX, and with the basics of python.

## The game file
### Global variables
First, the global variables are declared with global scope - these include constants and color values. The global variables act as syntactic sugar, and are plugged in at different places in the code instead of plugging in numbers directly.

### Main method and Game loop
The main function is where the main game processes take place. It contains all the objects needed in the program, like font and surface objects; and even buttons and banners. This is followed by the game loop which contains the main game logic, and runs indefinitely till some conditions have been fulfilled. 

### Supporting methods
Outside the main function, we have all the supporting methods that are used throughout the program. Each one has a short description of its function.
