<script lang="ts">
	import { error } from '@/lib/alert';
	import { preventDefault } from '@/lib/event';
	import pb, { auth } from '@/lib/pb';
	import { type RecordModel } from 'pocketbase';

	let settings: Partial<RecordModel> = $derived(
		$auth?.expand?.settings_via_user?.[0] || { theme: '' },
	);

	const update = () => {
		const saveSettings = (settings: RecordModel) => {
			const authStore = $pb.authStore;
			if (!authStore.record) return;

			const record = { ...authStore.record };
			((record.expand ??= {}).settings_via_user ??= [])[0] = settings;

			authStore.save(authStore.token, record);
		};

		// update existing settings
		if (settings.id)
			return $pb
				.collection('settings')
				.update(settings.id, settings)
				.then(saveSettings)
				.catch(error('Failed to save settings.'));

		// create new settings
		settings.user = $auth?.id;
		$pb.collection('settings')
			.create(settings)
			.then(saveSettings)
			.catch(error('Failed to create settings.'));
	};
</script>

<form
	onsubmit={preventDefault(update)}
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
