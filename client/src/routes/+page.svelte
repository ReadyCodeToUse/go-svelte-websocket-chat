<script>
	import Chatbox from '$lib/components/chatbox.svelte';
	import Preview from '$lib/components/preview.svelte';
	import Room from '$lib/components/room.svelte';
	import { onMount } from 'svelte';

	const titles = [
		'preview 1',
		'preview 2',
		'preview 3',
		'preview 4',
		'preview 5',
		'preview 6',
		'preview 7',
		'preview 8',
		'preview 9'
	];

	/** @type {Preview[]} */
	let previews = [];

	/** @type {Room}*/
	let room;

	/** @type {{active: number|undefined}}*/
	let state = {
		active: undefined
	};

	/**
	 * @param {{action: string, id?: number}} event
	 */
	function chnageSelected(event) {
		switch (event.action) {
			case 'activate':
				if (state.active !== undefined) {
					previews[state.active].toggleSelection();
				}
				if (event.id !== undefined) {
					previews[event.id].toggleSelection();
					state.active = event.id;
				}
				break;

			case 'deactivate':
				if (state.active !== undefined) {
					previews[state.active].toggleSelection();
					state.active = undefined;
				}
				break;
		}
	}

	onMount(() => {
		document.body.addEventListener('keyup', () => {
			chnageSelected({ action: 'deactivate' });
		});
	});
</script>

<main class="grid grid-cols-10 grid-rows-auto gap-1 content-start h-screen max-h-screen">
	<div class="row-span-1 col-start-1 col-end-4 h-20 flex items-center">
		<h1 class="text-lg font-bold">chat applcation</h1>
	</div>
	<div class="row-span-1 col-start-4 col-span-full h-20 flex items-center">
		<h1 class="text-lg font-bold">Selected preview name</h1>
	</div>
	<div class="col-start-1 col-span-3 row-start-2 row-span-3 h-auto overflow-auto">
		{#each titles as title, i}
			<a
				class="w-full"
				href={null}
				on:click={() => {
					chnageSelected({ action: 'activate', id: i });
				}}
			>
				<Preview title={`${i}: ${title}`} bind:this={previews[i]} />
			</a>
		{/each}
	</div>
	<div
		class="col-start-4 col-span-full row-start-2 row-span-2 h-auto overflow-auto flex flex-col-reverse"
	>
		<Room bind:this={room} />
	</div>
	<div class="col-start-4 col-span-full row-start-4 row-span-1 h-auto overflow-auto p-2">
		<Chatbox
			on:message={(e) => {
				room.reciveMessage(e.detail);
			}}
		/>
	</div>
</main>

<!-- <script>
    import {preventTabClose} from "$lib/prevetClose";

    let shouldWork = true;
</script>

<button on:click={() => shouldWork = !shouldWork}>
	change: {shouldWork}
</button>

<div use:preventTabClose={shouldWork}>
	Here would go a huge form that user doesn't want to lose in case of an accident.
</div> -->

<!-- {#if connection}
	connected
{:else}
	<button on:click={() => connection = !connection}>connect</button>
{/if} -->
