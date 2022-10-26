import { useEffect } from 'react';
import { NavLink } from 'react-router-dom';
import { logout, reset } from '../../features/auth/authSlice'
import { useAppDispatch, useAppSelector } from '../../store';
import styles from '../../styles/components/Nav.module.scss';

function Nav() {

    const dispatch = useAppDispatch()

    const { isSuccess, isAuthenticated } = useAppSelector((state) => state.auth)

    const handleLogout = (e: React.MouseEvent<HTMLAnchorElement, MouseEvent>) => {
        e.preventDefault()
        dispatch(logout())
    }

    useEffect(() => {
        if (isSuccess) {
            dispatch(reset());
        }
    }, [])

    return (
        <nav>
            <NavLink to="/" className={styles.nav}>Home</NavLink>
            <NavLink to="/dashboard" className={styles.nav}>Dashboard</NavLink>
            {
                isAuthenticated ? <NavLink to="/logout" onClick={(e) => handleLogout(e)} className={styles.nav}>Logout</NavLink>
                    : <NavLink to="/login" className={styles.nav}>Login</NavLink>
            }
            <NavLink to="/signup" className={styles.nav}>Signup</NavLink>
        </nav>
    )
}

export default Nav;