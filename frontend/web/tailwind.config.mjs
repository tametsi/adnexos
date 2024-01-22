import consts from './src/consts';

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
					primary: consts.COLOR.PRIMARY,
					secondary: consts.COLOR.SECONDARY,
					'secondary-content': consts.COLOR.SECONDARY_CONTENT,
				},
				dark: {
					...require('daisyui/src/theming/themes')['dracula'],
					primary: consts.COLOR.PRIMARY,
					secondary: consts.COLOR.SECONDARY,
				},
			},
		],
	},
	plugins: [require('daisyui')],
};
