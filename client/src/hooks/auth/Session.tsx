export const getCookie = (name: string): string | undefined => {
	const value = `; ${document.cookie}`;
	const parts = value.split(`; ${name}=`);
	if (parts.length === 2) return parts.pop()?.split(';').shift();
	return undefined;
};

export const decodeBase64 = (str: string): string => {
	// Replace URL-safe Base64 characters with standard Base64 characters
	const base64 = str.replace(/-/g, '+').replace(/_/g, '/');
	// Add padding if necessary
	const padding = base64.length % 4 === 0 ? '' : '='.repeat(4 - (base64.length % 4));
	const base64WithPadding = base64 + padding;

	console.log('Base64 with padding:', base64WithPadding);

	try {
		// Decode Base64
		const decoded = atob(base64WithPadding);
		console.log('Base64 decoded:', decoded);

		// Return the decoded value directly without URI decoding
		return decoded;
	} catch (e) {
		console.error('Failed to decode Base64:', e);
		return '';
	}
};

export const parseCookieValue = (encodedValue: string): any => {
	const decodedValue = decodeBase64(encodedValue);
	return JSON.parse(decodedValue);
};
