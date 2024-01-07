import React, { useState, useEffect } from "react";

const columns = ["A", "B", "C", "D", "E", "F", "G", "H"];
const rows = ["1", "2", "3", "4", "5", "6", "7", "8"];

const IndtoSq = (index) => {
	let row = Math.floor(index / 8) + 1;
	let column = columns[index % 8];
	return column + row;
};

// Create useful chess data from JSON data (from API)
const ChessData = (data) => {
	let message = data.message;
	let moveList = data.movelist;

	// ---- Move lists ----

	let HumanMoveListArray = [];
	let moveListArray = [];

	for (let i = 0; i < moveList.length; i++) {
		let startsq = moveList[i][1];
		let endsq = moveList[i][2];

		moveListArray.push([startsq, endsq]);

		// convert to human readable
		HumanMoveListArray.push([IndtoSq(startsq), IndtoSq(endsq)]);
	}

	return { message, moveListArray, HumanMoveListArray };
};

export default ChessData;
