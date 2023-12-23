import DrawChess from "../Components/chess/DrawChess";
import MyComponent from "../Components/demo";
import BallWithTrail from "../Components/chess/background_nero";
import { Button, Navbar, Container } from "react-bootstrap";

const Home = () => {
	return (
		<>
			<MyComponent />
			<DrawChess />
		</>
	);

	//return <BallWithTrail />;
};

export default Home;
