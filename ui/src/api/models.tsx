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


export interface Record {
    x: string
    y: number
}

export interface Chart {
    title: string
    records: Record[]
    type: string // pie, chart, bar
    description: string
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
