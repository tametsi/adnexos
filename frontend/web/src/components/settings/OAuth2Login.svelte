<script lang="ts">
	import pb from '@/lib/pb';

	export let redirect = '';

	const oAuth2 = (provider: string) => () =>
		$pb
			.collection('users')
			.authWithOAuth2({ provider, query: { expand: 'settings_via_user' } })
			.then(() => window.location.replace(redirect || '/settings'));
</script>

<button on:click={oAuth2('github')} class="btn btn-secondary btn-outline">
	Continue with GitHub
</button>
<button on:click={oAuth2('discord')} class="btn btn-secondary btn-outline">
	Continue with Discord
</button>
