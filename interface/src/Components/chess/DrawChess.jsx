import React, { useState, useEffect } from "react";
import ChessUItabs from "./chessTabs";
import ChessData from "./GameData/usefullData";
import ChessTabsFooter from "./chessFooterUI/chessTabsFooter";
import BoardUI from "./board_ui";

import "./chess.scss";

const startingBoard = [
	// default board
	3, 2, 1, 4, 5, 1, 2, 3, 0, 0, 0, 0, 0, 0, 0, 0, 12, 12, 12, 12, 12, 12, 12,
	12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12,
	12, 12, 12, 12, 12, 12, 6, 6, 6, 6, 6, 6, 6, 6, 9, 8, 7, 10, 11, 7, 8, 9,
];

const DrawChess = () => {
	// Draw chess UI (board, footer, tabs)

	// ========================================================================
	// API data / request

	// fetch game-data from API
	const [GSdata, setData] = useState(null); // gamestate state
	const [isLoading, setIsLoading] = useState(true);
	const [error, setError] = useState(null);

	// fetch game-data -> http://localhost:8080/chessgame
	const fetchData = async () => {
		try {
			const response = await fetch("http://localhost:8080/chessgame"); // Replace with your API endpoint
			if (!response.ok) {
				throw new Error(`HTTP error! Status: ${response.status}`);
			}
			const result = await response.json();

			// Decode data and handle it
			const decodedData = ChessData(result);

			setData(decodedData);

			console.log(`message: ${decodedData.message}`);
		} catch (error) {
			setError(error);
		} finally {
			setIsLoading(false);
		}
	};

	// fetch Move -> http://localhost:8080/move
	const sendMove = async (move) => {
		//console.log(`Send move: ${move}`);

		let jsondata = { move: move };
		//console.log(JSON.stringify(jsondata));
		try {
			//const response = await fetch("http://localhost:8080/move");
			const response = await fetch("http://localhost:8080/move", {
				method: "POST",
				headers: {
					"Content-Type": "application/json",
				},
				body: JSON.stringify(jsondata),
			});
			if (!response.ok) {
				throw new Error(`HTTP error! Status: ${response.status}`);
			}
			const result = await response.json();

			// Decode data and handle it
			const decodedData = ChessData(result);

			setData(decodedData);
		} catch (error) {
			setError(error);
		} finally {
			setIsLoading(false);
		}
	};

	// mount fetch data to UI
	useEffect(() => {
		fetchData();
	}, []);

	// ------------------------------------------------------------------------
	// Data handling

	let boardPieces = startingBoard;
	let moveList = {};
	let w_move = true;
	let moveHistory = {};
	let opp_pieces = []; //[6, 7, 8, 9, 10, 11];

	// decode data once loaded
	if (!isLoading && !error) {
		let chessData = GSdata;

		let gameData = chessData.gamestate;
		boardPieces = gameData.board;
		moveList = gameData.movelist;

		w_move = gameData.state.w_move;
		moveHistory = gameData.movehistory;
	}
	// ========================================================================
	// UI functions

	const newGame = () => {
		// reset game
		console.log("New Game");
	};

	const undoMove = () => {
		// undo last move
		console.log("Undo Move");
	};

	const flipBoard = () => {
		// flip board
		console.log("Flip Board");
	};

	// ========================================================================
	// Drawing Component

	let boardLength = 720;

	return (
		<div className="px-3 py-2 chess-ui flex">
			<div className="flex">
				<BoardUI
					boardLength={boardLength}
					boardPieces={boardPieces}
					w_move={w_move}
					movelist={moveList}
					sendMove={sendMove}
					newGame={newGame}
					undo={undoMove}
					flipBoard={flipBoard}
				/>

				<ChessTabsFooter
					moveList={moveList}
					w_move={w_move}
					moveHistory={moveHistory}
				/>
			</div>

			<ChessUItabs />
		</div>
	);
};

export default DrawChess;
