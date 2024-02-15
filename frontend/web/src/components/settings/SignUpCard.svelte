<script lang="ts">
	import DialogCard from '@/components/DialogCard.svelte';
	import OAuth2Login from '@/components/settings/OAuth2Login.svelte';
	import pb from '@/lib/pb';

	let data = {
		email: '',
		username: '',
		name: '',
		password: '',
		passwordConfirm: '',
	};
	const signUp = async () => {
		if (!data.email || data.password !== data.passwordConfirm || data.password.length < 8)
			return;

		if (!data.name) data.name = data.username;

		await $pb.collection('users').create(data).catch();

		await $pb.collection('users').authWithPassword(data.email, data.password).catch();
		window.location.replace('/settings');
	};
</script>

<DialogCard backUrl="/settings" on:submit={signUp}>
	<svelte:fragment slot="title">Sign up</svelte:fragment>

	<!-- email -->
	<label class="form-control w-full">
		<div class="label">
			<span class="label-text">Email</span>
		</div>
		<input
			type="email"
			bind:value={data.email}
			required
			placeholder="email"
			class="input input-bordered w-full"
		/>
	</label>

	<!-- username -->
	<label class="form-control w-full">
		<div class="label">
			<span class="label-text">Username (can be used for login)</span>
		</div>
		<input
			type="text"
			bind:value={data.username}
			placeholder="username"
			class="input input-bordered w-full"
		/>
	</label>
	<!-- name -->
	<label class="form-control w-full">
		<div class="label">
			<span class="label-text">Display Name (will be shown)</span>
		</div>
		<input
			type="text"
			bind:value={data.name}
			placeholder="display name"
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
			bind:value={data.password}
			required
			placeholder="password"
			class="input input-bordered w-full"
		/>
	</label>
	<!-- password - confirm -->
	<label class="form-control w-full">
		<div class="label">
			<span class="label-text">Confirm Password</span>
		</div>
		<input
			type="password"
			bind:value={data.passwordConfirm}
			required
			placeholder="password"
			class="input input-bordered w-full"
		/>
	</label>

	<!-- actions -->
	<svelte:fragment slot="actions">
		<button class="btn btn-primary" type="submit">Sign up</button>
		<a href="/login" class="btn btn-ghost">Login instead</a>
	</svelte:fragment>

	<svelte:fragment slot="bottom">
		<div class="divider">OR</div>
		<OAuth2Login />
	</svelte:fragment>
</DialogCard>
