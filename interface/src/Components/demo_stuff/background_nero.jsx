import React, { useEffect, useRef } from "react";

function BallWithTrail() {
	const canvasRef = useRef(null);
	const trailRef = useRef([]);

	useEffect(() => {
		const canvas = canvasRef.current;
		const ctx = canvas.getContext("2d");
		let x = canvas.width / 2;
		let y = canvas.height / 2;

		function draw() {
			// Draw the background to create the trail effect
			ctx.fillStyle = "rgba(255, 255, 255, 0.05)";
			ctx.fillRect(0, 0, canvas.width, canvas.height);

			// Draw the ball
			ctx.beginPath();
			ctx.arc(x, y, 10, 0, Math.PI * 2);
			ctx.fillStyle = "blue";
			ctx.fill();
			ctx.closePath();

			// Update the trail array with the ball's position
			// trailRef.current.push({ x, y });

			// // Limit the trail length to a specific number
			// if (trailRef.current.length > 40) {
			// 	trailRef.current.shift();
			// }

			// Draw the trail
			// ctx.beginPath();
			// ctx.strokeStyle = "blue";
			// for (const point of trailRef.current) {
			// 	ctx.lineTo(point.x, point.y);
			// }
			//ctx.stroke();

			// Move the ball
			y += 1; // Adjust the speed as needed

			if (x > canvas.width) {
				x = 0;
			}

			if (y > canvas.height) {
				y = 0;
			}

			requestAnimationFrame(draw);
		}

		draw();
		draw();
		draw();
	}, []);

	return (
		<canvas
			ref={canvasRef}
			width={800}
			height={600}
			style={{ border: "3px solid black", position: "absolute" }}
		/>
	);
}

export default BallWithTrail;
