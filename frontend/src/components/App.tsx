import React from "react";
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import Home from "../pages/Home";
import Cadastro from "../pages/Cadastro";
import Login from "../pages/Login";
import BoasVindas from "../pages/BoasVindas";
import AuthGuard from "../middleware/AuthGuard";
import { GuestGuard } from "../middleware/GuestGuard";

const App: React.FC = () => {
    return (
        <Router>
            <Routes>
                <Route path="/" element={<Home />} />
                <Route path="/cadastro" element={<GuestGuard><Cadastro /></GuestGuard>} />
                <Route path="/login" element={<GuestGuard><Login /></GuestGuard>} />
                <Route path="/boasvindas" element={<AuthGuard><BoasVindas /></AuthGuard>} />
            </Routes>
        </Router>
    );
};

export default App;