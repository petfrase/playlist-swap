import axios from "axios";

export const API = axios.create({
	baseURL: import.meta.env.VITE_SERVER_API_URL,
	timeout: 5000, // throw error if no response after 5 seconds
	withCredentials: true, // send cookies with requests
	validateStatus: (status) => status >= 200 && status < 300 // throw error if non-2xx response
})

// using axios interceptors to handle errors
// return the message property of the error response
// if it exists, otherwise return the error message
API.interceptors.response.use(
	(response) => response,
	(error) => {
		if (error.response?.data?.message) {
			return Promise.reject(error.response.data.message);
		}
		return Promise.reject(error.message);
	}
)

// export const setAuthToken = (token: string) => {
// 	API.defaults.headers.common['Authorization'] = `Bearer ${token}`;
// }