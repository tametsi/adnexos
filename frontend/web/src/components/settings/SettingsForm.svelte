<script lang="ts">
	import pb, { auth } from '@/lib/pb';

	const readSettings = () => ({
		theme: $auth?.theme || '',
	});

	let settings = readSettings();
	// super awesome way of resetting the settings <3
	auth.subscribe(() => (settings = readSettings()));

	const update = () => $pb.collection('users').update($auth?.id, settings);
</script>

<form
	on:submit|preventDefault={update}
	class="mx-auto flex w-64 max-w-full flex-col items-center gap-4"
>
	<div class="form-control w-full">
		<div class="label">
			<span class="label-text">Theme</span>
		</div>

		<div class="join w-full justify-between">
			<input
				type="radio"
				name="theme-buttons"
				bind:group={settings.theme}
				class="btn theme-controller join-item grow"
				aria-label="System"
				value=""
			/>
			<input
				type="radio"
				name="theme-buttons"
				bind:group={settings.theme}
				class="btn theme-controller join-item grow"
				aria-label="Light"
				value="light"
			/>
			<input
				type="radio"
				name="theme-buttons"
				bind:group={settings.theme}
				class="btn theme-controller join-item grow"
				aria-label="Dark"
				value="dark"
			/>
		</div>
	</div>

	<div class="flex w-full flex-col items-center gap-2 py-4">
		<button type="submit" class="btn btn-sm btn-primary btn-outline btn-wide max-w-full">
			Save
		</button>
	</div>
</form>
