import {useEffect} from 'react';
import {Button} from '../components/common/Button';
import {getUserSession} from '../hooks/auth/Auth';
import {useNavigate} from 'react-router-dom';

const Login: React.FC = () => {
	const navigate = useNavigate();

	useEffect(() => {
		const cookie = getUserSession('user');
		if (cookie) {
			localStorage.setItem('cookie', cookie);
		}

		if (cookie) {
			navigate('/home');
		}
	}, []);

	const handleLogin = () => {
		window.location.href = 'http://localhost:8080/auth/discord';
	};

	return (
		<div>
			<h1 className="text-center">Login</h1>
			<div className="w-screen flex justify-center">
				<Button content="Login With Discord" onClick={handleLogin} variant="default" className=""></Button>
			</div>
		</div>
	);
};

export default Login;
