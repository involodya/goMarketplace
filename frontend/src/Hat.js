import {Link} from "react-router-dom";
import React from "react";

import './Hat.css';

function Hat()  {
    const handleExitClick = () => {
        window.localStorage.setItem("access", "");
        window.location.reload();
    };

    return (
        <header className="hat">
            <div className="nav-left">
                <Link to="/" className="site-name">
                    MY MARKET
                </Link>
            </div>
            <div className="nav-center">
                <Link to="/" className="cabinet-name">
                    SELLER'S CABINET
                </Link>
            </div>
            <div className="nav-right">
                <Link to="/product/new" className="rectangle-button">
                    <div className="rectangle"></div>
                    <div className="button">Add product</div>
                </Link>
                <Link to="/" className="rectangle-button">
                    <div className="rectangle"></div>
                    <div className="button" onClick={handleExitClick}>Exit</div>
                </Link>
            </div>
        </header>
    );
}

export default Hat;