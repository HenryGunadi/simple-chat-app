import {useEffect, useState} from 'react';
import {HttpRequest} from '../hooks/api/HTTPReqs';
import {useNavigate} from 'react-router-dom';
import Button from '../components/common/Button';
import {handleInput} from '../hooks/common/Hook';
import {ChatType} from '../types/Types';

const Home: React.FC = () => {
	const [cookiesVal, setCookieVal] = useState<any>(undefined);
	const [message, setMessage] = useState<string>('');
	const [displayMessages, setDisplayMessages] = useState<ChatType[]>([]);
	const [socket, setSocket] = useState<WebSocket | null>(null);

	const navigate = useNavigate();

	useEffect(() => {
		HttpRequest<any, undefined>({path: 'user', method: 'GET', token: '', states: setCookieVal, navigate: navigate});
	}, []);

	useEffect(() => {
		console.log('user data : ', cookiesVal);
	}, [cookiesVal]);

	// connecting to ws connection
	useEffect(() => {
		const ws = new WebSocket(`ws://localhost:8080/ws`);
		setSocket(ws);

		ws.onopen = () => {
			console.log('Successfully Connected');
			ws.send(JSON.stringify('Hi From the Client!'));
		};

		ws.onmessage = (event) => {
			try {
				const data = JSON.parse(event.data);
				console.log(data.text);
				setDisplayMessages((prev) => [
					...prev,
					{
						id: data.HEADERS.ClientID,
						message: data.text,
					},
				]);
			} catch (err) {
				console.error('error parsing JSON data : ', err);
			}
		};

		ws.onclose = (event) => {
			console.log('Socket Closed Connection: ', event);
			ws.send('Client Closed!');
		};

		ws.onerror = (error) => {
			console.log('Socket Error: ', error);
		};

		return () => {
			if (ws) {
				ws.close();
			}
		};
	}, []);

	const handleSendButton = () => {
		if (socket && socket.readyState === WebSocket.OPEN) {
			socket.send(JSON.stringify({text: message}));
			setMessage('');
			console.log('is sent');
		}
	};

	const handleEnterKey = (e: React.KeyboardEvent<HTMLInputElement>) => {
		if (e.key === 'Enter') {
			handleSendButton();
		}
	};

	return (
		<div className="w-screen h-screen flex justify-center items-center bg-zinc-800">
			<div className="flex flex-col gap-6 justify-between w-2/3 h-5/6 border border-gray-500 rounded-md bg-zinc-100 p-6">
				<div className="flex-1">
					{displayMessages &&
						displayMessages.map((value: ChatType, index) => {
							return (
								<div key={index} className="w-full flex gap-2">
									<h2 className="font-semibold text-blue-500">{value.id} : </h2>
									<p>{value.message}</p>
								</div>
							);
						})}
				</div>

				<div className="w-full rounded-xl bg-zinc-300 py-2 px-3">
					<input
						type="text"
						className="px-3 py-1 border-none outline-none w-5/6 bg-transparent"
						onChange={(event) => {
							handleInput(event, setMessage);
						}}
						onKeyDown={handleEnterKey}
						value={message}
					/>

					<Button className="w-1/6" content="Send" variant="default" onClick={handleSendButton}></Button>
				</div>
			</div>
		</div>
	);
};

export default Home;
