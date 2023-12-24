import DrawBoard from "./board";
import ChessUItabs from "./chessTabs";

import "./chess.scss";

const startingBoard = [
	3, 2, 1, 4, 5, 1, 2, 3, 0, 0, 0, 0, 0, 0, 0, 0, 12, 12, 12, 12, 12, 12, 12,
	12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12,
	12, 12, 12, 12, 12, 12, 6, 6, 6, 6, 6, 6, 6, 6, 9, 8, 7, 10, 11, 7, 8, 9,
];
let boardPieces = startingBoard;
const DrawChess = () => {
	const squareSelected = (index) => {
		console.log(`Draw Chess: You clicked square ${index}`);

		if (boardPieces[index] == 12) {
			console.log(`Draw Chess: You clicked empty square`);
			return;
		}
		console.log(`Draw Chess: You clicked piece ${boardPieces[index]}`);
	};
	let boardLength = 920;
	//const playerWhite = bool; // is the player white or black

	return (
		<div className="container chess-ui d-flex">
			<DrawBoard
				onSquareSelect={squareSelected}
				boardLength={boardLength}
				pieces={boardPieces}
			/>
			<ChessUItabs />
		</div>
	);
};

export default DrawChess;
