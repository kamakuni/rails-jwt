import { BrowserRouter, Route, Routes } from 'react-router-dom';
import Dashboard from './components/Dashboard/Dashboard';
import Home from './components/Home/Home';
import Login from './components/Login/Login';
import SignUp from './components/SignUp/SignUp';
import ProtectedRoutes from './features/auth/ProtectedRoutes';
import Header from './components/Header/Header';
import styles from './styles/App.module.scss'

function App() {

  return (
    <div className="wrapper">
      <BrowserRouter>
        <Header></Header>
        <main className={styles.container}>
          <Routes>
            <Route path="/" element={<Home />} />
            <Route path="/signup" element={<SignUp />} />
            <Route path="/login" element={<Login />} />
            <Route element={<ProtectedRoutes />} >
              <Route path="/dashboard" element={<Dashboard />} />
            </Route>
          </Routes>
        </main>
      </BrowserRouter>
    </div >
  );

}

export default App;
