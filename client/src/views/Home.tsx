import { Card, Text } from '@mantine/core';

export default function Home() {
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
						size="xl"
						py={8}
						px={12}
					>
						Home
					</Text>
				</Card.Section>
			</Card>
		</>
	)
}