import React, { useState, useEffect } from "react";
import ChessUItabs from "./chessTabs";
import {
	fetchData,
	sendNewGame,
	sendUndo,
	sendMove,
	getAnalysis,
} from "./.api/api";
import ChessTabsFooter from "./chessFooterUI/chessTabsFooter";
import BoardUI from "./board_ui";

import "./chess.scss";

const startingBoard = [
	// default board
	3, 2, 1, 4, 5, 1, 2, 3, 0, 0, 0, 0, 0, 0, 0, 0, 12, 12, 12, 12, 12, 12, 12,
	12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12,
	12, 12, 12, 12, 12, 12, 6, 6, 6, 6, 6, 6, 6, 6, 9, 8, 7, 10, 11, 7, 8, 9,
];

let evalData = {};

let defaults_GS = {
	board: startingBoard,
	movelist: {
		index: [],
		human: [],
	},
	state: {
		w_move: true,
	},
	movehistory: {
		index: [],
		human: [],
	},
	opp_pieces: [],
};

const DrawChess = () => {
	// Draw chess UI (board, footer, tabs)

	// ========================================================================
	// API data / request

	// fetch game-data from API
	const [GSdata, setData] = useState(null); // gamestate state
	const [AnalysisData, setAnalysisData] = useState(null); // analysis parameters
	const [AIData, setAIData] = useState(null); // AI/search data

	const [isLoading, setIsLoading] = useState(true);
	const [error, setError] = useState(null);

	// mount fetch data to UI
	useEffect(() => {
		fetchData("http://localhost:8080/chessgame", {}, setData, setError);
		getAnalysis(setAnalysisData, setError);
		setIsLoading(false);
	}, []);

	// ------------------------------------------------------------------------
	// Data handling

	let bitbardData = {
		red: [1, 2, 3],
		green: [3, 10, 15, 16],
		blue: [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16],
	};

	if (AnalysisData !== null) {
		evalData = AnalysisData.evalScore;
	}
	// ========================================================================
	// UI functions

	const newGame = async () => {
		await sendNewGame(setData, setError);
		getAnalysis(setAnalysisData, setError);
	};
	const userMove = async (move) => {
		await sendMove(move, setData, setError);
		getAnalysis(setAnalysisData, setError);
	};
	const undoMove = async () => {
		await sendUndo(setData, setError);
		getAnalysis(setAnalysisData, setError);
	};

	// ========================================================================
	// Drawing Component

	let boardLength = 720;

	// TODO: pass gameData to board UI
	return (
		<div className="px-3 py-2 chess-ui flex">
			<div className="flex">
				<BoardUI
					gamestate={GSdata ? GSdata.gamestate : defaults_GS}
					boardBitboards={bitbardData}
					boardLength={boardLength}
					userMove={userMove}
					newGame={newGame}
					undo={undoMove}
				/>

				<ChessTabsFooter gamestate={GSdata ? GSdata.gamestate : defaults_GS} />
			</div>

			<ChessUItabs eval={evalData} />
		</div>
	);
};

export default DrawChess;
