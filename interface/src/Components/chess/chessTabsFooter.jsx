import React, { useState, useEffect } from "react";
import Nav from "react-bootstrap/Nav";
import "./chess.scss";

function ChessTabsFooter() {
	const [activeTab, setActiveTab] = useState("game");

	const handleSelect = (selectedKey) => setActiveTab(selectedKey);
	return (
		<div className="chessfooter w-10 d-flex">
			<Footertabs onSelect={handleSelect} />

			{activeTab === "game" && <div>Changed</div>}
			{activeTab === "link-1" && <div>Analysis content here</div>}
			{activeTab === "link-2" && <div>Settings content here</div>}
		</div>
	);
}

function Footertabs(props) {
	return (
		<Nav
			justify
			fill
			variant="tabs"
			className="flex-column bg-primary px-1 center-nav-items"
			style={{ width: "40px" }}
			defaultActiveKey="game"
			onSelect={props.onSelect}
		>
			<Nav.Item className="pt-2">
				<Nav.Link eventKey="game" className="chess-footer-nav-tab-item">
					1
				</Nav.Link>
			</Nav.Item>
			<Nav.Item className="pt-2">
				<Nav.Link eventKey="link-1" className="chess-footer-nav-tab-item">
					2
				</Nav.Link>
			</Nav.Item>
			<Nav.Item className="py-2">
				<Nav.Link eventKey="link-2" className="chess-footer-nav-tab-item">
					3
				</Nav.Link>
			</Nav.Item>
		</Nav>
	);
}

export default ChessTabsFooter;
