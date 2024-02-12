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

	// Clean Move lists -
	let moveList = makeMoveLists(data.gamestate.movelist);
	let moveHistory = makeMoveLists(data.gamestate.movehistory);

	// overwrite moveList and moveHistory
	data.gamestate.movelist = moveList;
	data.gamestate.movehistory = moveHistory;

	return data;
};

export default ChessData;
