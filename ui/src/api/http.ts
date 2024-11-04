import axios from 'axios';
import { API_URL } from '@/config';
import { LOCAL_STORAGE_KEY } from '@/auth/AuthProvider';

const axiosInstance = axios.create({
    baseURL: API_URL,
    headers: {
        'Content-Type': 'application/json',
    },
});

axiosInstance.interceptors.request.use(
    (config) => {
        const token = JSON.parse(localStorage.getItem('authUser') as string)?.user?.token;

        if (token) {
            config.headers['Authorization'] = `Bearer ${token}`;
        }

        config.headers['ngrok-skip-browser-warning'] = '*';

        return config;
    },
    (error) => {
        return Promise.reject(error);
    }
);

axiosInstance.interceptors.response.use(
    (response) => {
        return response;
    },
    (error) => {
        if (error.response && error.response.status === 401) {
            localStorage.removeItem(LOCAL_STORAGE_KEY);

            window.location.href = '/login';
        }
        return Promise.reject(error);
    }
);

export default axiosInstance;


export const get = async <T>(url: string, config = {}) => {
    const response = await axiosInstance.get<T>(url, config);
    return response.data;
};

export const post = async <T>(url: string, data: unknown, config = {}) => {
    const response = await axiosInstance.post<T>(url, data, config);
    return response.data;
};

export const put = async <T>(url: string, data: unknown, config = {}) => {
    const response = await axiosInstance.put<T>(url, data, config);
    return response.data;
};

export const del = async <T>(url: string, config = {}) => {
    const response = await axiosInstance.delete<T>(url, config);
    return response.data;
};