/** @type {import('tailwindcss').Config} */
export default {
	content: ['./src/**/*.{astro,html,js,jsx,md,mdx,svelte,ts,tsx,vue}'],
	theme: {
		extend: {},
	},
	daisyui: {
		themes: [
			{
				default: {
					...require('daisyui/src/theming/themes')['light'],
					primary: '#a3ccaa',
					secondary: '#fca5a5',
					'secondary-content': '#1f2937',
				},
				dark: {
					...require('daisyui/src/theming/themes')['dracula'],
					primary: '#a3ccaa',
					secondary: '#fca5a5',
				},
			},
		],
	},
	plugins: [require('daisyui')],
};
