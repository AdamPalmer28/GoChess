// api.js
export const fetchData = async (url, data, setData, setError, setIsLoading) => {
	setIsLoading(true);
	try {
		const response = await fetch(url, data);
		if (!response.ok) {
			throw new Error(`HTTP error! Status: ${response.status}`);
		}
		const result = await response.json();
		setData(result);
	} catch (error) {
		setError(error);
	} finally {
		setIsLoading(false);
	}
};

export const sendMove = (move, setData, setError, setIsLoading) => {
	const jsondata = { move };
	fetchData(
		"http://localhost:8080/move",
		{
			method: "POST",
			headers: { "Content-Type": "application/json" },
			body: JSON.stringify(jsondata),
		},
		setData,
		setError,
		setIsLoading
	);
};

export const sendUndo = (setData, setError, setIsLoading) => {
	fetchData(
		"http://localhost:8080/undo",
		{
			method: "POST",
			headers: { "Content-Type": "application/json" },
			body: JSON.stringify({}),
		},
		setData,
		setError,
		setIsLoading
	);
};

export const sendNewGame = (setData, setError, setIsLoading) => {
	fetchData(
		"http://localhost:8080/newgame",
		{
			method: "POST",
			headers: { "Content-Type": "application/json" },
			body: JSON.stringify({}),
		},
		setData,
		setError,
		setIsLoading
	);
};
