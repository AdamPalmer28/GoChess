import React, { useState, useEffect } from "react";

const MyComponent = () => {
	const [data, setData] = useState(null);
	const [isLoading, setIsLoading] = useState(true);
	const [error, setError] = useState(null);

	useEffect(() => {
		const fetchData = async () => {
			try {
				const response = await fetch("http://localhost:8080/api/data"); // Replace with your API endpoint
				if (!response.ok) {
					throw new Error(`HTTP error! Status: ${response.status}`);
				}
				console.log(response);
				const result = await response.json();
				setData(result);
			} catch (error) {
				setError(error);
			} finally {
				setIsLoading(false);
			}
		};

		fetchData();
	}, []); // The empty dependency array ensures that the effect runs only once, similar to componentDidMount

	if (isLoading) {
		return <p>Loading...</p>;
	}

	if (error) {
		return (
			<>
				<p>Error: {error.message}</p>
			</>
		);
	}

	return (
		<div>
			<p>Data:</p>
			<pre>{JSON.stringify(data, null, 2)}</pre>
		</div>
	);
};

export default MyComponent;
