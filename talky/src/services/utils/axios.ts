import axios from "axios";

const baseURL = `${import.meta.env.VITE_API_LOCATION}/api`;

export const axiosInstance = axios.create({
  baseURL,
});
