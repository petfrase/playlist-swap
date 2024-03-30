import { SET_SPOTIFY_ID, UserActionTypes } from '../../models';

export const setSpotifyId = (spotifyId: string): UserActionTypes => ({
	type: SET_SPOTIFY_ID,
	payload: spotifyId,
});