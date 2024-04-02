import { Button, Card, Group, Image, Text } from "@mantine/core";
import { API } from "../utils/Axios";
import './Playlists.css';
import spotifyLogo from '../../assets/images/Spotify-Logo-RGB-Green.png';
import spotifyIcon from '../../assets/images/Spotify_Icon_RGB_Green.png';
import { FaUserCircle } from "react-icons/fa";
import { useEffect, useState } from "react";


interface IPlaylist {
	id: string;
	name: string;
	description: string;
	imageUrl: string;
	owner: string;
}


export default function Playlists() {
	const [spotifyPlaylists, setSpotifyPlaylists] = useState<IPlaylist[]>([]);

	// load playlists from db
	useEffect(() => {
		const getPlaylists = async () => {
			try {
				const response = await API.get('/spotify/playlists?limit=10&offset=0');
				if (response.status === 200) {
					console.log(response);
					// Set the user's playlists in the redux store
					setSpotifyPlaylists(response.data);
				}
			} catch (error) {
				console.error('Error fetching playlists', error);
			}
		}

		getPlaylists();
	}, []);


	const getPlayImageUrl = (playlist: IPlaylist) => {
		if (playlist.imageUrl) {
			// check if the url is valid
			try {
				new URL(playlist.imageUrl);
				return playlist.imageUrl;
			} catch {
				return spotifyIcon;
			}
		} else {
			return spotifyIcon;
		}
	}


	return (
		<>
			<Card
				shadow="sm"
				radius="md"
				withBorder
			>
				<Card.Section>
					<Text
						ta="left"
						size="lg"
						py={8}
						px={12}
					>
						Playlists
					</Text>
				</Card.Section>
			</Card>

			<Card
				mt={12}
				shadow="sm"
				p={0}
				radius="md"
				withBorder
			>
				<Card.Section className="playlists-header">
					<Image
						src={spotifyLogo}
						style={{ width: 80 }}
						fit="contain"
					/>
					<Text
						ta="left"
						size="lg"
						py={8}
						px={12}
					>
						Playlists
					</Text>

					<Button
						style={{ marginLeft: 'auto' }}
						variant="light"
						color="#1DB954"
						mt={12}
					>
						View All
					</Button>

				</Card.Section>

				<Card.Section
					className="playlists-container"
				>
					{spotifyPlaylists.map((playlist) => {
						return (
							<Card
								className="playlists-card spotify-card"
								shadow="sm"
								radius="md"
								withBorder
							>
								<Card.Section>
									<Image
										px={4}
										py={12}
										src={getPlayImageUrl(playlist)}
										height={150}
										fit="contain"
									/>
								</Card.Section>

								<Text
									className="spotify-text"
									ta="left"
									size="sm"
									pt={8}
									pb={4}
								>
										{playlist.name}
								</Text>

								<Group className="playlists-owner">
									<FaUserCircle color="white"/>
									<Text
										ta="left"
										size="xs"
									>
										{playlist.owner}
									</Text>
								</Group>

								<Button
									className="spotify-button"
									variant="light"
									color="#1DB954"
									mt={12}
								>
									View Playlist
								</Button>
							</Card>
						)
					}
				)}
				</Card.Section>
			</Card>
		</>
	)
}