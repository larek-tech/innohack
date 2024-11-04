

import { API_URL } from '@/config';
import { SessionContentDto, SessionDto } from './models';
import { post, get } from "./http"


class ChatSessionService {
    public async createSession() {
        const response = await post<SessionDto>(`${API_URL}/api/session`, null);
        return response;
    }
    public async getSessions() {
        const response = await get<SessionDto[]>(`${API_URL}/api/session/list`);
        return response
    }

    public async getSessionContent(sessionId: string) {
        const response = await get<SessionContentDto[]>(`${API_URL}/api/session/${sessionId}`);
        return response
    }

}

export default new ChatSessionService();