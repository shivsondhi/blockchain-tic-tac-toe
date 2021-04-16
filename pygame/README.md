# Tic-Tac-Toe using PyGame and Python
## Tic-Tac-Toe game implementation using the pygame library in python.

This is my first independent game implementation using the pygame library and python. The poor UI/UX of this game is plainly visible, but I was more concerned with the porper working of my code. Once I'm a little more comfortable with the module, I will start focusing on the UI/UX aspect too, which without question is quite important too!

Since I seem to prefer informative README files, I will try to elaborate on what the file consists of a little bit.

### Constants
First off, I declare some global variable outside of any scope. These are my basic measurements and color values. Declaring these parameters in a similar way to mine, makes it a lot more intuitive when you are going through your own code anytime later. Of course you can just plug in the numbers wherever necessary, but that reduces readability greatly and must be avoioded. This is called syntactic sugar. 

### Main Function and Game loop
Next is the main function where most of the large processes take place. Here, I have created all of the objects I will be needing throughout the program, like the font objects and the surface objects. I even create a few banners and buttons. Most of the main logic of the game goes into the game loop which runs indefinitely, till some conditions have been fulfilled. 

### Supporting Functions
Outside the main function, we have all of the supporting functions that are used throughout the program. Each of them has a short description for itself and they're all quite intuitive once you're clear with the logic.

To end, I am not sure this is the most efficient method for this particular game. However, I feel like it is a good starting point and I will keep updating the game locally as and when I work on it. A huge shoutout must go to Al Sweigart for his book, [Making Games with Python and Pygame](https://www.amazon.com/Making-Games-Python-Pygame-Sweigart/dp/1469901730?ie=UTF8&tag=playwithpyth-20&linkCode=as2&camp=1789&creative=9325&creativeASIN=0982106017) from where I am learning about the pygame module. He has also been kind enough to provide some of his content as free .pdf files online, so you can go look through that if you'd like to!
