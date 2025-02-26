import DrawBoard from "./board";
import React, { useState } from "react";
import {
	PlusLg,
	ArrowCounterclockwise,
	ArrowDownUp,
} from "react-bootstrap-icons";
import Button from "react-bootstrap/Button";

const BoardUI = (props) => {
	let boardLength = props.boardLength;
	let boardPieces = props.gamestate.board;
	let moveList = props.gamestate.movelist.index;

	// selected squares [from, to]

	const [sqSelected, setSqSelected] = useState(64); // selected squares [from, to]
	const [lastMove, setLastMove] = useState([64, 64]); // last move [from, to]
	const [selectedSqMoves, setSqMoves] = useState([]); // selected square moves

	const resetSquares = () => {
		setSqSelected(64);
		setLastMove([64, 64]);
		setSqMoves([]);
	};

	const [is_flipBoard, changeFlipboard] = useState(false); // gamestate state

	const doFlipBoard = () => {
		changeFlipboard(!is_flipBoard);
	};

	let opp_pieces = [];
	if (props.gamestate.state.w_move) {
		opp_pieces = [6, 7, 8, 9, 10, 11]; // Opponent's pieces for white move
	} else {
		opp_pieces = [0, 1, 2, 3, 4, 5]; // Opponent's pieces for black move
	}

	// square selected / clicked
	const squareSelected = (index) => {
		// square already selected - therefore 2nd click is possible move

		if (sqSelected != 64) {
			let move = [sqSelected, index]; // possible move

			// check if move is valid
			if (selectedSqMoves.includes(index)) {
				// send move to API
				props.userMove(move);
				// set last move
				setLastMove(move);

				setSqSelected(64); // reset selected square
				setSqMoves([]); // reset moves
				return;
			}
			// else - invalid move
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

		// get available moves for square selected
		let new_moves = [];
		for (let i = 0; i < moveList.length; i++) {
			if (moveList[i][0] === index) {
				new_moves.push(moveList[i][1]);
			}
		}
		setSqMoves(new_moves);
	};

	// ------------------------------------------------------------------------
	// helper api functions

	const BoardNewGame = () => {
		// reset game
		props.newGame();

		resetSquares();
	};

	// undo last move
	const BoardUndo = () => {
		props.undo();

		resetSquares();
	};

	// ========================================================================
	// Drawing Component
	return (
		<div id="board-ui" className="d-flex">
			<BoardSettings
				newGame={BoardNewGame}
				undo={BoardUndo}
				flipBoard={doFlipBoard}
			/>

			<DrawBoard
				onSquareSelect={squareSelected}
				analysisHighlight={props.boardBitboards}
				boardLength={boardLength}
				pieces={boardPieces}
				flipBoard={is_flipBoard}
				sqSelected={sqSelected}
				lastMove={lastMove}
				moveOptions={selectedSqMoves}
			/>

			<EvalBar score={2.3} />
		</div>
	);
};

const BoardSettings = (props) => {
	/* Buttons:
		- New Game
		- Undo
		- Flip
		- Upload PGN
		
	*/
	let image_size = 24;

	return (
		<div id="board-settings" className="board-settings">
			<Button
				variant="black-light"
				id="new-game"
				className="setting-btn"
				onClick={props.newGame}
			>
				<PlusLg size={image_size} />
			</Button>
			<Button
				variant="black-light"
				id="Undo"
				className="setting-btn"
				onClick={props.undo}
			>
				<ArrowCounterclockwise size={image_size} />
			</Button>
			<Button
				variant="black-light"
				id="Flip"
				className="setting-btn"
				onClick={props.flipBoard}
			>
				<ArrowDownUp size={image_size} />
			</Button>
		</div>
	);
};

const EvalBar = (props) => {
	let score = props.score;
	let score_height = 0;

	if (Math.abs(score) > 2) {
		score_height = (50 * (score - 1)) / score;
	} else {
		score_height = (25 * score) / 2;
	}

	let top_height = 50 + score_height;

	return (
		<div id="eval-bar" className="eval-bar ms-3 mx-1">
			<div className="score-ticker " style={{ top: `${top_height}%` }} />
			<div id="black-2" className="eval-bar-item black">
				{score < 0 ? score : ""}
			</div>
			<div id="black-1" className="eval-bar-item black " />
			<div id="white-1" className="eval-bar-item white" />
			<div id="white-2" className="eval-bar-item white">
				{score > 0 ? score : ""}
			</div>
		</div>
	);
};

export default BoardUI;
