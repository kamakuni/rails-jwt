import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import AsyncState from "../models/AsyncState";
import { LoginUser, NewUser } from "./models/AuthModels"
import authService from "./services/authService";

interface AuthState extends AsyncState {
    isAuthenticated: boolean,
}

const initialState: AuthState = {
    isLoading: false,
    isSuccess: false,
    isError: false,
    isAuthenticated: false,
}

export const register = createAsyncThunk(
    'auth/register',
    async (user: NewUser, thunkAPI) => {
        try {
            return await authService.register(user);
        } catch (error) {
            return thunkAPI.rejectWithValue("Unabled to register!")
        }
    }
)

export const login = createAsyncThunk(
    'auth/login',
    async (user: LoginUser, thunkAPI) => {
        try {
            return await authService.login(user);
        } catch (error) {
            return thunkAPI.rejectWithValue("Unabled to login!")
        }
    }
)

export const isAlive = createAsyncThunk(
    'auth/isAlive',
    async (_, thunkAPI) => {
        try {
            return await authService.isAlive();
        } catch (error) {
            return thunkAPI.rejectWithValue("Unabled to login!")
        }
    }
)

export const logout = createAsyncThunk(
    'auth/logout',
    async (_, thunkAPI) => {
        try {
            return await authService.logout();
        } catch (error) {
            return thunkAPI.rejectWithValue("Unabled to login!")
        }
    }
)

export const authSlice = createSlice({
    name: 'auth',
    initialState,
    reducers: {
        reset: (state: AuthState) => {
            state.isLoading = false;
            state.isSuccess = false;
            state.isError = false;
        },
    },
    extraReducers: (builder) => {
        builder
            .addCase(register.pending, (state) => {
                state.isLoading = true;
            })
            .addCase(register.fulfilled, (state) => {
                state.isLoading = false;
                state.isSuccess = true;
            })
            .addCase(register.rejected, (state) => {
                state.isLoading = false;
                state.isError = true;
            })
        builder
            .addCase(login.pending, (state) => {
                state.isLoading = true;
            })
            .addCase(login.fulfilled, (state) => {
                state.isLoading = false;
                state.isSuccess = true;
                state.isAuthenticated = true;
            })
            .addCase(login.rejected, (state) => {
                state.isLoading = false;
                state.isError = true;
            })
        builder
            .addCase(isAlive.pending, (state) => {
                state.isLoading = true;
            })
            .addCase(isAlive.fulfilled, (state) => {
                state.isLoading = false;
                state.isSuccess = true;
                state.isAuthenticated = true;
            })
            .addCase(isAlive.rejected, (state) => {
                state.isLoading = false;
                state.isError = true;
            })
        builder
            .addCase(logout.pending, (state) => {
                state.isLoading = true;
            })
            .addCase(logout.fulfilled, (state) => {
                state.isLoading = false;
                state.isSuccess = true;
                state.isAuthenticated = false;
            })
            .addCase(logout.rejected, (state) => {
                state.isLoading = false;
                state.isError = true;
            })
    }
})

export const { reset } = authSlice.actions
export default authSlice.reducer;