import { Link } from "react-router-dom";
import React, { useState } from "react";
import axios from "axios";
import { Toast, toast } from "react-toastify";

function ChangePassword(props) {
    const [changePasswordForm, setChangePasswordForm] = useState({
        email: "",
        current_password: "",
        password: "",
        passwordConfirm: "",
    })
    const onChangeForm = (label, event) => {
        switch (label) {
            case "email":
                setChangePasswordForm({ ...changePasswordForm, email: event.target.value })
                break;
            case "current_password":
                setChangePasswordForm({ ...changePasswordForm, current_password: event.target.value })
                break;
            case "password":
                setChangePasswordForm({ ...changePasswordForm, password: event.target.value })
                break;
            case "passwordConfirm":
                setChangePasswordForm({ ...changePasswordForm, passwordConfirm: event.target.value })
                break;
        }
    };
    const onSubmitHandler = async (event) => {
        event.preventDefault();
        const accessToken = localStorage.getItem("auth_token");

        try {
            const response = await axios.put(
                "http://172.21.255.76:8888/api/user/change-password",
                changePasswordForm,
                {
                    headers: {
                        Authorization: `Bearer ${accessToken}`,
                    },
                }
            );

            // add success notification
            toast.success("Success!");
            setTimeout(() => {
                window.location.href = "/home"
            }, 1000)
        } catch (error) {
            console.error("Error changing password:", error);
            toast.error(error.response?.data.detail || "An error occurred while changing password.");
        }
    };

    return (
        <React.Fragment>
            <div>
                <h1 className="text-3xl font-bold text-center mb-4 cursor-pointer">
                    Change password
                </h1>
                <p className="w-80 text-center text-sm mb-8 font-semibold text-gray-700 tracking-wide cursor-pointer mx-auto">
                    Update your password!
                </p>
            </div>
            <form onSubmit={onSubmitHandler}>
                <div className="space-y-4">
                    <input
                        type="email"
                        placeholder="Email"
                        className="block text-sm py-3 px-4 rounded-lg w-full border outline-none focus:ring focus:outline-none focus:ring-yellow-400"
                        onChange={(event) => {
                            onChangeForm("email", event)
                        }}
                    ></input>
                    <input
                        type="password"
                        placeholder="Current Password"
                        className="block text-sm py-3 px-4 rounded-lg w-full border outline-none focus:ring focus:outline-none focus:ring-yellow-400"
                        onChange={(event) => {
                            onChangeForm("current_password", event)
                        }}
                    ></input>
                    <input
                        type="password"
                        placeholder="New Password"
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
                        Update Password
                    </button>
                    <p className="mt-4 text-sm">
                        <Link to="/home" onClick={() =>
                            props.setPage("home")}>
                            <span className="underline cursor-pointer">Back to home</span>
                        </Link>
                    </p>
                </div>
            </form>
        </React.Fragment >
    );
}

export default ChangePassword