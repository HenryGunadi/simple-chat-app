import Cookies from 'js-cookie';

export function getUserSession(name: string) {
	return Cookies.get(name);
}
