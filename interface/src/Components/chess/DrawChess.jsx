import React, { useState, useEffect } from "react";
import DrawBoard from "./board";
import ChessUItabs from "./chessTabs";
import ChessData from "./GameData/usefullData";
import ChessTabsFooter from "./chessTabsFooter";

import "./chess.scss";

const startingBoard = [
	// default board
	3, 2, 1, 4, 5, 1, 2, 3, 0, 0, 0, 0, 0, 0, 0, 0, 12, 12, 12, 12, 12, 12, 12,
	12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12,
	12, 12, 12, 12, 12, 12, 6, 6, 6, 6, 6, 6, 6, 6, 9, 8, 7, 10, 11, 7, 8, 9,
];

const DrawChess = () => {
	const [sqSelected, setSqSelected] = useState(64); // selected squares [from, to]
	const [lastMove, setLastMove] = useState([64, 64]); // last move [from, to]

	const [GSdata, setData] = useState(null);
	const [isLoading, setIsLoading] = useState(true);
	const [error, setError] = useState(null);

	// fetch game-data from API url = http://localhost:8080/chessgame
	useEffect(() => {
		const fetchData = async () => {
			try {
				const response = await fetch("http://localhost:8080/chessgame"); // Replace with your API endpoint
				if (!response.ok) {
					throw new Error(`HTTP error! Status: ${response.status}`);
				}
				console.log(response);
				const result = await response.json();
				setData(result);
			} catch (error) {
				setError(error);
			} finally {
				setIsLoading(false);
			}
		};
		fetchData();
	}, []);

	let boardPieces = startingBoard;

	// decode data once loaded
	if (!isLoading || error) {
		let gameData = ChessData(GSdata);
		console.log(`message: ${gameData.message}`);
		//console.log(`movelist: ${gameData.moveList}`);
		console.log(`board: ${gameData.board}`);
		console.log(`w_move: ${gameData.w_move}`);
		boardPieces = gameData.board;
	}

	const squareSelected = (index) => {
		// square selected / clicked
		if (sqSelected != 64) {
			console.log(`UserMove: ${sqSelected} -> ${index}`);

			setSqSelected(64); // reset selected square
			return;
		}

		if (boardPieces[index] == 12) {
			// empty square

			setSqSelected(64); // reset selected square
			return;
		}
		// clicked on a piece
		setSqSelected(index);
	};

	let boardLength = 720;
	//const playerWhite = bool; // is the player white or black

	return (
		<div className="px-3 py-2 chess-ui flex">
			<div className="flex">
				<DrawBoard
					onSquareSelect={squareSelected}
					boardLength={boardLength}
					pieces={boardPieces}
					sqSelected={sqSelected}
					lastMove={lastMove}
				/>
				<ChessTabsFooter />
			</div>
			<ChessUItabs />
		</div>
	);
};

export default DrawChess;
