<script>
	import Chatbox from '$lib/components/chatbox.svelte';
	import Modal from '$lib/components/modal.svelte';
	import Preview from '$lib/components/preview.svelte';
	import Room from '$lib/components/room.svelte';
	import { onMount } from 'svelte';
	import { writable } from 'svelte/store';

	let rooms = writable(new Map());
	$: titles = Array.from($rooms.keys());

	/** @type {Preview[]} */
	let previews = [];

	/** @type {Room}*/
	let room;

	/** @type {{active: number|undefined}}*/
	let state = {
		active: undefined
	};

	/** @type {boolean} */
	let isOpenModal = false;

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

<main class="container mx-auto grid grid-cols-10 gap-1 h-auto max-h-screen">
	<div class="col-span-3 row-span-5 overflow-auto py-8">
		<div
			class="sticky top-0 bg-inherit h-20 flex flex-row justify-between items-center bg-neutral-focus px-8"
		>
			<h1 class="text-lg font-bold">chat applcation</h1>
			<button
				class="text-xl bg-accent btn"
				title="Create a new room"
				on:click={() => {
					isOpenModal = true;
				}}>+</button
			>
		</div>
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

	<div class="col-span-7 row-span-5 flex flex-col justify-end h-screen py-8">
		<div class="overflow-auto flex flex-col-reverse">
			<Room bind:this={room} />
		</div>
		<div>
			<Chatbox
				on:message={(e) => {
					room.reciveMessage(e.detail);
				}}
			/>
		</div>
	</div>

	<Modal bind:isOpenModal>
		<h3 class="text-2xl">New room name</h3>
		<Chatbox on:message={(e) => console.log(e.detail)}/>
	</Modal>
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
