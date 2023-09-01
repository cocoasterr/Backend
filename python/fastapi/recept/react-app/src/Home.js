import React from 'react';
import { Toast, toast } from "react-toastify";
import axios from "axios";
import { Link } from "react-router-dom";

function Home(props) {
    React.useEffect(() => {
        const accessToken = localStorage.getItem('auth_token');
        if (!accessToken) {
            // Jika tidak ada akses token, arahkan ke halaman login
            setTimeout(() => {
                window.location.href = "/"
            }, 1)
        }
    });

    const openSwaggerDocs = async (event) => {
        event.preventDefault();
        try {
            const accessToken = localStorage.getItem("auth_token");

            await axios.get("http://172.21.255.76:8888/docs#", {
                headers: {
                    Authorization: `Bearer ${accessToken}`,
                },
            }).then((Response) => {
                toast.success(Response.data.detail);
                window.open("http://172.21.255.76:8888/docs#", "_blank"); // Membuka URL di tab baru
            }).catch((Error) => {
                toast.error(Error.response.data.detail);
            });
        } catch (error) {
            console.error("Error opening API docs:", error);
            toast.error("An error occurred while opening API docs.");
        }
    };
    const logout = async (event) => {
        event.preventDefault();
        try {
            const accessToken = localStorage.getItem("auth_token");

            await axios.get("http://172.21.255.76:8888/api/user/logout", {
                headers: {
                    Authorization: `Bearer ${accessToken}`,
                },
            });

            localStorage.removeItem("auth_token");
            localStorage.removeItem("auth_token_type");

            toast.success("Logout success!");
            //reload after success login
            setTimeout(() => {
                window.location.href = "/"
            }, 1000)
        } catch (error) {
            console.error("Error logging out:", error);
            toast.error("An error occurred while logging out.");
        }
    };

    // const changePassword = async (event) => {
    //     event.preventDefault();
    //     try {
    //         toast.success("success!");
    //         <Link to="/change-password" onClick={() =>
    //             props.setPage("change-password")}>
    //         </Link>
    //     } catch (error) {
    //         toast.error(Error.response.data.detail);
    //     }
    // };
    return (
        <div className="text-center">
            <h1 className="text-3xl font-bold mb-4">Welcome to API Documentation</h1>
            <p className="text-gray-700 mb-8">Explore and learn about the available API endpoints.</p>
            <div className="border border-gray-300 p-4 rounded-lg bg-white">
                <p className="text-gray-600">
                    Mohon maaf, menu beranda belum dapat ditampilkan. Namun, API telah saya siapkan beserta dengan dokumentasinya.
                </p>
            </div>
            <Link to="/change-password" onClick={() =>
                props.setPage("change-password")}>
                <button
                    type="button"
                    className="py-3 w-64 text-xl text-white bg-yellow-400 rounded-2xl hover:bg-yellow-300 active:bg-yellow-500 outline-none mt-8">
                    Change Password
                </button>
            </Link>
            <button
                type="button"
                className="py-3 w-64 text-xl text-white bg-yellow-400 rounded-2xl hover:bg-yellow-300 active:bg-yellow-500 outline-none mt-8"
                onClick={openSwaggerDocs}>
                API
            </button>
            <button
                type="button"
                className="py-3 w-64 text-xl text-white bg-yellow-400 rounded-2xl hover:bg-yellow-300 active:bg-yellow-500 outline-none mt-8"
                onClick={logout}>
                logout
            </button>
        </div>
    );
}

export default Home;
