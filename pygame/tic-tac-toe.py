'''
Tic Tac Toe!
Made by me, not sure how this should work more efficiently, but this is where I'll implement everything I learn. 

Notes:
For functions, the argument and parameter order is tilerow, tilecol; but while traversing or accessing the board data structure, order is tilecol and then tilerow.
Basic Flow:
_____________________________________________________________________________________________________________________________________________________________________
1. Create bottom panel as: 

	P1: Shape	   				  P2: Shape
------------------------------------------------
	New Game						Reset

2. Select P1 shape and P2 shape --> Question asked in top panel.

3. Create empty tic tac toe board.

4. Whoever's turn it is, their P1/P2 gets highlighted in bottom panel.

5. Start Game!

6. Game Won Animation!
_____________________________________________________________________________________________________________________________________________________________________

Changes:
	Reset Move doesn't reset the Player details.
	Pressing Reset Move multiple times will keep changing the shape to be printed to screen!
	Changing window size changes layout completely.
	Use keyboard!
		To select X or O choice at game start.
		To select where to put X or O on the board -> Highlight the selected cell / cell over which cursor is hovering.
		Spacebar to put shape (or X or O)
		R to Reset and N to start new game. 
	Between Player info and Menu put instructions for keyboard or just use the mouse! / Crosses and Knots.
	IMPROVE COLORS!!!!! 
'''

import pygame, sys, random
from pygame.locals import *

#Basic constants
FPS = 15
WINDOWWIDTH = 350
WINDOWHEIGHT = 566
BOARDWIDTH = 3
BOARDHEIGHT = 3
GAPSIZE = 1
FONTSIZE = 25																						#Non compromisable
SHAPESIZE = 50

#Relative constants
TOPPANELWIDTH = WINDOWWIDTH - 20
TOPPANELHEIGHT = WINDOWWIDTH - 20
TILESIZE = TOPPANELWIDTH / 3 
PLAYERHEIGHT = (WINDOWHEIGHT - WINDOWWIDTH) / 2 
MENUHEIGHT = (WINDOWHEIGHT - WINDOWWIDTH) / 2

XMARGIN = int((WINDOWWIDTH - (TILESIZE * BOARDWIDTH) - (BOARDWIDTH - 1) * GAPSIZE) / 2)
YMARGIN = int((WINDOWWIDTH - (TILESIZE * BOARDHEIGHT) - (BOARDHEIGHT - 1) * GAPSIZE) / 2)
BOTTOMPANELYMARGIN = int((PLAYERHEIGHT - FONTSIZE) / 2)												#X margin
BOTTOMPANELXMARGIN = 22																				#Y margin

#Colors			   R    G 	 B 
BLACK 		 =	 ( 0 ,	0,	 0	)																	#X color
WHITE 		 =   (255, 255, 255 )																	#O color
GRAY	 	 = 	 (153,  0,   0  )
TURQGREEN	 =	 ( 0 , 204, 102 )
FADEDORANGE	 =	 (255, 178, 102 )
DARKORANGE	 = 	 (204, 102,  0 	)
LEMONYELLOW	 =	 (255, 255, 153 )
BRIGHTBLUE 	 = 	 ( 0 , 50 ,	255	)

#Assign Colors
BGCOLOR 		= TURQGREEN
LIGHTBGCOLOR 	= LEMONYELLOW
DARKBGCOLOR 	= FADEDORANGE
FONTCOLOR 		= GRAY
BANNERCOLOR		= DARKORANGE
XOCOLOR 		= BLACK
TOPPANELBGCOLOR = WHITE
HIGHLIGHT 		= BRIGHTBLUE


