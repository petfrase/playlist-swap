// State interface
export interface IUserState {
	userName: string;
	spotifyId: string;
}

// Action types
export const SET_SPOTIFY_ID = 'SET_SPOTIFY_ID';

interface ISetSpotifyIdAction {
	type: typeof SET_SPOTIFY_ID;
	payload: string;
}

// Union action type
export type UserActionTypes = ISetSpotifyIdAction;
