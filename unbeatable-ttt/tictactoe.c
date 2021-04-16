/*
Tic Tac Toe. 

Learnt:
	Clear screen.
	Realloc.
	Random number generator + seeds.
	Sleep.
	Add leading whitespace for scanf of a single character. Otherwise it takes any stray whitespace (' ', \n, etc) as the input character!
	Cross platform support.

TO-DO:
	Make it work properly on Windows.
*/

#include <stdio.h>		// printf and scanf etc.
#include <stdlib.h>		// malloc and realloc and system("clear") etc. 
#include <unistd.h>		// for sleep()
#include <time.h>		// for srand's seed


int BSIZE = 3;			// The board size


void clearscr();
void get_board(char[BSIZE][BSIZE]);
void print_board(char[BSIZE][BSIZE]);
int gameOver(char[BSIZE][BSIZE]);
void getRowCol(int*, int);
int getMove(int, int);
int isLegalMove(char[BSIZE][BSIZE], int);
int getLegalMoves(int*, char[BSIZE][BSIZE]);
int getWinningMoves(int*, char[BSIZE][BSIZE], int, char);
int updateMoves(int*, int, int[6], int);
int inPos(char[BSIZE][BSIZE], char, char);
void findSpl(int[2], char[BSIZE][BSIZE], char, char);


int main(int argc, char *argv[]) {
	if (argc > 1) {
		printf("Error: Too many command-line arguments passed. Accepts none.");
		return 0;
	}

	int endGame = 0;
	while (!endGame) {
		char again;
		int mode;
		char cshape, pshape;
		printf("Use the numbers on your keyboard to select a move according to the move-to-number chart displayed during your turn.\nEnter 1 now to start game in normal mode.\nEnter 2 now to start game in hard mode\nEnter 0 at any time to quit.\n");
		scanf("%d", &mode);
	
		// if game quit then return
		if (mode == 0) {
			printf("Okay quiting game...\nGame Over!\n");
			return 0;
		}

		// if game is started then choose shape 
		printf("Choose shape - 'X' or 'O' -\n");
		scanf(" %c", &pshape);
		if ((pshape == 'X') || (pshape == 'x')) {
			pshape = 'X';
			cshape = 'O';
		}
		else if ((pshape == 'O') || (pshape == 'o')) {
			pshape = 'O';
			cshape = 'X';
		}
		else {
			printf("Unrecognized shape! Default to X. Restart game later to re-choose. Starting game...\n");
			sleep(4);
			pshape = 'X';
			cshape = 'O';
		}


		// continue with game
		char board[BSIZE][BSIZE];
		int result;
		int choice;
		int *moveset;
		int row, col, move;
		int movecounter = 1;
		int check = 1;
		int random;
		int rowcols[2];
		char won;
	
		// seed for random number
		srand( time(NULL) );

		// Get the board data structure
		get_board(board);
		//Print the board to screen
		clearscr();
		printf("Game Board:\n");
		print_board(board);

		// Start game
		char moves[3][3] = {{'1', '2', '3'}, {'4', '5', '6'}, {'7', '8', '9'}};
		while (gameOver(board) == 0) {
			printf("Choose your move:\n");
			print_board(moves);
			scanf("%d", &choice);
			// check if game was quit
			if (choice == 0) {
				printf("Okay quiting game...\nGame Over!\n");
				return 0;
			}
			// check if the move is legal
			if (isLegalMove(board, choice) == 1) {
				col = (choice % 3) - 1;
				row = (int)(choice / 3);
				if (col < 0) {
					col = 2;
					row--;
				}
				board[row][col] = pshape;
				clearscr();
				print_board(board);
			}
			else {
				clearscr();
				printf("Invalid move! Try again.\n");
					print_board(board);
				continue;
			}
			// check if game is over after this step
			if (gameOver(board) != 0) {
				won = 'P';
				break;
			}

			// Make the execution wait, as if the CPU is thinking
			printf("\nWait for CPU's turn...\n");
			sleep(2);
			printf("Yes?");

			// CPU's turn
			moveset = malloc(sizeof(int));
			// get a list of all possible moves
			int len = getLegalMoves(moveset, board);

			// hard mode code here!
			if (mode == 2) {
				if (movecounter == 1) {
					// First move of CPU
					// t = center, r = corner, e = edge
					if (inPos(board, pshape, 't') == 1) {
						// Take any corner and game is drawn.
						check = 0;
						int moves[] = {1, 3, 7, 9};
						len = updateMoves(moveset, len, moves, 4);
					}
					else if (inPos(board, pshape, 'r') == 1) {
						// Take the center
						check = 1;
						int moves[] = {5};
						len = updateMoves(moveset, len, moves, 1);
					}
					else if (inPos(board, pshape, 'e') == 1) {
						// Take a corner beside the player's move
						check = 2;
						int moves[2];
						findSpl(moves, board, pshape, 'e');
						len = updateMoves(moveset, len, moves, 2);
					}
				}
				else if ((movecounter == 2) && check) {
					// Second move of CPU
					if ((check == 1) && (inPos(board, pshape, 'r') == 2)) {
						// Take an edge
						int moves[] = {2, 4, 6, 8};
						len = updateMoves(moveset, len, moves, 4);
					}
					else if ((check == 1) && (inPos(board, pshape, 'e') == 1)) {
						// Take the corner between the two opponent moves
						int moves[1];
						findSpl(moves, board, pshape, 'b');
						len = updateMoves(moveset, len, moves, 1);
					}
					else if ((check == 2) && (inPos(board, pshape, 't') != 1)) {
						// Take the center
						int moves[] = {5};
						len = updateMoves(moveset, len, moves, 1);
					}
				}
			}

			// get a list of all possible winning moves
			len = getWinningMoves(moveset, board, len, cshape);
			if (len > 1) {
				// get a list of all possible blocking movess
				len = getWinningMoves(moveset, board, len, pshape);
			}
			// randomly select one of the moves form the list
			long random = rand();
			move = random % len;
			getRowCol(rowcols, moveset[move]);
			board[rowcols[0]][rowcols[1]] = cshape;
			clearscr();
			print_board(board);
			won = 'C';
			if (movecounter < 3)
				movecounter++;
		}
	
		// Game Over!
		printf("Game Over!\n");

		// Print result
		result = gameOver(board);
		switch (result) {
			case 1: 
				if (won == 'P')
					printf("You have won!");
				else if (won == 'C')
					printf("You have lost!");
				break;
			case 2: 
				printf("Match tied.");
				break;
			case 0: 
				break;
		}

		printf("\nPlay again?\t(y) Yes \t(n) No\n");
		scanf(" %c", &again);
		if ((again == 'n') || (again == 'N'))
			endGame = 1;
		else if ((again != 'y') && (again != 'Y'))
			printf("Invalid response, ending game!");
		else 
			clearscr();
	}
}


