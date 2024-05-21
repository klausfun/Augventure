import axios from "axios";

const instance = axios.create({
	baseURL: "/api/",
	// baseURL: "8000/api/",
});

instance.interceptors.request.use(function (config) {
	config.headers["Authorization"] = localStorage.getItem("token");
	return config;
});

export default instance;
