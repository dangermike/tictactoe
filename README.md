# Tic Tac Toe

In giving interviews for architecture or coding, I have recently been asking candidates to build tic-tac-toe. This could be writing the core engine of the game, an API, or a whole system to allow millions of simultaneous players locked in a global competition. Because of its simplicity, tic-tac-toe allows the sessions to get into non-trivial design challenges without the burden of needing to figure out what the problem means.

After hearing many candidates address this silly little game, it seemed only fair for me to do the same. Unlike those sitting across the table from me, I am under no time constraints. I have the opportunity to build tic-tac-toe unfettered by reason, unbound to purpose, and completely ridiculous. Since this is all for fun, ridiculousness in the implementation should be seen as a virtue.

# Game Implementation

## Board

With the unreasonable goal of supporting millions of games, memory efficiency is paramount. To that end, the board is represented as a single `uint32` and is immutable. This allows the board to be cheap to pass around and/or serialize. Not using a slice keeps all data on the stack. The board "number" uses 2 bits per square, 18 bits of total information. It is unclear if I'm going to use the other 14 bits, but we'll see. Each square can hold only 3 states, but using 2 bits allows the board to be manipulated with bit shifts, rather than `div` and `mod` by 3. These bit shifts should be significantly faster than the integer math.

The board has some special operations, specifically `Rotate` and `Minimize`. Rotations are transformations of the board. Minimization is rotating (or flipping) the board to minimize the integer representation. This allows all equivalent board positions to be transformed into a single, common value. Being honest, this should probably not be on the board as the board doesn't need this ability to run the game. Rather, it is useful to reduce the cases checked by automated players. I will likely factor these out into some kind of tic-tac-toe utility package.

## Game
Games contain a board and know the current player. While the board can tell whether a move is valid, it is up to the game to enforce the rules about turns. There isn't much in there, and likely won't be. Authentication and such will be handled elsewhere.

## Player
The `Player` interface is what defines a player in the game. Some players carry no state, but most intelligent players will have more to them than that. There are implementations in this project for both automated players and a terminal-based player. The plan is to build an API-exposing `Player` that will be used to communicate with outside agents.

* `terminal`: A player to get instructions from a human player at the command line. This shows the current state of the board to the user and parses the moves they make.
* `dumb`: This player just looks for an empty square at random. Useful for testing that things are working, the dumb player isn't much of an opponent.
* `learning`: Internally, this holds a map of boards it has seen to how well it has done in choosing a move on that board state. This player takes a long time to train, but can get to be pretty good. When trained, it has learned about all 4,480 possible board states.
* `learningminimizing`: Like the `learning` player, this keeps a map of boards to choice values, but it minimizes (rotates) the board before evaluating. That means that this player can learn much faster and carry far fewer states than the normal `learning` player. This player only has 956 states to track, about 1/5 as many as its non-minimizing counterpart.
* `heuristic`: (INCOMPLETE) Tic-tac-toe is a solved game. The rules are known and well-defined. This player knows those rules and simply applies them. Unless I make a mistake, it should never lose regardless of marker (X or O), and should be able to beat human players that haven't figured out all the rules yet.

If everything goes according to plan, the `learning` and `learningminimizing` players should achieve parity with the `heuristic` player.

# Main? Hardly.

At the time of this writing, `main` is a dumping ground of hacky, messy stuff to train and test the automated players. The plans for this are pretty grandiose.

## Functional areas

* Lobby -- where players can go to get matched for a game
* Gameplay -- transferred from the lobby, games happen here
* Scoreboard -- who's good, who plays

## Features:

* Millions of simultaneous games
* Multi-host
* Zero down time maintenance. I'd also like to allow games to proceed in the face of servers crashing, but I haven't decided how I want to do that yet.
* Signed requests and responses, allowing games to proceed without reaching out to other servers
* Possibly tracking moves. Since a game is just a series of numbers from 0-9, the whole thing can be held in only 29 bits.
