import logo from "./logo.svg";
import "./App.css";
import Login from "./form/Login";
import Register from "./form/Register";
import { useState } from "react";
import ChangePassword from "./form/ChangePassword";
import Home from "./Home";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";

function App(props) {
  const [page, setPage] = useState("login");
  const accessToken = localStorage.getItem("auth_token");

  if (!accessToken && page !== "login") {
    setPage("login");
  }

  return (
    <div className="min-h-screen bg-yellow-400 flex justify-center items-center">
      <div className="py-12 px-12 bg-white rounded-2xl shadow-x1 z-20">
        <Routes>
          <Route path="/register" element={<Register setPage={setPage} />} />
          <Route path="/change-password" element={<ChangePassword setPage={setPage} />} />
          <Route path="/home" element={<Home setPage={setPage} />} />
          {/* Default route */}
          <Route path="/" element={<Login setPage={setPage} />} />
        </Routes>
      </div>
    </div>
  );
}

export default App;
