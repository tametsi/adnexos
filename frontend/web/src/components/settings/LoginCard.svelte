<script lang="ts">
	import DialogCard from '@/components/DialogCard.svelte';
	import pb from '@/lib/pb';

	let identity = '',
		password = '';
	const login = async () =>
		await $pb
			.collection('users')
			.authWithPassword(identity, password)
			.then(() => window.location.replace('/settings'))
			.catch();
</script>

<DialogCard backUrl="/settings" on:submit={login}>
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
		<a href="/signup" class="btn btn-ghost">Sign up instead</a>
	</svelte:fragment>
</DialogCard>
