import './App.css';
import Home from './pages/Home';
import Login from './pages/Login';
import {Navigate, Route, Routes} from 'react-router-dom';

function App() {
	return (
		<div className="font-nunito antialiased">
			<Routes>
				<Route path="/" element={<Home />} />
				<Route path="/login" element={<Login />} />
			</Routes>
		</div>
	);
}

export default App;
