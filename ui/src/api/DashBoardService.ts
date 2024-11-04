

import { API_URL } from '@/config';
import { post } from "./http"
import { ChartReport } from './models';


export interface ChartReportRequest {
    startDate: Date;
    endDate: Date;
}

class DashSessionService {
    public async getReport(data: ChartReportRequest) {
        const response = await post<ChartReport>(`${API_URL}/api/dashboard/`, data);
        return response;
    }


}

export default new DashSessionService();