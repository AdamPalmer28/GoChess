import React, { useState, useEffect } from "react";
import ChessUItabs from "./chessTabs";
import ChessData from "./GameData/gamestate";
import { fetchData, sendNewGame, sendUndo, sendMove } from "./.api/api";
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

	// mount fetch data to UI
	useEffect(() => {
		fetchData(
			"http://localhost:8080/chessgame",
			{},
			setData,
			setError,
			setIsLoading
		);
	}, []);

	// ------------------------------------------------------------------------
	// Data handling

	let boardPieces = startingBoard;
	let moveList = {};
	let w_move = true;
	let moveHistory = {};
	let opp_pieces = []; //[6, 7, 8, 9, 10, 11];
	let evalData = {};

	// decode data once loaded
	if (!isLoading && !error) {
		let chessData = GSdata;

		let gameData = chessData.gamestate;
		boardPieces = gameData.board;
		moveList = gameData.movelist;

		w_move = gameData.state.w_move;
		moveHistory = gameData.movehistory;

		evalData = gameData.evalScore;
	}

	// ========================================================================
	// UI functions

	const newGame = () => sendNewGame(setData, setError, setIsLoading);
	const userMove = (move) => sendMove(move, setData, setError, setIsLoading);
	const undoMove = () => sendUndo(setData, setError, setIsLoading);

	// TODO: flip board
	const flipBoard = () => {
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
					userMove={userMove}
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

			<ChessUItabs eval={evalData} />
		</div>
	);
};

export default DrawChess;