def main():
	global FPSCLOCK, DISPLAYSURF, GAMESURF, BASICFONT, SHAPEFONT, NEW_TEXT, NEW_BUTTON, RESET_TEXT, RESET_BUTTON, P, P1_TEXT, P1_POS, P2_TEXT, P2_POS
	FPSCLOCK = pygame.time.Clock()																	#Create clock object
	P = [None] * 2																					#Two-item list that acts as a queue to store current player details
	lastMove = None																					#Records the last committed move

	pygame.init()
	#Create main display surface and fill it with color
	DISPLAYSURF = pygame.display.set_mode((WINDOWWIDTH, WINDOWHEIGHT))
	DISPLAYSURF.fill(BGCOLOR)
	pygame.display.set_caption("Tic Tac Toe!")
	#Set the Game Screen boundaries
	pygame.draw.rect(DISPLAYSURF, HIGHLIGHT, (10, 10, TOPPANELWIDTH, TOPPANELHEIGHT), 4)
	#Create Game Surface and fill it with color
	GAMESURF = pygame.Surface((TOPPANELWIDTH+1, TOPPANELHEIGHT+1))									#Surface for gameplay
	GAMESURF.fill(TOPPANELBGCOLOR)																	#Will be used later but must be initialized in main()
	
	#Create font objects
	BASICFONT = pygame.font.Font('freesansbold.ttf', FONTSIZE)										#Font for menus
	SHAPEFONT = pygame.font.Font('freesansbold.ttf', SHAPESIZE)										#Font for Shapes

	#Create start-screen buttons and text.
	#Player Information
	P1_TEXT, P1_POS = makePlainText("P1: ", FONTCOLOR, BGCOLOR, BOTTOMPANELXMARGIN, TOPPANELHEIGHT+BOTTOMPANELYMARGIN)
	P2_TEXT, P2_POS = makePlainText("P2: ", FONTCOLOR, BGCOLOR, WINDOWWIDTH-120, TOPPANELHEIGHT+BOTTOMPANELYMARGIN)
	#Game Title
	BANNER_TEXT, BANNER_POS = makePlainText("TIC TAC TOE!", BANNERCOLOR, BGCOLOR, 90, 100)
	#Shape Selection
	START_TEXT, START_POS = makePlainText("P1 Select: ", BANNERCOLOR, BGCOLOR, 70, 200)
	SELECTX_TEXT, SELECTX_BUTTON = makeClickableText("  X  ", BANNERCOLOR, BGCOLOR, 220, 200)
	SELECTO_TEXT, SELECTO_BUTTON = makeClickableText("  O  ", BANNERCOLOR, BGCOLOR, 280, 200)
	#Game Options
	NEW_TEXT, NEW_BUTTON = makeClickableText("New Game", FONTCOLOR, BGCOLOR, BOTTOMPANELXMARGIN, TOPPANELHEIGHT+PLAYERHEIGHT+BOTTOMPANELYMARGIN)
	RESET_TEXT, RESET_BUTTON = makeClickableText("Reset Move", FONTCOLOR, BGCOLOR, WINDOWWIDTH-160, TOPPANELHEIGHT+PLAYERHEIGHT+BOTTOMPANELYMARGIN)

	#Add all buttons and text to the main display surface
	DISPLAYSURF.blit(BANNER_TEXT, BANNER_POS)
	DISPLAYSURF.blit(START_TEXT, START_POS)
	DISPLAYSURF.blit(SELECTX_TEXT, SELECTX_BUTTON)
	DISPLAYSURF.blit(SELECTO_TEXT, SELECTO_BUTTON)
	DISPLAYSURF.blit(NEW_TEXT, NEW_BUTTON)
	DISPLAYSURF.blit(RESET_TEXT, RESET_BUTTON)
	DISPLAYSURF.blit(P1_TEXT, P1_POS) 
	DISPLAYSURF.blit(P2_TEXT, P2_POS)
	
	#start of main game loop
	while True:
		for event in pygame.event.get():
			if event.type == QUIT or (event.type==KEYUP and event.key==K_ESCAPE):					#To shut game window
				pygame.quit()
				sys.exit()
			elif event.type == MOUSEBUTTONUP and (P[0]==None or P[1]==None):						#Activated only in the beginning to select P1 and P2
				if SELECTX_BUTTON.collidepoint(event.pos):											#event.pos contains the coordinates of the mouse click
					P[0] = P1 = 'X'
					P[1] = 'O'
					updatePlayerInfo(P)
					mainBoard = getStartingBoard(P1)
				elif SELECTO_BUTTON.collidepoint(event.pos):										#collidepoint() checks whether event.pos collides with the button
					P[0] = P1 = 'O'
					P[1] = 'X'
					updatePlayerInfo(P)
					mainBoard = getStartingBoard(P1)
			elif event.type == MOUSEBUTTONUP:														#Activated at any subsequent mouse click
				if NEW_BUTTON.collidepoint(event.pos):												#If New Game button pressed
					main()
				elif RESET_BUTTON.collidepoint(event.pos):											#If Reset button pressed
					if lastMove == None:
						break
					else:
						mainBoard = undoMove(lastMove, mainBoard)
				else:																				#For any other mouse click!
					mousex, mousey = event.pos
					if isValid(mousex, mousey, mainBoard):
						mainBoard, lastMove = updateBoard(mousex, mousey, mainBoard, P1)
						if hasEnded(mainBoard):
							pygame.time.wait(1500)
							endAnimation(P1)
		#Update the display surface. All blits don't get printed to the screen until you call the display.update()pygame function!
		pygame.display.update()
		FPSCLOCK.tick(FPS)

