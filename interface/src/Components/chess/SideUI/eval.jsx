import "./sideUI.scss";

function DrawFunctionBars(props) {
	// draw bar chart of values  bar
	let w_val = parseFloat(props.eval[0]);
	let b_val = parseFloat(props.eval[1]);

	let max_value = props.maxEval;
	console.log(max_value);

	//console.log(w_val, b_val);
	let w_width = Math.abs(Math.round(((w_val / max_value) * 100) / 2));
	let b_width = Math.abs(Math.round(((b_val / max_value) * 100) / 2));

	console.log("Values: ", w_val, b_val);
	console.log("Widths: ", w_width, b_width);

	// build retangles from the center
	return (
		<div className="FunctionEval">
			<div
				className="FunctionEvalBar white"
				style={{
					width: `${w_width}%`,

					marginLeft: w_val >= 0 ? "50%" : "0",
					marginRight: w_val < 0 ? "50%" : "0",

					// align right
					alignSelf: "right",
				}}
			>
				{w_val.toFixed(2)}
			</div>
			<div className="FunctionEvalBar black">{b_val.toFixed(2)}</div>
		</div>
	);
}

function DrawEvalScore(props) {
	// draw eval score

	let evalData = props.eval;

	// slice the first element from the object
	let evalKeys = Object.keys(evalData).slice(1);
	let evalValues = Object.values(evalData).slice(1);

	// get max value of evalValues (taking the absolute value)
	let maxEval = Math.max(...evalValues.flat().map((val) => Math.abs(val)));

	return (
		<div className="analysisPage">
			<div className="w-100">
				<DrawEvalBar eval={evalData.total} />
			</div>
			<table className="evalTable">
				<thead>
					<tr>
						<th>Function</th>
						<th>Score Value</th>
					</tr>
				</thead>
				<tbody>
					{evalKeys.map((key, index) => (
						<tr key={index}>
							<td>{key}</td>
							<td>
								<DrawFunctionBars eval={evalValues[index]} maxEval={maxEval} />
							</td>
						</tr>
					))}
				</tbody>
			</table>
		</div>
	);
}

// Draw horizontal evalBar - on the eval tab
function DrawEvalBar(props) {
	let score = props.eval;
	let score_width = 0;

	if (Math.abs(score) > 2) {
		score_width = (50 * (score - 1)) / score;
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

export default DrawEvalScore;
