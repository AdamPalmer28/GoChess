package moves

/*
Moves are represented as a 16 bit integer
0000 000000 000000
special | finish | start

0000 - special moves (info below)
000000 - index of square

special moves
0000 - quite move
0001 - double pawn push
0010 - king side castle
0011 - queen side castle
0100 - capture
0101 - enpassent capture
1000 - promotion knight
1001 - promotion bishop
1010 - promotion rook
1011 - promotion queen
1100 - promotion knight capture
1101 - promotion bishop capture
1110 - promotion rook capture
1111 - promotion queen capture

*/




