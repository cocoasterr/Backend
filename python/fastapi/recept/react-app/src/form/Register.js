import React, { useState } from "react";
import ReactDatePicker from "react-datepicker";
import "react-datepicker/dist/react-datepicker.module.css";
import { Link } from "react-router-dom";
import axios from "axios";
import { toast } from "react-toastify";
import { withRouter } from "react-router-dom";


function Register(props) {
  const options = [
    { value: "", label: "Select your Gender!" },
    { value: "MALE", label: "Male" },
    { value: "FEMALE", label: "Female" },
  ];
  //register form
  const [formRegister, setFormRegister] = useState({
    username: "",
    email: "",
    password: "",
    passwordConfirm: "",
    name: "",
    sex: "",
    birth: "",
    // profile: "",
    phone_number: "",
  });
  // default value date picker
  const [birthDate, setBirthDate] = useState(null);

  const onChangeForm = (label, event) => {
    switch (label) {
      case "email":
        setFormRegister({ ...formRegister, email: event.target.value })
        break;
      case "username":
        setFormRegister({ ...formRegister, username: event.target.value })
        break;
      case "password":
        setFormRegister({ ...formRegister, password: event.target.value })
        break;
      case "passwordConfirm":
        setFormRegister({ ...formRegister, passwordConfirm: event.target.value })
        break;
      case "name":
        setFormRegister({ ...formRegister, name: event.target.value })
        break;
      case "sex":
        setFormRegister({ ...formRegister, sex: event.target.value })
        break;
      case "birth":
        setFormRegister({ ...formRegister, birth: event.target.value })
        break;
      case "profile":
        setFormRegister({ ...formRegister, birth: event.target.value })
        break;
      case "phone_number":
        setFormRegister({ ...formRegister, birth: event.target.value })
        break;
    }
  };
  // console.log(loginForm)
  const onSubmitHandler = async (event) => {
    event.preventDefault()
    // call api auth
    await axios.post("http://172.21.255.76:8888/api/user/register", formRegister).then((Response) => {
      // console.log(Response.data.access_token)
      // localStorage.setItem("auth_token", Response.data.access_token)
      // localStorage.setItem("auth_token_type", "Barier")

      //add successfull notif
      console.log(Response)
      toast.success(Response.data.status)
      //reload after success login

      setTimeout(() => {
        window.location.href = "/"
      }, 1000)
    }).catch((Error) => {
      toast.error(Error.response.data.detail)
    })
  }
  return (
    <React.Fragment>
      <div>
        <h1 className="text-3xl font-bold text-center mb-4 cursor-pointer">
          Register
        </h1>
        <p className="w-80 text-center text-sm mb-8 font-semibold text-gray-700 tracking-wide cursor-pointer mx-auto">
          Insert your Biodata
        </p>
      </div>
      <form onSubmit={onSubmitHandler}>
        <div className="space-y-4">
          <input
            type="text"
            placeholder="Username"
            className="block text-sm py-3 px-4 rounded-lg w-full border outline-none focus:ring focus:outline-none focus:ring-yellow-400"
            onChange={(event) => {
              onChangeForm("username", event)
            }}
          ></input>
          <ReactDatePicker
            className="block text-sm py-3 px-4 rounded-lg w-full border outline-none focus:ring focus:outline-none focus:ring-yellow-400"
            dateFormat={"dd-MM-yyyy"}
            placeholderText="Birth date"
            selected={birthDate}
            onChange={(date) => setBirthDate(date)}
          // onChange={(event) => {
          //   onChangeForm("birth", event)
          // }}
          />
          <select
            value={formRegister.sex}
            onChange={(e) =>
              setFormRegister({
                ...formRegister,
                sex: e.target.value,
              })
            }
            className="block text-sm py-3 px-4 rounded-lg w-full border outline-none focus:ring focus:outline-none focus:ring-yellow-400"
          >
            {options.map((data) => (
              <option key={data.label} value={data.value} disabled={data.value === ""}>
                {data.label}
              </option>
            ))}
          </select>
          <input
            type="text"
            placeholder="Name"
            className="block text-sm py-3 px-4 rounded-lg w-full border outline-none focus:ring focus:outline-none focus:ring-yellow-400"
            onChange={(event) => {
              onChangeForm("name", event)
            }}
          ></input>
          <input
            type="number"
            placeholder="Phone number"
            className="block text-sm py-3 px-4 rounded-lg w-full border outline-none focus:ring focus:outline-none focus:ring-yellow-400"
            onChange={(event) => {
              onChangeForm("phone_number", event)
            }}
          ></input>
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
          <input
            type="password"
            placeholder="Confirm Password"
            className="block text-sm py-3 px-4 rounded-lg w-full border outline-none focus:ring focus:outline-none focus:ring-yellow-400"
            onChange={(event) => {
              onChangeForm("passwordConfirm", event)
            }}
          ></input>
        </div>
        <div className="text-center mt-6">
          <button
            type="submit"
            className="py-3 w-64 text-xl text-white bg-yellow-400 rounded-2xl hover:bg-yellow-300 active:bg-yellow-500 outline-none"
          >
            Register
          </button>
          <p className="mt-4 text-sm">
            Already have account?{" "}
            <Link to="/?login" onClick={() =>
              props.setPage("login")}>
              <span className="underline cursor-pointer">Sign in</span>
            </Link>
          </p>
        </div>
      </form>
    </React.Fragment>
  );
}

export default Register;
