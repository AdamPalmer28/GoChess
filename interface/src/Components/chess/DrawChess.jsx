import React, { useState, useEffect } from "react";
import DrawBoard from "./board";
import ChessUItabs from "./chessTabs";
import ChessData from "./GameData/usefullData";
import ChessTabsFooter from "./chessFooterUI/chessTabsFooter";

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
	const [GSdata, setData] = useState(null);
	const [isLoading, setIsLoading] = useState(true);
	const [error, setError] = useState(null);

	// fetch game-data -> http://localhost:8080/chessgame
	const fetchData = async () => {
		try {
			const response = await fetch("http://localhost:8080/chessgame"); // Replace with your API endpoint
			if (!response.ok) {
				throw new Error(`HTTP error! Status: ${response.status}`);
			}
			console.log(response);
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
		console.log(`Send move: ${move}`);
		try {
			// ! this request is not working
			const response = await fetch("http://localhost:8080/move", {
				method: "GET",
				headers: {
					"Content-Type": "application/json",
				},
				body: JSON.stringify(move),
			});
			if (!response.ok) {
				throw new Error(`HTTP error! Status: ${response.status}`);
			}
			console.log(response);
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

	// mount fetch data to UI
	useEffect(() => {
		fetchData();
	}, []);

	let boardPieces = startingBoard;
	let moveList = {};
	let w_move = true;
	let moveHistory = {};
	let opp_pieces = [6, 7, 8, 9, 10, 11];

	// decode data once loaded
	if (!isLoading && !error) {
		let chessData = GSdata;

		let gameData = chessData.gamestate;
		boardPieces = gameData.board;
		moveList = gameData.movelist;

		w_move = gameData.state.w_move;
		moveHistory = gameData.movehistory;

		let opp_pieces = [];

		if (w_move) {
			opp_pieces = [6, 7, 8, 9, 10, 11]; // Opponent's pieces for white move
		} else {
			opp_pieces = [0, 1, 2, 3, 4, 5]; // Opponent's pieces for black move
		}
	}

	// ========================================================================
	// Interactions

	// selected squares [from, to]
	const [sqSelected, setSqSelected] = useState(64); // selected squares [from, to]
	const [lastMove, setLastMove] = useState([64, 64]); // last move [from, to]
	const [selectedSqMoves, setSqMoves] = useState([]); // selected square moves

	// square selected / clicked
	const squareSelected = (index) => {
		// square already selected - therefore 2nd click is possible move
		if (sqSelected != 64) {
			let move = [sqSelected, index]; // possible move

			// check if move is valid
			if (selectedSqMoves.includes(index)) {
				console.log("Valid Move");

				// send move to API

				sendMove(move);

				setSqSelected(64); // reset selected square
				setSqMoves([]); // reset moves
				return;
			} else {
				console.log("Invalid Move");
			}
		}

		// (empty square) or (selected opponent pieces)
		let piece_selected = boardPieces[index];
		if (piece_selected == 12 || opp_pieces.includes(piece_selected)) {
			setSqSelected(64); // reset selected square
			setSqMoves([]); // reset moves
			return;
		}

		// clicked on a piece
		setSqSelected(index);

		// get available moves
		let new_moves = [];
		for (let i = 0; i < moveList.index.length; i++) {
			if (moveList.index[i][0] === index) {
				new_moves.push(moveList.index[i][1]);
			}
		}
		setSqMoves(new_moves);
		//console.log(`Moves: ${new_moves}`);
	};

	// ========================================================================
	// Drawing Component

	let boardLength = 720;
	//const playerWhite = bool; // is the player white or black

	return (
		<div className="px-3 py-2 chess-ui flex">
			<div className="flex" style={{ width: boardLength }}>
				<DrawBoard
					onSquareSelect={squareSelected}
					boardLength={boardLength}
					pieces={boardPieces}
					sqSelected={sqSelected}
					lastMove={lastMove}
					moveOptions={selectedSqMoves}
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
