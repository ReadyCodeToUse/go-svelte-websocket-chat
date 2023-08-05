<script>
	import Preview from '$lib/components/preview.svelte';
	import { useMachine } from '$lib/stateMachine';

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
	
	/**
	 * state: {active: number|undefined}
	 * event: {action: string, id: number|undefined}
	 */


	/**
	 * Usually there should be a switch case over the current state to 
	 * handle the event. In this implementation of this state machine 
	 * switch is over the event to adapt the state to it.
	 * @param {{active: number|undefined}} state
	 * @param {{action: string, id: number|undefined}} event
	 */
	function chatMachine(state, event) {
		switch (event.action) {
			case 'activate':
				if (state.active !== undefined) {
					previews[state.active].toggleSelection()
				}
				if (event.id !== undefined) {
					previews[event.id].toggleSelection()
					state.active = event.id
				}
				break

			case 'deactivate':
				if (state.active !== undefined) {
					previews[state.active].toggleSelection()
					state.active = undefined
				}
				break

			default:
				return state;
		}
	}

	const { state, send } = useMachine(chatMachine, {active: undefined})

	state.subscribe((v) => console.log(v))

	/** @type {Preview[]} */
	let previews = [];
</script>

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

<main class="grid grid-cols-10 grid-rows-auto content-start h-screen">
	<div class="row-span-1 col-span-full h-20 flex items-center">
		<h1 class="text-lg font-bold">chat applcation</h1>
	</div>
	<div class="col-start-1 col-span-3 row-start-2 h-auto bg-purple-300 overflow-auto">
		{#each titles as title, i}
			<a
				class="w-full"
				href={null}
				on:click={() => {
					// previews[i].toggleSelection();
					send({action: 'activate', id: i})
				}}
			>
				<Preview title={`${i}: ${title}`} bind:this={previews[i]} />
			</a>
		{/each}
	</div>
</main>

<!-- {#if connection}
	connected
{:else}
	<button on:click={() => connection = !connection}>connect</button>
{/if} -->
