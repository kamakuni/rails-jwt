interface LoginUser {
    email: string,
    password: string,
}

interface NewUser {
    email: string,
    password: string,
}

interface ResponseSignup {
    name: string,
    email: string,
    message?: string,
}

interface ResponseLogin {
    message: string,
}

interface ResponseLogout {
    message: string,
}

export type { LoginUser, NewUser, ResponseSignup, ResponseLogin, ResponseLogout };