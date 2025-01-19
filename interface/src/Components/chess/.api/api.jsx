// api.js
export const fetchData = async (url, data, setData, setError) => {
	//setIsLoading(true);
	try {
		const response = await fetch(url, data);
		if (!response.ok) {
			throw new Error(`HTTP error! Status: ${response.status}`);
		}
		const result = await response.json();
		setData(result);
	} catch (error) {
		setError(error);
	}
};

export const sendMove = (move, setData, setError) => {
	const jsondata = { move };
	fetchData(
		"http://localhost:8080/move",
		{
			method: "POST",
			headers: { "Content-Type": "application/json" },
			body: JSON.stringify(jsondata),
		},
		setData,
		setError
	);
};

export const sendUndo = (setData, setError) => {
	fetchData(
		"http://localhost:8080/undo",
		{
			method: "POST",
			headers: { "Content-Type": "application/json" },
			body: JSON.stringify({}),
		},
		setData,
		setError
	);
};

export const sendNewGame = (setData, setError) => {
	fetchData(
		"http://localhost:8080/newgame",
		{
			method: "POST",
			headers: { "Content-Type": "application/json" },
			body: JSON.stringify({}),
		},
		setData,
		setError
	);
};

export const getAnalysis = async (setAnalysisData, setError) => {
	try {
		const response = await fetch("http://localhost:8080/analysis", {
			method: "GET",
			headers: { "Content-Type": "application/json" },
		});
		if (!response.ok) {
			throw new Error(`HTTP error! Status: ${response.status}`);
		}
		const result = await response.json();
		setAnalysisData(result);
	} catch (error) {
		setError(error);
	}
};
