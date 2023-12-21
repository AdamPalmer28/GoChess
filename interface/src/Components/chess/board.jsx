import Container from "react-bootstrap/Container";
import Row from "react-bootstrap/Row";
import Col from "react-bootstrap/Col";
import React, { useState } from "react";

import "./chess.scss";

// Chess Square
const DrawSquare = (prop) => {
	let sqLength = prop.sqLength; // square length

	// use variables from chess.scss
	let color_sq_sass = "black-sq";
	if ((prop.index + ((Math.floor(prop.index / 8) - 1) % 2)) % 2 === 0) {
		// if white square
		color_sq_sass = "white-sq";
	}

	return (
		<Col id={prop.index} className="m-0 p-0 ">
			<div
				className={`${color_sq_sass}  ${prop.highlightSq ? "highlight" : ""}`}
				style={{ height: sqLength, width: sqLength }}
				onClick={() => {
					// left click
					prop.onLeftClick(prop.index, prop.id);
				}}
				onContextMenu={(e) => {
					// right click
					e.preventDefault(); // prevent context menu from showing
					prop.onContextMenu(prop.index);
				}}
			>
				{prop.id}
			</div>
		</Col>
	);
};

// Chess board
const DrawBoard = (props) => {
	const columns = ["A", "B", "C", "D", "E", "F", "G", "H"];
	const rows = ["1", "2", "3", "4", "5", "6", "7", "8"];

	//const playerWhite = bool; // is the player white or black

	let boardSize = props.boardLength; // board size
	let sqLength = boardSize / 8; // square length
	const [highlight_sq, setHighlight] = useState(Array(64).fill(false));

	// left click on square
	function LeftSqClick(index, id) {
		props.onSquareSelect(index);
		resetHighlight(); // reset all highlighted squares
	}

	// right click on square
	function highlightSquare(index) {
		console.log("right click ", index);
		let highlight_sq_copy = [...highlight_sq]; // copy array
		highlight_sq_copy[index] = !highlight_sq_copy[index]; // set square to true
		setHighlight(highlight_sq_copy); // update state
	}

	// reset all highlighted squares
	function resetHighlight() {
		// reset all highlighted state on squares to false
		setHighlight(Array(64).fill(false));
	}

	function CreateRow(props) {
		let row = props.row; // row number
		let RowSquareDivs = [];

		// loop through columns of a row
		columns.forEach((column, i) => {
			const squareNum = (row - 1) * 8 + i; // square number
			const square = column + row; // square name

			RowSquareDivs.push(
				<DrawSquare
					id={square}
					index={squareNum}
					onLeftClick={LeftSqClick}
					onContextMenu={highlightSquare}
					sqLength={sqLength}
					highlightSq={highlight_sq[squareNum]}
				></DrawSquare>
			);
		});

		return (
			<Row
				id={`row ${rows[row]}`}
				class="d-inline-block position-absolute"
				style={{ width: boardSize }}
			>
				{RowSquareDivs}
			</Row>
		);
	}

	const rowsRender = rows.reverse().map((row, index) => {
		return <CreateRow row={row} key={index} />;
	});

	return (
		<div id="board-layout">
			<Container fluid className="grid" id="board">
				{rowsRender}
			</Container>
		</div>
	);
};
// className="flipped"
export default DrawBoard;