#Functions start
def makeClickableText(text, font_color, bg_color, left, top):
	#Creates a text object that can be clicked on (= a button)
	textSurf = BASICFONT.render(text, True, font_color, bg_color)
	textRect = textSurf.get_rect()
	textRect.topleft = (left, top)
	return textSurf, textRect

def makePlainText(text, font_color, bg_color, left, top):
	#Creates a text object that will not be clicked on (= a banner)
	textSurf = BASICFONT.render(text, True, font_color, bg_color)
	pos_tuple = (left, top)
	return textSurf, pos_tuple

def makeShapeText(text, left, top):
	#Creates a text object for the X and O shapes. Font object is different for different size.
	textSurf = SHAPEFONT.render(text, True, XOCOLOR, TOPPANELBGCOLOR)
	pos_tuple = (left, top)
	return textSurf, pos_tuple

def updatePlayerInfo(p):
	#Updates the player information in the menu below the game surface. Note how the changes won't take effect until a call to display.update() is made.
	P1_TEXT, P1_POS = makePlainText("P1: {0}".format(p[0]), FONTCOLOR, BGCOLOR, BOTTOMPANELXMARGIN, TOPPANELHEIGHT+BOTTOMPANELYMARGIN)
	P2_TEXT, P2_POS = makePlainText("P2: {0}".format(p[1]), FONTCOLOR, BGCOLOR, WINDOWWIDTH-120, TOPPANELHEIGHT+BOTTOMPANELYMARGIN)
	DISPLAYSURF.blit(P1_TEXT, P1_POS)
	DISPLAYSURF.blit(P2_TEXT, P2_POS)
	pygame.display.update()

def getStartingBoard(P1):
	#Creates the empty game board and prints it to the screen.
	mainBoard = getEmptyBoard()
	drawBoard(mainBoard)
	drawHighlight(P1, HIGHLIGHT)
	DISPLAYSURF.blit(GAMESURF, (10, 10))
	return mainBoard

def getEmptyBoard():
	#Creates an empty board data structure. (2D list, 3x3)
	board = []
	for x in range(BOARDWIDTH):
		col = []
		for y in range(BOARDHEIGHT):
			col.append(None)
		board.append(col)
	return board 

def drawBoard(board):
	#Draws the board to screen, tile by tile.
	for tilecol in range(len(board)):
		for tilerow in range(len(board[0])):
			drawTile(tilerow, tilecol)

def drawTile(tilerow, tilecol, msg=''):
	#Draws a single tile, with an optional text message at the center(X or O)
	left, top = getLeftTopOfTile(tilerow, tilecol)
	pygame.draw.rect(GAMESURF, XOCOLOR, (left, top, TILESIZE, TILESIZE), 1)
	#Shift top and left co-ordinates to print the message at the approximate center of the tile
	top += int((TILESIZE - SHAPESIZE) / 2)
	left += 41
	textSurf, textRect = makeShapeText(msg, left, top)
	GAMESURF.blit(textSurf, textRect)
	return 

