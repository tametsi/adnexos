import svelte from '@astrojs/svelte';
import tailwind from '@astrojs/tailwind';
import astroPwa from '@vite-pwa/astro';
import { defineConfig } from 'astro/config';
import manifest from './src/manifest';

// https://astro.build/config
export default defineConfig({
	vite: {
		define: {
			__APP_VERSION__: `'${process?.env?.VERSION ?? 'UNKNOWN'}'`,
		},
	},
	integrations: [
		tailwind(),
		svelte(),
		astroPwa({
			manifest,
			registerType: 'prompt',
			includeAssets: ['logo.svg', 'apple-touch-icon-180x180.png'],
			workbox: {
				ignoreURLParametersMatching: [/.*/],
				navigateFallbackDenylist: [
					/^\/_(\/.*)?$/,
					/^\/api\/.*$/,
					/^\/\.well-known\/.*$/,
					/^.*\.txt$/,
				],
			},
		}),
	],
});
