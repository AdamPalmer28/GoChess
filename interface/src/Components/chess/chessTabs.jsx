import React, { useState } from "react";
import Nav from "react-bootstrap/Nav";

function ChessUItabs() {
	const [activeTab, setActiveTab] = useState("game");

	const handleSelect = (selectedKey) => setActiveTab(selectedKey);
	return (
		<div className="chess-ui-tabs flex-grow-1">
			<Navtabs onSelect={handleSelect} />

			{activeTab === "game" && <div>Game content here</div>}
			{activeTab === "link-1" && <div>Analysis content here</div>}
			{activeTab === "link-2" && <div>Settings content here</div>}
		</div>
	);
}

function Navtabs(props) {
	return (
		<Nav
			fill
			variant="tabs"
			className="w-100 border-0 background"
			defaultActiveKey="game"
			onSelect={props.onSelect}
		>
			<Nav.Item>
				<Nav.Link eventKey="game" className="chess-nav-tab-item">
					Game
				</Nav.Link>
			</Nav.Item>
			<Nav.Item>
				<Nav.Link eventKey="link-1" className="chess-nav-tab-item">
					Analysis
				</Nav.Link>
			</Nav.Item>
			<Nav.Item>
				<Nav.Link eventKey="link-2" className="chess-nav-tab-item">
					Settings
				</Nav.Link>
			</Nav.Item>
		</Nav>
	);
}

export default ChessUItabs;