def getLeftTopOfTile(tilerow, tilecol):
	#Get the top and left pixel co-ordinates of any tile, given its board co-ordinates. 
	#Pixel co-ordinates calculated on the game surface and NOT the display surface!
	left = (TILESIZE * tilecol)
	top = (TILESIZE * tilerow)
	return left, top

def getTile(mousex, mousey, board):
	#Given the pixel co-ordinates of mouse click, find the corresponding board co-ordinates of the tile clicked on.
	for col in range(len(board)):
		for row in range(len(board[0])):
			left, top = getLeftTopOfTile(row, col)
			tileRect = pygame.Rect((left, top, TILESIZE, TILESIZE))
			if tileRect.collidepoint(mousex, mousey):
				return row, col
	return (None, None)																				#Return None if mouse click coincides with no tile

def isValid(mousex, mousey, board):
	#Determine if the move determined by mouse click is valid.
	#2 conditions checked: 1. Mouse click does not coincide with a tile(check getTile function) and 2. the coinciding tile already has a value.
	tilerow, tilecol = getTile(mousex, mousey, board)
	if (tilerow == None) or (board[tilecol][tilerow] != None):
		return False
	else:
		return True

def updateBoard(mousex, mousey, board, P1):
	#Update the board data structure and the board on-screen according to the last move.
	tilerow, tilecol = getTile(mousex, mousey, board)
	board[tilecol][tilerow] = P[0]
	drawTile(tilerow, tilecol, P[0])
	#Clear highlights over both players
	drawHighlight(P[0], BGCOLOR)
	drawHighlight(P[1], BGCOLOR)
	P[0], P[1] = P[1], P[0]																			#Change current player queue
	drawHighlight(P1, HIGHLIGHT)																	#Highlight the current player
	DISPLAYSURF.blit(GAMESURF, (10, 10))
	pygame.display.update()
	lastMove = (tilerow, tilecol)																	#Record the last move to be able to undo it
	return board, lastMove

def hasEnded(board):
	#Checks whether the game has ended by checking the board state.
	#Checks 4 conditions: Horizontal / Vertical / Diagonal / Board Full
	#Check for horizontal strike
	for col in range(len(board)):
		row = 1
		if (board[col][row] != None) and (board[col][row - 1] == board[col][row]) and (board[col][row] == board[col][row + 1]):
			strikethrough(row, col, 'v')
			return True
	#Check for horizontal strike
	for row in range(len(board[0])):
		col = 1
		if (board[col][row] != None) and (board[col - 1][row] == board[col][row]) and (board[col][row] == board[col + 1][row]):
			strikethrough(row, col, 'h')
			return True 
	#Check for diagonal strike
	row = col = 1
	if (board[col][row] == None):
		return False
	elif (board[col - 1][row - 1] == board[col][row]) and (board[col][row] == board[col + 1][row + 1]):
		strikethrough(row, col, 'd1')
		return True 
	elif (board[col - 1][row + 1] == board[col][row]) and (board[col][row] == board[col + 1][row - 1]):
		strikethrough(row, col, 'd2')
		return True 
	#Check for filled board
	elif boardFull(board):
		P[0] = 'Tie'
		return True
	else:
		return False

def boardFull(board):
	#Check whether the board is full or not.
	for col in range(len(board)):
		for row in range(len(board[0])):
			if board[col][row] == None:
				return False
	return True

