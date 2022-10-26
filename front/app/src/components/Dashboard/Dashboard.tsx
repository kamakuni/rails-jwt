import React, { useEffect, useState } from 'react'
import { createList, getAllLists, removeList, createCard, removeCard, reset } from '../../features/dashboard/dashboardSlice'
import { List } from '../../features/dashboard/models/DashboardModels'
import { useAppDispatch, useAppSelector } from '../../store'
import styles from '../../styles/components/Dashoboard.module.scss'

interface CardProps {
    body: string
    title: string
}

const MemorizedCard = React.memo(function Card(props: CardProps) {

    return (
        <div>
            <div className={styles.control}>
                <label className={styles.label}>title:</label>
                {props.title}
            </div>
            <div>
                <label className={styles.label}>body:</label>
                {props.body}
            </div>
        </div>
    )
})

interface ListItemProps {
    list: List,
}

function ListItem(props: ListItemProps) {

    const [card, setCard] = useState({ title: "", body: "" })

    const dispatch = useAppDispatch();

    const handleCardTitleChange = (value: string) => {
        const copyed = { ...card }
        copyed.title = value
        setCard(copyed)
    }

    const handleCardBodyChange = (value: string) => {
        const copyed = { ...card }
        copyed.body = value
        setCard(copyed)
    }

    const handleAddCards = () => {
        dispatch(createCard({ title: card.title, body: card.body, list_id: props.list.id }))
    }

    const handleRemoveCards = (index: number) => {
        dispatch(removeCard({ list_id: props.list.id, card_id: props.list.cards[index].id }))
    }

    const handleRemoveList = () => {
        dispatch(removeList(props.list.id))
    }

    return (
        <div>
            <div className={styles.control}>
                <label className={styles.label}>{props.list.title}</label>
            </div>
            <div className={styles.control}>
                <label className={styles.label}>title</label>
            </div>
            <div className={styles.control}>
                <input className={styles.input} onChange={(e) => handleCardTitleChange(e.target.value)}></input>
            </div>
            <div className={styles.control}>
                <label className={styles.label}>body</label>
            </div>
            <div className={styles.control}>
                <input className={styles.input} onChange={(e) => handleCardBodyChange(e.target.value)}></input>
            </div>
            <div className={styles.control}>
                <button className={styles.button} onClick={handleAddCards} >Add Cards</button>
            </div>
            <div>
                <ul>
                    {props.list.cards.map((card, index) => {
                        return (
                            <li key={index}>
                                <MemorizedCard
                                    body={card.body}
                                    title={card.title} />
                                <div className={styles.control}>
                                    <button className={styles.button} onClick={() => handleRemoveCards(index)} >Remove Cards</button>
                                </div>
                            </li>
                        )
                    })}
                </ul>
            </div>
            <div className={styles.control}>
                <button className={styles.button} onClick={handleRemoveList}>Remove List</button>
            </div>
        </div>
    )
}

function Dashboard() {

    const [title, setTitle] = useState("")

    const { isLoading, isSuccess, lists } = useAppSelector((state) => state.dashboard)
    const dispatch = useAppDispatch();

    const handleTitleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        setTitle(e.target.value)
    }

    const handleAddListClick = (e: React.MouseEvent<HTMLButtonElement, MouseEvent>) => {
        dispatch(createList({ title: title }))
        setTitle("")
    }

    useEffect(() => {
        dispatch(getAllLists())
        if (isSuccess) {
            dispatch(reset())
        }
    }, [])

    /*    if (isLoading) {
            return (
                <div>
                    <p>Loading...</p>
                </div>
            );
        }*/

    return (
        <div>
            <h2>Dashboard</h2>
            <div>
                <div className={styles.control}>
                    <label className={styles.label}>title</label>
                </div>
                <div className={styles.control}>
                    <input className={styles.input} onChange={handleTitleChange}></input>
                </div>
                <div className={styles.control}>
                    <button className={styles.button} onClick={handleAddListClick}>Add Lists</button>
                </div>
            </div>
            <h3>List</h3>
            <div>
                <ul>
                    {lists.map((list, index) => {
                        return <li key={index}>
                            <ListItem list={list} ></ListItem>
                        </li>
                    })}
                </ul>
            </div>
        </div>);
}

export default Dashboard;