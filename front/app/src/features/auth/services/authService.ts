import api from "../../../api";
import { LoginUser, NewUser, ResponseLogin, ResponseLogout, ResponseSignup } from "../models/AuthModels"

const register = async (user: NewUser): Promise<ResponseSignup> => {
    const response = await api.post('/api/v1/signup', user, {
        withCredentials: false
    })
    return response.data
}

const login = async (user: LoginUser): Promise<ResponseLogin> => {
    const response = await api.post('/api/v1/login', user)
    return response.data
}
const isAlive = async (): Promise<ResponseLogin> => {
    const response = await api.get('/api/v1/login')
    return response.data
}

const logout = async (): Promise<ResponseLogout> => {
    const response = await api.delete('/api/v1/logout')
    return response.data
}

const authService = {
    register,
    login,
    isAlive,
    logout,
}

export default authService