def strikethrough(tilerow, tilecol, dir):
	#For each strikethrough; find the starting tile and the ending tile of the winning, three-tile sequence.
	#Find the start and end co-ordinates of the line that will strikethrough the three-tile sequence.
	#The method is different for all the four types of strikethroughs; horizontal (--), vertical (|) and the two types of diagonals (/ and \).
	if dir is 'h':
		left_start, top_start = getLeftTopOfTile(tilerow, tilecol - 1)
		left_end, top_end = getLeftTopOfTile(tilerow, tilecol + 1)
		strike_y = top_start + (TILESIZE / 2)
		start_pos = (left_start + 5, strike_y)
		end_pos = (left_end + TILESIZE - 5, strike_y)
		pygame.draw.line(GAMESURF, BLACK, start_pos, end_pos, 4)
		DISPLAYSURF.blit(GAMESURF, (10, 10))
		pygame.display.update()
	elif dir is 'v':
		left_start, top_start = getLeftTopOfTile(tilerow - 1, tilecol)
		left_end, top_end = getLeftTopOfTile(tilerow + 1, tilecol)
		strike_x = left_start + (TILESIZE / 2)
		start_pos = (strike_x, top_start + 5)
		end_pos = (strike_x, top_end + TILESIZE - 5)
		pygame.draw.line(GAMESURF, BLACK, start_pos, end_pos, 4)
		DISPLAYSURF.blit(GAMESURF, (10, 10))
		pygame.display.update()	
	elif dir is 'd1':
		left_start, top_start = getLeftTopOfTile(tilerow - 1, tilecol - 1)
		left_end, top_end = getLeftTopOfTile(tilerow + 1, tilecol + 1)
		start_pos = (left_start + 5, top_start + 5)
		end_pos = (left_end + TILESIZE - 5, top_end + TILESIZE - 5)
		pygame.draw.line(GAMESURF, BLACK, start_pos, end_pos, 4)
		DISPLAYSURF.blit(GAMESURF, (10, 10))
		pygame.display.update()	
	elif dir is 'd2':
		left_start, top_start = getLeftTopOfTile(tilerow + 1, tilecol - 1)
		left_end, top_end = getLeftTopOfTile(tilerow - 1, tilecol + 1)
		start_pos = (left_start + 5, top_start + TILESIZE - 5)
		end_pos = (left_end + TILESIZE - 5, top_end + 5)
		pygame.draw.line(GAMESURF, BLACK, start_pos, end_pos, 4)
		DISPLAYSURF.blit(GAMESURF, (10, 10))
		pygame.display.update()

def endAnimation(P1):
	#Flashy animation to celebrate game end! Also displays the game outcome as a banner.
	color1 = DARKBGCOLOR
	color2 = LIGHTBGCOLOR

	drawHighlight(P[0], BGCOLOR)
	drawHighlight(P[1], BGCOLOR)
	for i in range(11):
		color1, color2 = color2, color1
		GAMESURF.fill(color1)
		if P[0] == "Tie":
			winSurf, winRect = makePlainText("Game has been Tied!", HIGHLIGHT, color1, 45, 90)
		elif P[0] == P1:
			winSurf, winRect = makePlainText("Player 2 Has Won!", HIGHLIGHT, color1, 60, 90)
		else:
			winSurf, winRect = makePlainText("Player 1 Has Won!", HIGHLIGHT, color1, 60, 90)
		GAMESURF.blit(winSurf, winRect)
		DISPLAYSURF.blit(GAMESURF, (10, 10))
		pygame.display.update()
		pygame.time.wait(300)
	main()																							#Resets the game to the start state.

def drawHighlight(P1, color):
	#Draws highlight around a text object.
	if P[0] == P1:
		P1_RECT = P1_TEXT.get_rect()
		pygame.draw.rect(DISPLAYSURF, color, (P1_POS[0]-7, P1_POS[1]-7, P1_RECT.width*1.8, P1_RECT.height+10), 3)
	else:
		P2_RECT = P2_TEXT.get_rect()
		pygame.draw.rect(DISPLAYSURF, color, (P2_POS[0]-7, P2_POS[1]-7, P2_RECT.width*1.8, P2_RECT.height+10), 3)
	pygame.display.update()

def undoMove(lastMove, board):
	#Undo move sand return to all previous conditions.
	tilerow, tilecol = lastMove
	board[tilecol][tilerow] = None
	drawTile(tilerow, tilecol, '   ')																#Covers up the existing shape with a blank space.
	DISPLAYSURF.blit(GAMESURF, (10, 10))
	P[0], P[1] = P[1], P[0]
	pygame.display.update()
	return board

if __name__ == '__main__':
	main()

'''
#Imperfect, buggy code, needs some work.

'''