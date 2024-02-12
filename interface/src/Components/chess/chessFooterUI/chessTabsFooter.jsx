import React, { useState, useEffect } from "react";
import Nav from "react-bootstrap/Nav";
import "./footer.scss";
// icons
import { Sliders2, Cpu, Activity } from "react-bootstrap-icons";
import GameStateTab from "./gamestateTab";

function ChessTabsFooter(props) {
	const [activeTab, setActiveTab] = useState("ai");

	const handleSelect = (selectedKey) => setActiveTab(selectedKey);
	return (
		<div className="chessFooter d-flex my-2 bg-dark">
			<Footertabs onSelect={handleSelect} />

			<div className="ps-2">
				{activeTab === "ai" && <div>AI Info</div>}
				{activeTab === "gamestate" && (
					<GameStateTab
						moveList={props.moveList}
						w_move={props.w_move}
						moveHistory={props.moveHistory}
					/>
				)}
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
			className="nav-tabs-footer"
			style={{ width: "40px" }}
			defaultActiveKey="ai"
			onSelect={props.onSelect}
		>
			<Nav.Item>
				<Nav.Link eventKey="ai" className="chess-footer-nav-tab-item">
					<Cpu size={image_size} />
				</Nav.Link>
			</Nav.Item>
			<Nav.Item>
				<Nav.Link eventKey="gamestate" className="chess-footer-nav-tab-item">
					<Activity size={image_size} />
				</Nav.Link>
			</Nav.Item>
			<Nav.Item>
				<Nav.Link eventKey="settings" className="chess-footer-nav-tab-item">
					<Sliders2 size={image_size} />
				</Nav.Link>
			</Nav.Item>
		</Nav>
	);
}

export default ChessTabsFooter;
