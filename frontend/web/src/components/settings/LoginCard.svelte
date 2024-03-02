<script lang="ts">
	import DialogCard from '@/components/DialogCard.svelte';
	import OAuth2Login from '@/components/settings/OAuth2Login.svelte';
	import pb from '@/lib/pb';
	import { onMount } from 'svelte';

	let identity = '',
		password = '',
		redirect = '';
	const login = async () =>
		await $pb
			.collection('users')
			.authWithPassword(identity, password, { expand: 'settings_via_user' })
			.then(() => window.location.replace(redirect || '/settings'))
			.catch();

	onMount(() => (redirect = new URLSearchParams(window.location.search).get('redirect') || ''));
</script>

<DialogCard backUrl={redirect ? '' : '/settings'} on:submit={login}>
	<svelte:fragment slot="title">Login</svelte:fragment>

	<!-- email or username -->
	<label class="form-control w-full">
		<div class="label">
			<span class="label-text">Email or username</span>
		</div>
		<input
			type="text"
			bind:value={identity}
			required
			placeholder="email or username"
			class="input input-bordered w-full"
		/>
	</label>

	<!-- password -->
	<label class="form-control w-full">
		<div class="label">
			<span class="label-text">Password</span>
		</div>
		<input
			type="password"
			bind:value={password}
			required
			placeholder="password"
			class="input input-bordered w-full"
		/>
	</label>

	<!-- actions -->
	<svelte:fragment slot="actions">
		<button class="btn btn-primary" type="submit">Login</button>
		<a href="/signup{redirect ? `?redirect=${redirect}` : ''}" class="btn btn-ghost">
			Sign up instead
		</a>
	</svelte:fragment>

	<svelte:fragment slot="bottom">
		<div class="divider">OR</div>
		<OAuth2Login {redirect} />
	</svelte:fragment>
</DialogCard>
