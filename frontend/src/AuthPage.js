import React, {useState} from "react";
import {API_URL} from './config';

import './AuthPage.css';

function AuthPage() {
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    const [loginError, setLoginError] = useState(false);
    const [registerError, setRegisterError] = useState(false);

    const handleLoginResponse = (response) => {
        const authHeader = response.headers.get('Authorization');
        console.log(authHeader);
        console.log(response.headers);
        if (authHeader) {
            const token = authHeader.split(' ')[1];
            window.localStorage.setItem("access", token);
            window.location.reload();
        } else {
            console.error("Login failed: Authorization header not found");
        }
    };

    const handleLoginClick = async () => {
        try {
            const response = await fetch(`${API_URL}/login`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({ email, password })
            });
            const data = await response.json();
            if (response.ok) {
                handleLoginResponse(response);
                setLoginError(false);
            } else {
                console.error('Error logging in:', data);
                setLoginError(true);
                setTimeout(() => setLoginError(false), 2000);
            }
        } catch (error) {
            console.error('Error fetching login data:', error);
            setLoginError(true);
            setTimeout(() => setLoginError(false), 2000);
        }
    };

    const handleCreateClick = async () => {
        try {
            const response = await fetch(`${API_URL}/register`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({ email, password })
            });
            const data = await response.json();
            if (response.ok) {
                handleLoginResponse(response);
                setRegisterError(false);
            } else {
                console.error('Error creating account:', data);
                setRegisterError(true);
                setTimeout(() => setRegisterError(false), 2000);
            }
        } catch (error) {
            console.error('Error fetching create data:', error);
            setRegisterError(true);
            setTimeout(() => setRegisterError(false), 2000);
        }
    };

    return (
        <div className="auth-page">
            <div className="auth-container">
                <h2 className="auth-title">Auth account for Seller</h2>
                <div className="input-group">
                    <input
                        type="email"
                        className="auth-input"
                        placeholder="Enter e-mail"
                        value={email}
                        onChange={(e) => setEmail(e.target.value)}
                    />
                </div>
                <div className="input-group">
                    <input
                        type="password"
                        className="auth-input"
                        placeholder="Enter password"
                        value={password}
                        onChange={(e) => setPassword(e.target.value)}
                    />
                </div>
                <div className="button-group">
                    <button className={`auth-button ${registerError ? 'register-error' : ''}`}
                            onClick={handleCreateClick}>Create
                    </button>
                    <button className={`auth-button login-button ${loginError ? 'login-error' : ''}`}
                            onClick={handleLoginClick}>Login
                    </button>
                </div>
            </div>
        </div>
    );
}

export default AuthPage;
