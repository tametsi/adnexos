<script lang="ts">
	import DialogCard from '@/components/DialogCard.svelte';
	import OAuth2Login from '@/components/settings/OAuth2Login.svelte';
	import { error } from '@/lib/alert';
	import pb from '@/lib/pb';
	import { onMount } from 'svelte';

	let identity = $state(''),
		password = $state(''),
		redirect = $state('');
	const login = () =>
		$pb
			.collection('users')
			.authWithPassword(identity, password, { expand: 'settings_via_user' })
			.then(() => window.location.replace(redirect || '/groups'))
			.catch(error('Login failed.'));

	onMount(() => (redirect = new URLSearchParams(window.location.search).get('redirect') || ''));
</script>

<DialogCard backUrl="/" onsubmit={login}>
	{#snippet title()}
		Login
	{/snippet}

	<!-- email or username -->
	<label class="fieldset">
		<span class="label">Email or username</span>
		<input
			type="text"
			bind:value={identity}
			required
			placeholder="Email or username"
			class="input w-full"
		/>
	</label>

	<!-- password -->
	<label class="fieldset">
		<span class="label">Password</span>
		<input
			type="password"
			bind:value={password}
			required
			placeholder="Password"
			class="input w-full"
		/>
	</label>

	<!-- actions -->
	{#snippet actions()}
		<button class="btn btn-primary" type="submit">Login</button>
		<a href="/signup{redirect ? `?redirect=${redirect}` : ''}" class="btn btn-ghost">
			Sign up instead
		</a>
	{/snippet}

	{#snippet bottom()}
		<div class="divider">OR</div>
		<OAuth2Login {redirect} />
	{/snippet}
</DialogCard>
