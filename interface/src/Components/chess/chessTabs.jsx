import Nav from "react-bootstrap/Nav";

function ChessUItabs() {
	return (
		<div className="chess-ui-tabs flex-grow-1">
			<Navtabs />
			<div>Tabs here</div>
		</div>
	);
}

function Navtabs(props) {
	return (
		<Nav justify fill variant="tabs" className="w-100" defaultActiveKey="/home">
			<Nav.Item>
				<Nav.Link eventKey="game" className="chess-nav-tab">
					Game
				</Nav.Link>
			</Nav.Item>
			<Nav.Item>
				<Nav.Link eventKey="link-1" className="chess-nav-tab">
					Analysis
				</Nav.Link>
			</Nav.Item>
			<Nav.Item>
				<Nav.Link eventKey="link-2" className="chess-nav-tab">
					Settings
				</Nav.Link>
			</Nav.Item>
		</Nav>
	);
}

export default ChessUItabs;
