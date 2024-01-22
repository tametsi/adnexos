import consts from './consts';
import type { ManifestOptions } from 'vite-plugin-pwa';

export default {
	name: consts.DEFAULT_TITLE,
	short_name: consts.DEFAULT_TITLE,
	description: consts.DEFAULT_DESC,
	theme_color: consts.THEME_COLOR,
	background_color: consts.COLOR.BACKGROUND,

	icons: [
		{
			src: 'pwa-64x64.png',
			sizes: '64x64',
			type: 'image/png',
		},
		{
			src: 'pwa-192x192.png',
			sizes: '192x192',
			type: 'image/png',
		},
		{
			src: 'pwa-512x512.png',
			sizes: '512x512',
			type: 'image/png',
		},
		{
			src: 'maskable-icon-512x512.png',
			sizes: '512x512',
			type: 'image/png',
			purpose: 'maskable',
		},
	],
} satisfies Partial<ManifestOptions>;
