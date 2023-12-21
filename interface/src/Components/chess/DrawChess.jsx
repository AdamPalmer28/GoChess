import DrawBoard from "./board";

const DrawChess = () => {
	const squareSelected = (index) => {
		console.log(`Draw Chess: You clicked square ${index}`);
	};
	let boardLength = 920;
	return (
		<div class="container Chess">
			<div className="Chess-board d-flex justify-content-center">
				<DrawBoard onSquareSelect={squareSelected} boardLength={boardLength} />
			</div>
		</div>
	);
};

export default DrawChess;
