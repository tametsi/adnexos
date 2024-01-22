import type { Preset } from '@vite-pwa/assets-generator/config';
import {
	createAppleSplashScreens,
	defineConfig,
	minimal2023Preset,
} from '@vite-pwa/assets-generator/config';
import consts from './src/consts';

const preset: Preset = {
	...minimal2023Preset,

	maskable: {
		sizes: [512],
		resizeOptions: { background: consts.COLOR.BACKGROUND },
	},
	apple: {
		sizes: [180],
		resizeOptions: { background: consts.COLOR.BACKGROUND },
	},

	appleSplashScreens: createAppleSplashScreens({
		darkResizeOptions: { background: consts.COLOR.BACKGROUND },
		resizeOptions: { background: consts.COLOR.BACKGROUND_LIGHT },
	}),
};

export default defineConfig({
	headLinkOptions: {
		preset: '2023',
	},
	preset: preset,
	images: ['public/logo.svg'],
	manifestIconsEntry: false,
});
