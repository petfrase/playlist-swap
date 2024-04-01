import { Button, Card, Text } from "@mantine/core";
import { useEffect, useState } from "react";
import { useParams, useSearchParams } from "react-router-dom";
import { API } from "../utils/Axios";

// This component is used to handle when a user logs into an external service (spotify)
// and is redirected back to the application.
// Route: /settings/:type/redirect

// First get the type of service that the user is trying to login to.
// Then based on the service, that will determine how to get the access token
// e.g. for spotify, the auth token is in the query params of the url
export default function ServiceRedirectHandler() {
	const [loading, setLoading] = useState(true);

	// Get the type of service that the user is trying to login to
	const params = useParams();
	const [searchParams] = useSearchParams();

	// Send the auth token to the server
	// The server will then use the token to get the access token
	const handleSpotifyAuthToken = async (authToken: string, state: string) => {
		// Send the auth token to the server
		try {
			const query = `?code=${authToken}&state=${state}`;
			const response = await API.get(`/spotify/callback${query}`);
			if (response.status === 200) {
				// Set the user's id in the redux store
				// dispatch(setSpotifyId(response.data.spotifyId));
			}
		} catch (error) {
			console.error('Error handling spotify auth token', error);
		}
	};

	useEffect(() => {
		const init = async () => {
			if (params.type === 'spotify') {
				const token = searchParams.get('code');
				const state = searchParams.get('state');
				if (token && state) {
					await handleSpotifyAuthToken(token, state);
					// After async operation, set loading to false
					setLoading(false);
				} else {
					// Handle case where token is not available
					console.error("Token not found");
					setLoading(false);
				}
			}
		};

		init();
	}, [params.type, searchParams]);

	const getPlaylists = async () => {
		try {
			const response = await API.get('/spotify/playlists?limit=10&offset=0');
			if (response.status === 200) {
				// Set the user's playlists in the redux store
				// dispatch(setPlaylists(response.data.playlists));
			}
		} catch (error) {
			console.error('Error fetching playlists', error);
		}
	}


	if (loading) {
		return <div>Loading...</div>;
	}

	return (
		<Card shadow="sm" padding="lg" radius="md" withBorder>
			<Card.Section>
				<Text ta="left" size="xl" py={8} px={12}>
					Thank you for logging in to your {params.type} account! Please wait while we redirect you...
					<Button onClick={getPlaylists}>Get Playlists</Button>
				</Text>
			</Card.Section>
			{/* Display content based on the fetched data */}
		</Card>
	);
}