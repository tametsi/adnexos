<script lang="ts">
	import { error } from '@/lib/alert';
	import pb from '@/lib/pb';

	interface Props {
		redirect?: string;
	}

	let { redirect = '' }: Props = $props();

	const oAuth2 = (provider: string) => () =>
		$pb
			.collection('users')
			.authWithOAuth2({ provider, query: { expand: 'settings_via_user' } })
			.then(() => window.location.replace(redirect || '/settings'))
			.catch(error('Failed to authenticate.'));
</script>

<button onclick={oAuth2('github')} class="btn btn-secondary dark:btn-outline">
	Continue with GitHub
</button>
<button onclick={oAuth2('discord')} class="btn btn-secondary dark:btn-outline">
	Continue with Discord
</button>
