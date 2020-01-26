# GO Chess

Inspired by: https://www.geeksforgeeks.org/design-a-chess-game/
https://codereview.stackexchange.com/questions/71790/design-a-chess-game-using-object-oriented-principles

x y => x y

1 0 => 7 6

2 0 => 0 2

2 3 => 0 1

2 3 => 6 7

### Gui

https://github.com/fyne-io/fyne


- [ ] Player chooses piece to move.
- [ ] Piece makes legal move according to its own move rules.
- [ ] In addition to purely move-based rules, there's also capture logic, so a bishop cannot move from a1-h8 if there's a piece sitting on c3.
- [ ] If the player was previous under check and the move does not remove the check, it must be undone.
- [ ] If the move exposes check, it must be undone / disallowed.
- [ ] If player captures a piece, remove the piece (including en passant!)
- [ ] If the piece is a pawn reaching the back rank, promote it.
- [ ] If the move is a castling, set the new position of the rook accordingly. But a king and rook can only castle if they haven't moved, so you need to keep track of that. And if the king moves through a check to castle, that's disallowed, too.
- [ ] If the move results in a stalemate or checkmate, the game is over.