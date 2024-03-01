import svelte from '@astrojs/svelte';
import tailwind from '@astrojs/tailwind';
import astroPwa from '@vite-pwa/astro';
import { defineConfig } from 'astro/config';
import manifest from './src/manifest';

// https://astro.build/config
export default defineConfig({
	integrations: [
		tailwind(),
		svelte(),
		astroPwa({
			manifest,
			registerType: 'prompt',
			includeAssets: ['logo.svg', 'apple-touch-icon-180x180.png'],
			workbox: {
				ignoreURLParametersMatching: [/.*/],
			},
		}),
	],
});
