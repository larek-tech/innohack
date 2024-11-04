export interface LoginResponse {
    token: string
    type: string
}

export interface LoginParams {
    email: string
    password: string
}

export interface SignUpParams {
    email: string
    password: string
}


export interface QueryDto {
    id: number
    prompt: string
    createdAt: Date
}

export interface SessionDto {
    id: number
    title: string
    createdAt: Date | null
}

export interface QueryDto {
    id: number
    prompt: string
    created_at: Date
}

export interface SessionContentDto {
    query: QueryDto
    response: ResponseDto
}


export interface Multiplier {
    key: string
    value: number
}


export interface ResponseDto {
    queryId: number
    sources: string[]
    filenames: string[]
    charts: Chart[]
    description: string
    multipliers: Multiplier[]
    createdAt: Date
    error: string
    isLast: boolean
}

export interface Record {
    x: string;
    y: number;
}

export interface Chart {
    color: string;
    records: Record[];
    type: string; // pie, chart, bar
}

export interface Legend {
    [key: string]: string;
}

export interface Info {
    charts: Chart[];
    legend: Legend;
}

export interface ChartReport {
    summary: string;
    info: {
        [key: string]: Info;
    };
    multipliers: Multiplier[];
    startDate: number;
    endDate: number;
}

export interface Multiplier {
    key: string;
    value: number;
}