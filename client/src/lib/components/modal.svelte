<script context="module">
	/**
	 * Request to change the state of the machine given and event.
	 * @param {boolean} state
	 * @param {boolean} event
	 */
	export function modalMachine(state, event) {
		switch (state) {
			case true:
				if (event == false) {
					return false;
				}
			case false:
				if (event == true) {
					return true;
				}
			default:
				return state;
		}
	}
</script>

<script>
	import { createEventDispatcher } from "svelte";

	/** @type {boolean} */
	export let isOpenModal;

	/** @type {import('svelte').EventDispatcher<{closeModal: {open: boolean}}>}*/
	const dispatch = createEventDispatcher();

	function dispatchClose() {
		dispatch('closeModal', {
			open: false
		})
	}
</script>

<!-- You can open the modal using ID.showModal() method -->
<dialog class="modal" class:modal-open={isOpenModal}>
	<form method="dialog" class="modal-box">
		<button
			class="btn btn-sm btn-circle btn-ghost absolute right-2 top-2"
			on:click={dispatchClose}>✕</button
		>
		<slot />
	</form>
</dialog>
