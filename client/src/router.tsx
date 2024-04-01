import { createBrowserRouter } from "react-router-dom";
import App from './App.tsx';

const router = createBrowserRouter([
	{
		path: "/",
		element: <App />,
		children: [
			{
				path: "/profile"
			},
			{
				path: "/settings"
			},
			{
				path: "/settings/:type/redirect"
			}
		]
	}
]);

export default router;