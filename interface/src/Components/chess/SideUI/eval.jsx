import "./sideUI.scss";

// draw Eval tab on side UI -- Root div
function DrawEvalTab(props) {
	let evalData = props.eval;

	return (
		<div className="analysisPage">
			<div className="w-100">
				<DrawEvalBar eval={evalData.total} />
			</div>
			<DrawEvalFunctionTable evalData={evalData} />
		</div>
	);
}

// draw table of eval paired scores
function DrawEvalFunctionTable(props) {
	let evalData = props.evalData;

	// slice the first element from the object
	let evalKeys = Object.keys(evalData).slice(1);
	let evalValues = Object.values(evalData).slice(1);
	// get max value of evalValues (taking the absolute value)
	let maxEval = Math.max(...evalValues.flat().map((val) => Math.abs(val)));

	return (
		<table className="evalTable">
			{/* <thead>
				<tr>
					<th>Function</th>
					<th>Score Value</th>
				</tr>
			</thead> */}
			<tbody>
				{evalKeys.map((key, index) => (
					<tr key={index}>
						<td>{key}</td>
						<td>
							{evalValues[index].length > 1 ? (
								<DrawFunctionBars eval={evalValues[index]} maxEval={maxEval} />
							) : (
								<>{evalValues[index]}</>
							)}
						</td>
					</tr>
				))}
			</tbody>
		</table>
	);
}

// draw 2 (white and black) evaluation pairs blocks
function DrawFunctionBars(props) {
	let w_val = parseFloat(props.eval[0]);
	let b_val = -parseFloat(props.eval[1]);

	let max_value = props.maxEval;

	let w_width = Math.abs(Math.round(((w_val / max_value) * 100) / 2));
	let b_width = Math.abs(Math.round(((b_val / max_value) * 100) / 2));

	function getStyle(width, evalVal) {
		const w_eval = evalVal >= 0; // does eval favour white?
		return {
			width: `${width}%`,

			marginLeft: w_eval ? "50%" : `${50 - width}%`,
			textAlign: !w_eval ? "right" : "left",
		};
	}

	// build retangles from the center
	return (
		<div className="FunctionEval">
			<div
				className="FunctionEvalBar white transition"
				style={getStyle(w_width, w_val)}
			>
				{w_val.toFixed(2)}
			</div>
			<div
				className="FunctionEvalBar black transition"
				style={getStyle(b_width, b_val)}
			>
				{b_val.toFixed(2)}
			</div>
		</div>
	);
}

// Draw a single horizontal evalBar
function DrawEvalBar(props) {
	let score = props.eval;
	let score_width = 0;

	if (Math.abs(score) > 2) {
		score_width = (50 * (Math.abs(score) - 1)) / score;
	} else {
		score_width = (25 * score) / 2;
	}

	// round score to 2 decimal places
	score = parseFloat(score);
	score = score.toFixed(2);

	let top_width = 50 + score_width;

	return (
		<div id="analysis-eval-bar" className="analysis-eval-bar my-3">
			<div
				className="analysis-evalscore-ticker"
				style={{ left: `${top_width}%` }}
			/>
			<div id="analysis-eval-black-2" className="analysis-eval-bar-item black">
				{score < 0 ? score : ""}
			</div>
			<div
				id="analysis-eval-black-1"
				className="analysis-eval-bar-item black "
			/>
			<div
				id="analysis-eval-white-1"
				className="analysis-eval-bar-item white"
			/>
			<div id="analysis-eval-white-2" className="analysis-eval-bar-item white">
				{score >= 0 ? score : ""}
			</div>
		</div>
	);
}

export default DrawEvalTab;
