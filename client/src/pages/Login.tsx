import {useEffect, useState} from 'react';
import {Button} from '../components/common/Button';
import {HttpRequest} from '../hooks/api/HTTPReqs';
import UseNavigation from '../hooks/navigate/Navigate';
import {LoginResponse} from '../types/Types';

const Login: React.FC = () => {
	const [loginStatus, setIsLoginStatus] = useState<LoginResponse | undefined>(undefined);

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
