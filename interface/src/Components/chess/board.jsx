import Container from "react-bootstrap/Container";
import Row from "react-bootstrap/Row";
import Col from "react-bootstrap/Col";
import React, { useState } from "react";

import "./chess.scss";

// Chess board
const DrawBoard = (props) => {
	const columns = ["A", "B", "C", "D", "E", "F", "G", "H"];
	const rows = ["1", "2", "3", "4", "5", "6", "7", "8"];

	let analysisHighlight = props.analysisHighlight.green;

	// board prorperties
	let boardSize = props.boardLength; // board size
	let sqLength = boardSize / 8; // square length

	const [highlight_sq, setHighlight] = useState(Array(64).fill(false)); // right click highlight squares

	// highlight variables
	let lastMoveFrom = props.lastMove[0];
	let lastMoveTo = props.lastMove[1];
	let sqSelected = props.sqSelected;
	let moveOptions = props.moveOptions;

	// helper functions -------------------------------------------------------

	// Sq - left click
	function LeftSqClick(index, id) {
		props.onSquareSelect(index);
		resetHighlight(); // reset all highlighted squares
	}

	// Sq - right click
	function highlightSquare(index) {
		let highlight_sq_copy = [...highlight_sq]; // copy array

		highlight_sq_copy[index] = !highlight_sq_copy[index]; // set square to true

		setHighlight(highlight_sq_copy); // update highligh state of the board state
	}

	// reset all highlighted state on squares to false
	function resetHighlight() {
		setHighlight(Array(64).fill(false));
	}

	// Draw board -------------------------------------------------------------

	// Intialise row of board - divs of squares
	function CreateRow(props) {
		let row = props.row; // row number
		let RowSquareDivs = [];

		// loop through columns of a row
		columns.forEach((column, i) => {
			// create propserties for square
			const squareNum = (row - 1) * 8 + i;
			const square = column + row;

			RowSquareDivs.push(
				<DrawSquare
					id={square}
					index={squareNum}
					onLeftClick={LeftSqClick}
					onContextMenu={highlightSquare}
					sqLength={sqLength}
					piece={props.pieces[squareNum]}
					highlightSq={highlight_sq[squareNum]}
					styleClasses={
						(squareNum == sqSelected ? "selected " : "") +
						(moveOptions.includes(squareNum) ? "move-option " : "") +
						(squareNum == lastMoveFrom ? "last-move-from " : "") +
						(squareNum == lastMoveTo ? "last-move-to " : "") +
						(analysisHighlight.includes(squareNum) ? "analysis-highlight " : "")
					}
				></DrawSquare>
			);
		});

		return (
			<Row id={`row ${rows[row]}`} style={{ width: boardSize }}>
				{RowSquareDivs}
			</Row>
		);
	}
	// Render rows of board - reverse order so that row 1 is at the bottom
	let row_order = rows;
	if (!props.flipBoard) {
		row_order = rows.reverse();
	}

	const rowsRender = row_order.map((row, index) => {
		return <CreateRow row={row} key={index} pieces={props.pieces} />;
	});

	return (
		<div
			id="board-ui"
			className={`${props.className}`}
			style={{ width: boardSize }}
		>
			<div fluid className="chess-board grid" id="board">
				{rowsRender}
			</div>
		</div>
	);
};

// ============================================================================
// Square

// Chess Square
const DrawSquare = (props) => {
	let sqLength = props.sqLength; // square length

	let pieceHeight, pieceWidth;
	// determine piece scaling
	if (props.piece != 12) {
		switch (props.piece % 6) {
			case 0: // pawn
				pieceHeight = 0.7;
				pieceWidth = 0.6;
				break;
			case 1: // knight
				pieceHeight = 0.75;
				pieceWidth = 0.72;
				break;
			case 2: // bishop
				pieceHeight = 0.75;
				pieceWidth = 0.75;
				break;
			case 3: // rook
				pieceHeight = 0.75;
				pieceWidth = 0.7;
				break;
			case 4: // queen
				pieceHeight = 0.75;
				pieceWidth = 0.75;
				break;
			case 5: // king
				pieceHeight = 0.75;
				pieceWidth = 0.75;
				break;
		}
	}

	// determine square color
	let color_sq_sass = "black square";
	if ((props.index + ((Math.floor(props.index / 8) - 1) % 2)) % 2 === 0) {
		color_sq_sass = "white square";
	}

	return (
		<Col id={props.index} className="m-0 p-0">
			{/* SQUARE ------------------------------------*/}
			<div
				className={`${color_sq_sass}  ${props.highlightSq ? "highlight" : ""} 
					${props.styleClasses}`}
				style={{
					height: sqLength,
					width: sqLength,
					position: "relative",
					display: "flex",
					justifyContent: "center",
					alignItems: "center",
					zIndex: 0,
				}}
				onClick={() => {
					// left click
					props.onLeftClick(props.index, props.id);
				}}
				onContextMenu={(e) => {
					// right click
					e.preventDefault(); // prevent context menu from showing
					props.onContextMenu(props.index);
				}}
			>
				{/* ID / TEXT -----------------------------*/}
				{props.index % 8 == 0 || Math.floor(props.index / 8) == 0 ? (
					<div
						style={{
							position: "absolute",
							bottom: 3,
							left: 3,
							fontSize: "0.85rem",
							zIndex: 2,
						}}
					>
						{props.id}
					</div>
				) : (
					<></>
				)}

				{/* PIECE ---------------------------------*/}
				{props.piece != 12 ? (
					<DrawPiece
						piece={props.piece}
						style={{
							height: pieceHeight * sqLength,
							width: pieceWidth * sqLength,
							position: "absolute",
							zIndex: 1,
						}}
					/>
				) : (
					<></>
				)}
			</div>
		</Col>
	);
};

const piecePath = {
	0: "w_pawn",
	1: "w_knight",
	2: "w_bishop",
	3: "w_rook",
	4: "w_queen",
	5: "w_king",
	6: "b_pawn",
	7: "b_knight",
	8: "b_bishop",
	9: "b_rook",
	10: "b_queen",
	11: "b_king",
};

// Chess piece
const DrawPiece = (props) => {
	let piece = piecePath[props.piece];
	let imgPath = `src/assets/chess_pieces/${piece}.png`;

	return <img src={imgPath} alt={piece} style={props.style} />;
};

// className="flipped"
export default DrawBoard;
