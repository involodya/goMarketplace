import {BrowserRouter as Router, Route, Routes} from "react-router-dom";
import './App.css';
import MainPage from './MainPage';
import AuthPage from './AuthPage';
import ProductPage from "./ProductPage";

function App() {
    const isAuth = Boolean(window.localStorage.getItem("access"));

    return (
        <Router>
            <Routes>
                {!isAuth && <Route path="/*" element={<AuthPage/>}/>}
                {isAuth && <Route path="/" element={<MainPage/>}/>}
                {isAuth && <Route path="/product/new" element={<ProductPage/>}/>}
                {isAuth && <Route path="/product/:id" element={<ProductPage/>}/>}
            </Routes>
        </Router>
    )
}

export default App;