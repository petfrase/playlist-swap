import '@mantine/core/styles.css';
import { AppShell, MantineProvider, Burger } from '@mantine/core';
import { NavLink } from '@mantine/core';
import { useDisclosure } from '@mantine/hooks';
import './App.css';
import Settings from './views/Settings';
import { Route, Routes, useNavigate } from 'react-router-dom';
import Home from './views/Home';
import { useEffect, useState } from 'react';
import ServiceRedirectHandler from './views/ServiceRedirectHandler';
import { API } from './utils/Axios';
import Playlists from './views/Playlists';

interface INavigationItem {
	id: number;
	label: string;
	to: string;
	icon?: React.ReactNode;
}

const navigationItems: INavigationItem[] = [
	{ id: 0, label: 'Home', to: '/' },
	{ id: 1, label: 'Playlists', to: '/playlists' },
	{ id: 2, label: 'Settings', to: '/settings' },
]

function App() {
	const navigate = useNavigate();
	const [opened, { toggle }] = useDisclosure();
	const [active, setActive] = useState<number>(0);

	useEffect(() => {
		// Call the server to get the session
		const fetchSession = async () => {
			try {
				const response = await API.get('/auth/session');
				if (response.status === 200) {
					// A cookie should be set, in the header
					// Make sure all requests are sent with the cookie
				}
			} catch (error) {
				console.error('Error fetching session', error);
			}
		}

		fetchSession();

		const activeItem = navigationItems.find((item) => item.to === window.location.pathname);
		if (activeItem) {
			setActive(activeItem.id);
		}
	}, []);


	const handleNavigation = (to: string, id: number) => () => {
		if (to === window.location.pathname) {
			// Already on the page, do nothing
			return;
		}
		navigate(to);
		setActive(id);
	}

	return <MantineProvider>
		{
			<AppShell
				header={{ height: 60 }}
				navbar={{
					width: 180,
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
					<h1>Playlist Swap</h1>
				</AppShell.Header>

				<AppShell.Navbar p="sm">
					{/* Add the Nav Links below */}
						{navigationItems.map((item: INavigationItem, index: number) => {
							return (
								<NavLink
									key={item.label}
									active={active === index}
									label={item.label}
									onClick={handleNavigation(item.to, index)}
								/>
							)
						})}
				</AppShell.Navbar>

				<AppShell.Main>
				{/* Add the Routes below */}
					<Routes>
						<Route path="/" element={<Home />} />
						<Route path="/playlists" element={<Playlists />} />
						<Route path="/settings" element={<Settings />} />
						<Route path="/settings/:type/redirect" element={<ServiceRedirectHandler />} />
					</Routes>
				</AppShell.Main>
			</AppShell>
		}
	</MantineProvider>
}

export default App
