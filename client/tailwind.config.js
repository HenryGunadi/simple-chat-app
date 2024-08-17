/** @type {import('tailwindcss').Config} */
module.exports = {
	content: ['./src/**/*.{js,jsx,ts,tsx}', './public/index.html'],
	theme: {
		extend: {
			fontFamily: {
				nunito: ['Nunito', 'system-ui'],
			},
		},
	},
	plugins: [
		function ({addUtilities}) {
			addUtilities(
				{
					'.font-smooth': {
						'-webkit-font-smoothing': 'antialiased',
						'-moz-osx-font-smoothing': 'grayscale',
					},
				},
				['responsive', 'hover']
			);
		},
	],
};
