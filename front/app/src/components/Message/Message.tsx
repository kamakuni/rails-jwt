import styles from '../../styles/components/Message.module.scss'

type Level = "error" | "warn" | "info" | "success"
interface MessageProps {
    body: string
    level: Level
}

const Message = (props: MessageProps) => {
    let className;
    switch (props.level) {
        case "error":
            className = styles.error;
            break;
        case "warn":
            className = styles.warn;
            break;
        case "info":
            className = styles.info;
            break;
        case "success":
            className = styles.success;
            break;
    }
    return (
        <div className={className}>
            <p>{props.body}</p>
        </div>
    )
}

export default Message;