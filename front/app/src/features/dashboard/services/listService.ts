import api from "../../../api";
import { NewCard, NewList, ResponseCard, ResponseList, ResponseRemoveList, ResponseRemoveCard } from "../models/DashboardModels";

const create = async (list: NewList): Promise<ResponseList> => {
    const response = await api.post('/api/v1/lists', list);
    return response.data;
}

const createCard = async (card: NewCard): Promise<ResponseCard> => {
    const response = await api.post('/api/v1/cards', card);
    return response.data;
}

const getAllLists = async (): Promise<ResponseList[]> => {
    const listsResponse = await api.get('/api/v1/lists')
    const cardsResponse = await Promise.all(listsResponse.data.map(async (list: any) => {
        return getCards(list.id)
    }))

    const lists = listsResponse.data.map((list: any) => ({ ...list, cards: cardsResponse.flat().filter(card => card.list_id == list.id) }))
    return lists;
}

const removeList = async (list_id: string): Promise<ResponseRemoveList> => {
    await api.delete(`/api/v1/lists/${list_id}`)
    return { "list_id": list_id }
}

const updateList = async (list_id: string): Promise<any> => {
    const response = await api.delete(`/api/v1/lists/${list_id}`)
    return response
}

const getCards = async (list_id: string): Promise<ResponseCard[]> => {
    const response = await api.get(`/api/v1/cards?list_id=${list_id}`)
    return response.data;
}

const updateCard = async (card_id: string): Promise<any> => {
    const response = await api.put('/api/v1/cards')
    return response.data;
}

const removeCard = async (list_id: string, card_id: string): Promise<ResponseRemoveCard> => {
    await api.delete(`/api/v1/cards/${card_id}`)
    return { "list_id": list_id, "card_id": card_id };
}

const listService = {
    create,
    createCard,
    getAllLists,
    removeList,
    updateList,
    getCards,
    updateCard,
    removeCard,
}

export default listService;