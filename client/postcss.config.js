const {config} = require('process');

module.exports = {
	plugins: {
		tailwindcss: {config: './tailwind.config.js'},
		autoprefixer: {},
	},
};
