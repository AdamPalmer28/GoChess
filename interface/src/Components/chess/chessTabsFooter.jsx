import React, { useState, useEffect } from "react";
import Nav from "react-bootstrap/Nav";
import "./chess.scss";
// icons
import { Sliders2, Cpu, Activity } from "react-bootstrap-icons";

function ChessTabsFooter() {
	const [activeTab, setActiveTab] = useState("game");

	const handleSelect = (selectedKey) => setActiveTab(selectedKey);
	return (
		<div className="chessfooter d-flex my-3">
			<Footertabs onSelect={handleSelect} />

			<div className="">
				{activeTab === "gamestate" && <div>GameState</div>}
				{activeTab === "ai" && <div>AI Info</div>}
				{activeTab === "settings" && <div>Settings content here</div>}
			</div>
		</div>
	);
}

function Footertabs(props) {
	let image_size = 24;
	return (
		<Nav
			justify
			fill
			variant="tabs"
			className="center-nav-items flex-column"
			style={{ width: "40px" }}
			defaultActiveKey="gamestate"
			onSelect={props.onSelect}
		>
			<Nav.Item>
				<Nav.Link eventKey="ai" className="chess-footer-nav-tab-item">
					<Cpu size={image_size} />
				</Nav.Link>
			</Nav.Item>
			<Nav.Item className="pt-3">
				<Nav.Link eventKey="gamestate" className="chess-footer-nav-tab-item">
					<Activity size={image_size} />
				</Nav.Link>
			</Nav.Item>
			<Nav.Item className="pt-3">
				<Nav.Link eventKey="settings" className="chess-footer-nav-tab-item">
					<Sliders2 size={image_size} />
				</Nav.Link>
			</Nav.Item>
		</Nav>
	);
}

export default ChessTabsFooter;
