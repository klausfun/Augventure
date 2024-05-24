import axios from "axios";

// Базовый URL вашего бэкенда
const API_BASE_URL = "http://localhost:8000";

const loginRest = async (username, secret) => {
    try {
        const response = await axios.post(`${API_BASE_URL}/auth/sign-in`, {
            username,
            secret
        });
        return response.data;
    } catch (error) {
        console.error("Login failed:", error);
        throw error;
    }
};

const signupRest = async (username, secret, email, first_name, last_name) => {
    try {
        const response = await axios.post(`${API_BASE_URL}/auth/sign-up`, {
            username,
            secret,
            email,
            first_name,
            last_name
        });
        return response.data;
    } catch (error) {
        console.error("Registration failed:", error);
        throw error;
    }
};

export { loginRest, signupRest };
