import axios from 'axios';
import {HttpRequestProps} from '../../types/Types';

const APIUrl = process.env.REACT_APP_API_URL;
const frontendUrl = process.env.REACT_APP_FRONTEND_URL;

export const HttpRequest = async <T, D>({path, method, data, token, states, navigate}: HttpRequestProps<T, D>) => {
	const withCredentials: boolean = true;
	try {
		const config = {
			method,
			url: `${APIUrl}/${path}`,
			headers: {
				Authorization: `Bearer ${token}`,
			},
			data,
			withCredentials,
		};

		if (method === 'GET' || method === 'DELETE') {
			delete config.data;
		}

		const response = await axios.request<T>(config);
		states(response.data);
	} catch (err) {
		if (axios.isAxiosError(err) && err.response?.status === 401 && navigate) {
			console.error('Unauthorized access');
			navigate(`/login`);
		}
		console.error('Error sending request:', err);
	}
};
