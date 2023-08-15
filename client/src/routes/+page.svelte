<script>
	import { useMachine } from '$lib';
	import Chatbox from '$lib/components/chatbox.svelte';
	import Modal, { modalMachine } from '$lib/components/modal.svelte';
	import Preview from '$lib/components/preview.svelte';
	import Room from '$lib/components/room.svelte';
	import { onMount } from 'svelte';
	import { writable } from 'svelte/store';

	/** @type {import('svelte/store').Writable< {room: string, messages: import('$lib/components/message').Message[]}[]>} */
	let rooms = writable([
		{
			room: 'welcome',
			messages: [{ sender: false, content: 'Welcome to this message client', name: 'app' }]
		}
	]);

	/** @type {Preview[]} */
	let previews = [];

	/** @type {Room}*/
	let room;

	/** @type {{active: number|undefined}}*/
	let selectedPreview = {
		active: undefined
	};

	/**
	 * @param {{action: string, id?: number}} event
	 */
	function chnageSelected(event) {
		switch (event.action) {
			case 'activate':
				if (selectedPreview.active !== undefined) {
					previews[selectedPreview.active].toggleSelection();
				}
				if (event.id !== undefined) {
					previews[event.id].toggleSelection();

					let previewName = previews[event.id].getTitle();
					let filtered = $rooms.filter((obj) => obj.room === previewName);
					let oldMessages = room.updateMessages(filtered[0].messages);
					if (selectedPreview.active !== undefined) {
						let oldTitle = previews[selectedPreview.active].getTitle();
						$rooms.map((obj) => {
							if (obj.room === oldTitle) {
								obj.messages = [...oldMessages];
							}
						});
					}

					selectedPreview.active = event.id;
				}

				break;

			case 'deactivate':
				if (selectedPreview.active !== undefined) {
					previews[selectedPreview.active].toggleSelection();
					selectedPreview.active = undefined;
				}
				break;
		}
	}

	const modal = useMachine(modalMachine, false);
	$: modalSate = modal.state;

	/**
	 * Create a new room if the name doesn't exists yet.
	 * @param {string} name
	 */
	function createRoom(name) {
		let filtered = $rooms.filter((obj) => obj.room === name);

		if (filtered.length > 0) return;

		/** @type {{room: string, messages: import('$lib/components/message').Message[]}}*/
		let messages = { room: name, messages: [{ content: 'Room created', sender: true }] };
		$rooms = [messages, ...$rooms];
	}

	onMount(() => {
		document.body.addEventListener('keyup', (e) => {
			switch (e.key) {
				case 'Escape':
					if ($modalSate === true) $modalSate = false;
					else chnageSelected({ action: 'deactivate' });
					break;

				default:
					break;
			}
		});
	});
</script>

<div class="container mx-auto grid grid-cols-10 gap-1 h-screen">
	<div class="col-span-3 row-span-5 overflow-auto py-8">
		<div
			class="sticky top-0 bg-inherit h-20 flex flex-row justify-between items-center bg-neutral-focus px-8"
		>
			<h1 class="text-lg font-bold">chat applcation</h1>
			<button
				class="text-xl bg-accent btn"
				title="Create a new room"
				on:click={() => {
					modal.send(true);
				}}>+</button
			>
		</div>
		{#each $rooms as room, i}
			<a
				class="w-full"
				href={null}
				on:click={() => {
					chnageSelected({ action: 'activate', id: i });
				}}
			>
				<Preview title={`${room.room}`} bind:this={previews[i]} />
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

	<Modal bind:isOpenModal={$modalSate} on:closeModal={(e) => modal.send(e.detail.open)}>
		<h3 class="text-2xl">New room name</h3>
		<Chatbox
			on:message={(e) => {
				createRoom(e.detail.content);
				modal.send(false);
			}}
		/>
	</Modal>
</div>

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
