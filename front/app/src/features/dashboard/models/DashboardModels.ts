interface Card {
    id: string,
    title: string,
    body: string,
}

interface List {
    id: string,
    title: string,
    cards: Card[],
}

interface NewCard {
    title: string,
    body: string,
    list_id: string,
}

interface NewList {
    title: string
}

interface ResponseList {
    id: string,
    title: string,
    user_id: string,
    created_at: Date,
    updated_at: Date,
    cards: ResponseCard[],
}

interface ResponseCard {
    id: string,
    title: string,
    body: string,
    list_id: string,
    created_at: Date,
    updated_at: Date,
}

interface ResponseRemoveList {
    list_id: string
}

interface ResponseRemoveCard {
    list_id: string,
    card_id: string,
}

export type { Card, List, NewCard, NewList, ResponseList, ResponseCard, ResponseRemoveList, ResponseRemoveCard };