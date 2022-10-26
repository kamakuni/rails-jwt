import { createAsyncThunk, createSlice } from "@reduxjs/toolkit"
import AsyncState from "../models/AsyncState"
import { List, NewCard, NewList } from "./models/DashboardModels"
import listService from "./services/listService"

interface DashboardState extends AsyncState {
    lists: List[]
}

const initialState: DashboardState = {
    lists: [],
    isLoading: false,
    isSuccess: false,
    isError: false,
}

export const createList = createAsyncThunk(
    'dashboard/createList',
    async (list: NewList, thunkAPI) => {
        try {
            return await listService.create(list);
        } catch (error) {
            thunkAPI.rejectWithValue(error)
        }
    }
)

export const getAllLists = createAsyncThunk(
    'dashboard/getAlllists',
    async (_, thunkAPI) => {
        try {
            return await listService.getAllLists();
        } catch (error) {
            thunkAPI.rejectWithValue(error)
        }
    }
)

export const removeList = createAsyncThunk(
    'dashboard/removeList',
    async (list_id: string, thunkAPI) => {
        try {
            return await listService.removeList(list_id)
        } catch (error) {
            thunkAPI.rejectWithValue(error)
        }
    }
)

export const getCards = createAsyncThunk(
    'dashboard/getCards',
    async (list_id: string, thunkAPI) => {
        try {
            return await listService.getCards(list_id)
        } catch (error) {
            thunkAPI.rejectWithValue(error)
        }
    }
)

export const createCard = createAsyncThunk(
    'dashboard/createCard',
    async (card: NewCard, thunkAPI) => {
        try {
            return await listService.createCard(card);
        } catch (error) {
            thunkAPI.rejectWithValue(error)
        }
    }
)

export const removeCard = createAsyncThunk(
    'dashboard/removeCard',
    async (params: any, thunkAPI) => {
        try {
            return await listService.removeCard(params.list_id, params.card_id)
        } catch (error) {
            thunkAPI.rejectWithValue(error)
        }
    }
)

const dashboardSlice = createSlice({
    name: "dashboard",
    initialState,
    reducers: {
        reset: (state: DashboardState) => {
            state.isLoading = false;
            state.isSuccess = false;
            state.isError = false;
        },
    },
    extraReducers: (builder) => {
        builder
            .addCase(createList.pending, (state) => {
                state.isLoading = true;
            })
            .addCase(createList.fulfilled, (state, action) => {
                const copyed = [...state.lists]
                if (action && action.payload) {
                    copyed.push({ id: action.payload.id, title: action.payload.title, cards: [] })
                }
                state.lists = copyed
                state.isLoading = false;
                state.isSuccess = true;
            })
            .addCase(createList.rejected, (state) => {
                state.isLoading = false;
                state.isError = true;
            })
        builder
            .addCase(getAllLists.pending, (state) => {
                state.isLoading = true;
            })
            .addCase(getAllLists.fulfilled, (state, action) => {
                if (action && action.payload) {
                    state.lists = action.payload.map((list: any) => { return { id: list.id, title: list.title, cards: [...list.cards] } })
                }
                state.isLoading = false;
                state.isSuccess = true;
            })
            .addCase(getAllLists.rejected, (state) => {
                state.isLoading = false;
                state.isError = true;
            })
        builder
            .addCase(removeList.pending, (state) => {
                state.isLoading = true
            })
            .addCase(removeList.fulfilled, (state, action) => {
                if (action && action.payload) {
                    const list_id = action.payload.list_id
                    const filtered = state.lists.filter((list) => { return list.id !== list_id })
                    state.lists = filtered
                }
                state.isLoading = false
                state.isSuccess = true
            })
            .addCase(removeList.rejected, (state) => {
                state.isError = true
                state.isLoading = false
            })
        builder
            .addCase(getCards.pending, (state) => {
                state.isLoading = true;
            })
            .addCase(getCards.fulfilled, (state, action) => {
                //state.lists = action.payload.map((card: any) => { return { id: ca.id, title: list.title, cards: [] } })
                state.isLoading = false;
                state.isSuccess = true;
            })
            .addCase(getCards.rejected, (state) => {
                state.isLoading = false;
                state.isError = true;
            })
        builder
            .addCase(createCard.pending, (state) => {
                state.isLoading = true;
            })
            .addCase(createCard.fulfilled, (state, action) => {
                const copyed = [...state.lists]
                if (action && action.payload) {
                    const list_id = action.payload.list_id
                    const index = state.lists.findIndex((list: any) => list.id === list_id)
                    copyed[index].cards.push({ id: action.payload.id, title: action.payload.title, body: action.payload.body })
                }
                state.lists = copyed
                state.isLoading = false;
                state.isSuccess = true;
            })
            .addCase(createCard.rejected, (state) => {
                state.isLoading = false;
                state.isError = true;
            })
        builder
            .addCase(removeCard.pending, (state) => {
                state.isLoading = true;
            })
            .addCase(removeCard.fulfilled, (state, action) => {
                //const copyed = [...state.lists]
                if (action && action.payload) {
                    const list_id = action.payload.list_id
                    const card_id = action.payload.card_id
                    const index = state.lists.findIndex((list: any) => list.id === list_id)
                    state.lists[index].cards = state.lists[index].cards.filter((card) => { return card.id !== card_id })
                }
                state.isLoading = false;
                state.isSuccess = true;
            })
            .addCase(removeCard.rejected, (state) => {
                state.isLoading = false;
                state.isError = true;
            })
    }
})

export const { reset } = dashboardSlice.actions
export default dashboardSlice.reducer;