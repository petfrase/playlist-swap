import { useEffect, useState } from "react";
import { Button } from "@mantine/core";
import { useSelector, useDispatch } from "react-redux";
import { RootState, AppDispatch } from "../redux/stores/Store";
import { setSpotifyId } from "../redux/slices/UserSlice";
import { API } from "../utils/Axios";


// PKCE flow for spotify login
// function generateRandomString(length: number) {
// 	const possible = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789';
// 	const values = crypto.getRandomValues(new Uint8Array(length));
// 	return values.reduce((acc, x) => acc + possible[x % possible.length], "");
// }
// async function sha256(plain: string) {
// 	const encoder = new TextEncoder()
// 	const data = encoder.encode(plain)
// 	return window.crypto.subtle.digest('SHA-256', data)
// }

// function base64encode(input: ArrayBuffer) {
// 	return btoa(String.fromCharCode(...new Uint8Array(input)))
// 		.replace(/=/g, '')
// 		.replace(/\+/g, '-')
// 		.replace(/\//g, '_');
// }


function SpotifyLogin() {
	const [loginUrl, setLoginUrl] = useState('');
	const [loading, setLoading] = useState(true);
	const spotifyId = useSelector((state: RootState) => state.user.spotifyId);
	const dispatch = useDispatch<AppDispatch>();

	useEffect(() => {
		// PKCE flow below
		// const setupSpotifyLogin = async () => {
		// 	// Only setup the spotify login once and if the user is not logged in
		// 	if (spotifyId !== '') return;


		// 	const codeVerifier = generateRandomString(100);
		// 	const hashed = await sha256(codeVerifier)
		// 	const codeChallenge = base64encode(hashed);

		// 	const clientId = import.meta.env.VITE_SPOTIFY_CLIENT_ID;
		// 	const redirectUri = import.meta.env.VITE_SPOTIFY_REDIRECT_URI;

		// 	const scope = import.meta.env.VITE_SPOTIFY_SCOPES;
		// 	const authUrl = new URL(import.meta.env.VITE_SPOTIFY_AUTH_URL);

		// 	// generated in the previous step
		// 	window.localStorage.setItem('code_verifier', codeVerifier);

		// 	const params = {
		// 		response_type: 'code',
		// 		client_id: clientId,
		// 		scope,
		// 		code_challenge_method: 'S256',
		// 		code_challenge: codeChallenge,
		// 		redirect_uri: redirectUri,
		// 	}

		// 	authUrl.search = new URLSearchParams(params).toString();
		// 	setLoginUrl(authUrl.toString());
		// };

		const setupSpotifyLogin = () => {
			// Only setup the spotify login once and if the user is not logged in
			if (spotifyId !== '') return;
			const url = import.meta.env.VITE_SERVER_SPOTIFY_API_URL;
			setLoginUrl(url);
		}

		setupSpotifyLogin();
	}, []);

	// we should redrect the window to the loginUrl
	const loginSpotify = () => {
		try {
			setLoading(true);
			window.location.href = loginUrl;
		} catch (error) {
			console.error('Error logging in with Spotify', error);
		} finally {
			setLoading(false);
		}

	}


	// PKCE flow below
	// const openUrl = () => {
	// 	console.log('Opening URL', loginUrl);
	// 	//dispatch(setSpotifyId('asdf'));
	// 	window.open(loginUrl, '_self');
	// }


	return (
		<Button
			color="blue"
			fullWidth
			onClick={loginSpotify}
			disabled={spotifyId !== ''}
			mt="md"
			radius="md"
		>
			Sign in with Spotify
		</Button>
	)
}

export default SpotifyLogin;
