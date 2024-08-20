import './App.css';
import Home from './pages/Home';
import Login from './pages/Login';
import {Navigate, Route, Routes, redirect} from 'react-router-dom';

function App() {
	return (
		<div className="font-nunito antialiased">
			<Routes>
				<Route path="/" element={<Navigate to="/login" />} />
				<Route path="/login" element={<Login />} />
				<Route path="/home" element={<Home />} />
			</Routes>
		</div>
	);
}

export default App;
