import { Card, Text } from '@mantine/core';
import SpotifySettings from '../components/SpotifySettings';
import './Settings.css';

function Settings() {


	return (
		<>
			<Card
				shadow="sm"
				padding="lg"
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
						Settings
					</Text>
				</Card.Section>
			</Card>
			<div className="settings-container">
				<SpotifySettings />
			</div>
		</>
	)
}


export default Settings;
