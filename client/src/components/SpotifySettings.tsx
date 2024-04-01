import { Card, Text, Image } from '@mantine/core';
import spotifyLogo from '../../assets/images/Spotify-Logo-RGB-Green.png';
import SpotifyLogin from '../services/SpotifyLogin';

function SpotifySettings() {
	return (
		<>
			<Card
				className='settings-card'
				shadow="sm"
				padding="lg"
				radius="md"
				withBorder
			>
				<Card.Section>
					{/* SpotifyLogin */}
					<Image
						p={20}
						src={spotifyLogo}
						height={160}
						fit="contain"
						alt="Norway"
					/>
				</Card.Section>

				<Text size="sm" c="dimmed">
					Connect your Spotify account to Playlist Swap
				</Text>

				<SpotifyLogin />

			</Card>
		</>
	)
}

export default SpotifySettings;