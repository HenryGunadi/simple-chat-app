import {useEffect, useState} from 'react';
import {HttpRequest} from '../hooks/api/HTTPReqs';
import UseNavigation from '../hooks/navigate/Navigate';
import {useNavigate} from 'react-router-dom';

const Home: React.FC = () => {
	const [cookiesVal, setCookieVal] = useState<any>(undefined);
	const navigate = useNavigate();

	useEffect(() => {
		HttpRequest<any, undefined>({path: 'user', method: 'GET', token: '', states: setCookieVal, navigate: navigate});
	}, []);

	useEffect(() => {
		console.log('user data : ', cookiesVal);
	}, [cookiesVal]);

	return (
		<div className="">
			<h1 className="text-center text-4xl font-semibold">Home</h1>
		</div>
	);
};

export default Home;
