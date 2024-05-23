import axios from "axios";

const instance = axios.create({
	baseURL: "/",
});

instance.interceptors.request.use(function (config) {
	config.headers["Authorization"] = localStorage.getItem("token");
	return config;
});

export default instance;
