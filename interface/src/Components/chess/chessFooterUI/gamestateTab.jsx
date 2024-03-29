import "./footer.scss";

const GameStateTab = (props) => {
	return (
		<div>
			<h5 className="tab-title">GameState data</h5>
			Move:{" "}
			<span className="gamestate-text">{props.w_move ? "White" : "Black"}</span>
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
			<div className="ps-3 move-list">
				{props.ml.human.map((move, index) => (
					<div className="move-text" key={index}>
						{formatMove(move)}
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