void clearscr() {
	#if defined(__linux__) || defined(__unix__) || defined (__APPLE__)
	system("clear");
	#endif

	#if defined(_WIN32) || defined(_WIN64)
	system("cls");
	#endif
}


void get_board(char board[3][3]) {
	for (int row=0; row<3; row++) {
		for (int col=0; col<3; col++) {
			board[row][col] = ' ';
		}
	}
}


void print_board(char board[3][3]) {
	for (int i=0; i<3; i++) {
		printf(" _____ _____ _____\n|");
		for (int j=0; j<3; j++) {
			printf("  %c  |", board[i][j]);
		}
		printf("\n");
	}
	printf(" _____ _____ _____\n");
}


void getRowCol(int *rowcol, int choice) {
	int row, col;
	col = (choice % 3) - 1;
	row = (int)(choice / 3);
	if (col < 0) {
		col = 2;
		row--;
	}
	rowcol[0] = row;
	rowcol[1] = col;
}


int getMove(int row, int col) {
	return ((row*3) + col + 1);
}


int isLegalMove(char board[3][3], int choice) {
	int rowcol[2];
	getRowCol(rowcol, choice);
	if (board[rowcol[0]][rowcol[1]] == ' ')
		return 1;
	else
		return 0;
}


int getLegalMoves(int moveset[], char board[3][3]) {
	int move, size=1;
	for (int row=0; row<3; row++) {
		for (int col=0; col<3; col++) {
			if ((board[row][col] != 'X') && (board[row][col] != 'O')) {
				move = (row * 3) + col + 1;
				moveset[size-1] = move;
				size++;
				moveset = realloc(moveset, size*sizeof(int));
			}
		}
	}
	return (size-1);
}


int inPos(char board[BSIZE][BSIZE], char shape, char region) {
	int rowcols[2];
	int inpos = 0;
	if (region == 't') {
		int movepos = 5;
		getRowCol(rowcols, movepos);
		if (board[rowcols[0]][rowcols[1]] == shape)
			inpos++;
	}
	else if (region == 'r') {
		int movepos[] = {1, 3, 7, 9};
		for (int i=0; i<4; i++) {
			getRowCol(rowcols, movepos[i]);
			if (board[rowcols[0]][rowcols[1]] == shape)
				inpos++;
		}
	}
	else {
		int movepos[] = {2, 4, 6, 8};
		for (int i=0; i<4; i++) {
			getRowCol(rowcols, movepos[i]);
			if (board[rowcols[0]][rowcols[1]] == shape)
				inpos++;
		}
	}
	return inpos;
}


