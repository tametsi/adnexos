<script lang="ts">
	import DialogCard from '@/components/DialogCard.svelte';
	import pb, { auth } from '@/lib/pb';

	let data = { ...$auth };
	const edit = () =>
		$pb
			.collection('users')
			.update(
				data.id,
				{ username: data.username, name: data.name },
				{ expand: 'settings_via_user' },
			)
			.then(() => window.location.replace('/settings'));
</script>

<DialogCard backUrl="/settings" on:submit={edit}>
	<svelte:fragment slot="title">Edit Profile</svelte:fragment>

	<!-- username -->
	<label class="form-control w-full">
		<div class="label">
			<span class="label-text">Avatar (coming soon™)</span>
		</div>
		<input type="file" class="file-input file-input-bordered w-full" />
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

	<a href="/users/password" class="link">Change Password</a>

	<!-- actions -->
	<svelte:fragment slot="actions">
		<button class="btn btn-primary" type="submit">Edit</button>
		<a href="/settings" class="btn btn-ghost">Nothing to do here</a>
	</svelte:fragment>
</DialogCard>
