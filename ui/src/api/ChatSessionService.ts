

import { API_URL } from '@/config';
import { SessionDto } from './models';
import { post, get } from "./http"


class ChatSessionService {
    public async createSession() {
        const response = await post<SessionDto>(`${API_URL}/api/chat/session`, null);
        return response;
    }
    public async getSessions() {
        const response = await get<SessionDto[]>(`${API_URL}/api/chat/session/list`);
        return response
    }

}

export default new ChatSessionService();