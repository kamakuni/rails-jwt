import { useEffect } from 'react'
import { Outlet } from 'react-router'
import Message from '../../components/Message/Message'
import { useAppDispatch, useAppSelector } from '../../store'
import { isAlive } from './authSlice'

// https://v5.reactrouter.com/web/example/auth-workflow
const ProtectedRoutes = () => {
    const { isAuthenticated } = useAppSelector((state) => state.auth)
    const dispatch = useAppDispatch();

    useEffect(() => {
        if (!isAuthenticated) {
            dispatch(isAlive())
        }
    }, []);

    return isAuthenticated ? <Outlet /> : <Message body="Users is unauthorized." level="info"></Message>
}

export default ProtectedRoutes;