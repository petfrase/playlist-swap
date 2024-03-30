import { PayloadAction, createSlice } from "@reduxjs/toolkit";
import { IUserState } from "../../models";

const initialState: IUserState = {
	userName: "",
	spotifyId: "",
};

const userSlice = createSlice({
	name: "userStore",
	initialState,
	reducers: {
		setSpotifyId: (state, action: PayloadAction<string>) => {
			state.spotifyId = action.payload;
		},
	},
});

export const {
	setSpotifyId
} = userSlice.actions;

export default userSlice.reducer;