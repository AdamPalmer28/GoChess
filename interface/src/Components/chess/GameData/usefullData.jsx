import React, { useState, useEffect } from "react";

const columns = ["A", "B", "C", "D", "E", "F", "G", "H"];
const rows = ["1", "2", "3", "4", "5", "6", "7", "8"];

// Creates Human readable and index move from data
const makeMoveLists = (ml) => {
	const IndtoSq = (index) => {
		let row = Math.floor(index / 8) + 1;
		let column = columns[index % 8];
		return column + row;
	};

	if (ml == null) {
		// ?happens at start of game
		return { human: [], index: [] };
	}
	let humanList = [];
	let indexlist = [];

	for (let i = 0; i < ml.length; i++) {
		let startsq = ml[i][1];
		let endsq = ml[i][2];

		indexlist.push([startsq, endsq]);

		// convert to human readable
		humanList.push([IndtoSq(startsq), IndtoSq(endsq)]);
	}

	return { human: humanList, index: indexlist };
};

// Create useful chess data from JSON data (from API)
const ChessData = (data) => {
	let message = data.message;
	let moveList = data.movelist;
	let moveHistory = data.movehistory;
	let board = data.board;
	let w_move = data.w_move;

	// Clean Move lists -
	moveList = makeMoveLists(moveList);
	moveHistory = makeMoveLists(moveHistory);

	return {
		message: message,
		movelist: moveList,
		movehistory: moveHistory,
		board: board,
		w_move: w_move,
	};
};

export default ChessData;
