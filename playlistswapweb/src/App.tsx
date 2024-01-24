import '@mantine/core/styles.css';
import { AppShell, MantineProvider, Burger } from '@mantine/core';
import { useDisclosure } from '@mantine/hooks';
import './App.css';

function App() {
	const [opened, { toggle }] = useDisclosure();

	return <MantineProvider>
		{
			<AppShell
				header={{ height: 60 }}
				navbar={{
					width: 300,
					breakpoint: 'sm',
					collapsed: { mobile: !opened },
				}}
				padding="md"
			>
				<AppShell.Header className="header-bar">
					<Burger
						opened={opened}
						onClick={toggle}
						hiddenFrom="sm"
						size="sm"
						color="white"
					/>
					<h1>PlayList Swap</h1>
				</AppShell.Header>

				<AppShell.Navbar p="md">
					<AppShell.Section>Home</AppShell.Section>
					<AppShell.Section>Profile</AppShell.Section>
					<AppShell.Section>Settings</AppShell.Section>
				</AppShell.Navbar>

				<AppShell.Main>Main</AppShell.Main>
			</AppShell>
		}
	</MantineProvider>
}

export default App
