import axios from 'axios'
import React, { useEffect } from 'react'
import { login, reset } from '../../features/auth/authSlice'
import { useAppDispatch, useAppSelector } from '../../store'
import Message from '../Message/Message'
import styles from '../../styles/components/Login.module.scss'

const Login: React.FC = () => {

    const [email, setEmail] = React.useState("")
    const [password, setPassword] = React.useState("")
    const dispatch = useAppDispatch()
    const { isLoading, isSuccess, isAuthenticated } = useAppSelector((state) => state.auth)

    const handleEmailChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        setEmail(e.target.value);
    }

    const handlePasswordChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        setPassword(e.target.value);
    }

    const handleClick = (e: React.MouseEvent<HTMLButtonElement, MouseEvent>) => {
        e.preventDefault()

        dispatch(login({ email: email, password: password }))

    }

    useEffect(() => {
        if (isSuccess) {
            dispatch(reset())
        }
    }, [isSuccess])

    if (isLoading) {
        return (
            <div>
                <p>Loading...</p>
            </div>
        );
    }

    return (
        <div>
            {isAuthenticated
                ? <Message body="You're already logged in." level="info"></Message>
                : <form>
                    <div className={styles.control}>
                        <div>
                            <label className={styles.label} >Email</label>
                        </div>
                        <div>
                            <input className={styles.input} type="text" placeholder="email" onChange={handleEmailChange} />
                        </div>
                    </div>
                    <div className={styles.control}>
                        <div>
                            <label className={styles.label} >Password</label>
                        </div>
                        <div>
                            <input className={styles.input} type="text" placeholder="password" onChange={handlePasswordChange} />
                        </div>
                    </div>
                    <div className={styles.control}>
                        <button className={styles.button} type="submit" onClick={handleClick} >Submit</button>
                    </div>
                    <div className={styles.control}>
                        <a className={styles.button} href="" >OAuth</a>
                    </div>
                </form>
            }
        </div>
    );

}

export default Login;