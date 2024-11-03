import axios from 'axios';

import { API_URL } from '@/config';
import { LoginParams, LoginResponse, SignUpParams } from './models';


class AuthApiService {
    public async signup(body: SignUpParams) {
        const response = await axios.post<LoginResponse>(`${API_URL}/auth/signup`, body);
        return response.data;
    }
    public async login(body: LoginParams) {
        const response = await axios.post<LoginResponse>(`${API_URL}/auth/login`, body);
        return response.data;
    }
}

export default new AuthApiService();