void findSpl(int moves[2], char board[BSIZE][BSIZE], char shape, char region) {
	int rowcole[2];
	int rowcolc[2];
	int loc;
	if (region == 'e') {
		int movepos[] = {2, 4, 6, 8};
		for (int i=0; i<4; i++) {
			getRowCol(rowcole, movepos[i]);
			if (board[rowcole[0]][rowcole[1]] == shape) {
				if (rowcole[0] == 1) {
					loc = getMove(rowcole[0], rowcole[1]);
					moves[0] = loc - 3;
					moves[1] = loc + 3;
					return;
				}
				else {
					loc = getMove(rowcole[0], rowcole[1]);
					moves[0] = loc - 1;
					moves[1] = loc + 1;
					return;
				}
			}
		}
	}
	else {
		int movepose[] = {2, 4, 6, 8};
		int moveposc[] = {1, 3, 7, 9};
		int edgerow, edgecol, cornrow, corncol;
		for (int i=0; i<4; i++) {
			getRowCol(rowcole, movepose[i]);
			getRowCol(rowcolc, moveposc[i]);
			if (board[rowcole[0]][rowcole[1]] == shape) {
				edgerow = rowcole[0];
				edgecol = rowcole[1];
			}
			if (board[rowcolc[0]][rowcolc[1]] == shape) {
				cornrow = rowcolc[0];
				corncol = rowcolc[1];
			}
		}
		if (edgerow == cornrow) {
			loc = getMove(edgerow, edgecol);
			if (edgecol > corncol)
				moves[0] = loc + 1;
			else
				moves[0] = loc - 1;
			return;
		}
		else if (edgecol == corncol) {
			loc = getMove(edgerow, edgecol);
			if (edgerow > cornrow)
				moves[0] = loc + 3;
			else
				moves[0] = loc - 3;
			return;
		}
		else if ((edgerow-cornrow == 2) || (cornrow-edgerow == 2)) {
			loc = getMove(edgerow, corncol);
			moves[0] = loc;
		}
		else if ((edgecol-corncol == 2) || (corncol-edgecol == 2)) {
			loc = getMove(cornrow, edgecol);
			moves[0] = loc;
		}
		return;
	}
}


int getWinningMoves(int moveset[], char board[3][3], int movelen, char p) {
	int winset[6]; 	// max possible size of winset
	int i = 0;
	for (int r=0; r<3; r++) {
		for (int c=0; c<3; c++) {
			if ((r == 1) && (c == 1) && (board[r][c] == p)) {
				if (board[r-1][c-1] == p)
					winset[i++] = getMove(r+1, c+1);
				else if (board[r+1][c+1] == p)
					winset[i++] = getMove(r-1, c-1);
				if (board[r-1][c+1] == p)
					winset[i++] = getMove(r+1, c-1);
				else if (board[r+1][c-1] == p)
					winset[i++] = getMove(r-1, c+1);
			}
			else if ((r == 1) && (c == 1) && (board[r][c] == ' ')) {
				if ((board[r-1][c-1] == board[r+1][c+1]) && (board[r-1][c-1] == p))
					winset[i++] = getMove(r, c);
				if ((board[r-1][c+1] == board[r+1][c-1]) && (board[r-1][c+1] == p))
					winset[i++] = getMove(r, c);
			}
			if ((r == 1) && (board[r][c] == p)) {
				if (board[r-1][c] == p)
					winset[i++] = getMove(r+1, c);
				else if (board[r+1][c] == p)
					winset[i++] = getMove(r-1, c);
			}
			else if ((r == 1) && (board[r][c] == ' ')) {
				if ((board[r-1][c] == board[r+1][c]) && board[r-1][c] == p)
					winset[i++] = getMove(r, c);
			}
			if ((c == 1) && (board[r][c] == p)) {
				if (board[r][c-1] == p)
					winset[i++] = getMove(r, c+1);
				else if (board[r][c+1] == p)
					winset[i++] = getMove(r, c-1);
			}
			else if ((c == 1) && (board[r][c] == ' ')) {
				if ((board[r][c-1] == board[r][c+1]) && board[r][c-1] == p)
				winset[i++] = getMove(r, c);
			}
		}
	}
	// update moveset if there are any winning moves.
	if (i > 0) 
		movelen = updateMoves(moveset, movelen, winset, i);
	return movelen;
}


int updateMoves(int moveset[], int movelen, int winset[6], int winlen) {
	int intersection[winlen];
	int i = 0;
	for (int w=0; w<winlen; w++) {
		for (int m=0; m<movelen; m++) {
			if (moveset[m] == winset[w]) {
				intersection[i++] = winset[w];
				continue;
			}
		}
	}
	if (i == 0) 
		return movelen;
	for (int m=0; m<i; m++) {
		if (m < winlen) {
			moveset[m] = intersection[m];
		}
	}
	return i;
}


int gameOver(char board[3][3]) {
// Returns 1 if the game is over. Else returns 0.
	int full = 1;
	for (int row=0; row<3; row++) {
		for (int col=0; col<3; col++) {
			if (board[row][col] != ' ') {
				if ((col == 1) && (row == 1)) {
					if (((board[row-1][col-1] == board[row][col]) && (board[row][col] == board[row+1][col+1])) || ((board[row-1][col+1] == board[row][col]) && (board[row+1][col-1] == board[row][col])))
						return 1;
				}
				if (col == 1) {
					if ((board[row][col-1] == board[row][col]) && (board[row][col] == board[row][col+1]))
						return 1;
				}
				if (row == 1) {
					if ((board[row-1][col] == board[row][col]) && (board[row][col] == board[row+1][col]))
					return 1;
				}
			}
			else 
				full = 0;
		}
	}
	if (full == 0)
		return 0;
	else if (full == 1)
		return 2;
}