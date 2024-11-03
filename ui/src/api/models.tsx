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