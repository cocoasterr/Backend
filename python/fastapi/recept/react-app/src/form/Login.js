import axios from "axios";
import React, { useState } from "react";
import { Link, useNavigate } from "react-router-dom";
import { Toast, toast } from "react-toastify";

export default function Login(props) {
  // const navigate = useNavigate();
  const [loginForm, setLoginForm] = useState({
    email: "",
    password: "",
  })
  const onChangeForm = (label, event) => {
    switch (label) {
      case "email":
        setLoginForm({ ...loginForm, email: event.target.value })
        break;
      case "password":
        setLoginForm({ ...loginForm, password: event.target.value })
        break;
    }
  };
  // console.log(loginForm)

  const onSubmitHandler = async (event) => {
    event.preventDefault()
    // call api auth
    await axios.post("http://172.21.255.76:8888/api/user/login", loginForm).then((Response) => {
      // console.log(Response.data.access_token)
      localStorage.setItem("auth_token", Response.data.access_token)
      localStorage.setItem("auth_token_type", "Barier")

      //add successfull notif
      toast.success(Response.data.status)
      // props.setPage("home")
      setTimeout(() => {
        window.location.href = "/home"
      }, 1000)
    }).catch((Error) => {
      toast.error(Error.response.data.detail)
    })

  };
  return (
    <React.Fragment>
      <div>
        <h1 className="text-3xl font-bold text-center mb-4 cursor-pointer">
          Welcome to Receipt.Ido
        </h1>
        <p className="w-80 text-center text-sm mb-8 font-semibold text-gray-700 tracking-wide cursor-pointer mx-auto">
          Please login to your account
        </p>
      </div>
      <form onSubmit={onSubmitHandler}>
        <div className="space-y-4">
          <input
            type="text"
            placeholder="Email"
            className="block text-sm py-3 px-4 rounded-lg w-full border outline-none focus:ring focus:outline-none focus:ring-yellow-400"
            onChange={(event) => {
              onChangeForm("email", event)
            }}
          ></input>
          <input
            type="password"
            placeholder="Password"
            className="block text-sm py-3 px-4 rounded-lg w-full border outline-none focus:ring focus:outline-none focus:ring-yellow-400"
            onChange={(event) => {
              onChangeForm("password", event)
            }}
          ></input>
        </div>
        <div className="text-center mt-6">
          <button
            type="submit"
            className="py-3 w-64 text-xl text-white bg-yellow-400 rounded-2xl hover:bg-yellow-300 active:bg-yellow-500 outline-none">
            Sign in
          </button>
          <p className="mt-4 text-sm">
            You dont have an account?{" "}
            <Link to="/register" onClick={() =>
              props.setPage("register")}>
              <span className="underline cursor-pointer">Register </span>
            </Link>
          </p>
        </div>
      </form>
    </React.Fragment>
  );
}
// export default Login;
