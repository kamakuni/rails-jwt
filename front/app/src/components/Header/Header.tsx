import Nav from "../Nav/Nav"
import styles from "../../styles/components/Header.module.scss"

const Header = () => {
    return (
        <header className={styles.header}>
            <div className={styles.content}>
                <Nav></Nav>
            </div>
        </header>
    )
}

export default Header;