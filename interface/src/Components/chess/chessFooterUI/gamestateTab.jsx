const GameStateTab = (props) => {
	return (
		<div>
			<h5>Game State</h5>
			Move: {props.w_move ? "White" : "Black"}
			<MoveList ml={props.moveList} />
			<MoveHistory mh={props.moveHistory} />
		</div>
	);
};

const MoveList = (props) => {
	// format the move for human readability in the UI
	const formatMove = (move) => {
		return `${move[0]}-${move[1]}`;
	};

	// display the move list
	return (
		<div>
			Moves (ordered by engine): <br />
			<div className="d-inline-flex">
				{props.ml.human.map((move, index) => (
					<div className="text-nowrap" key={index}>
						{formatMove(move)} &nbsp;
					</div>
				))}
			</div>
		</div>
	);
};

const MoveHistory = (props) => {
	return (
		<div>
			Moves History: <br />
			{props.mh.human}
		</div>
	);
};

export default GameStateTab;
