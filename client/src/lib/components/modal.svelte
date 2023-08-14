<script>
	import { createEventDispatcher } from "svelte";

    /** @type {boolean} */
    export let isOpenModal;

</script>

<!-- You can open the modal using ID.showModal() method -->
<dialog class="modal" class:modal-open={isOpenModal}>
	<form method="dialog" class="modal-box">
		<button class="btn btn-sm btn-circle btn-ghost absolute right-2 top-2" on:click={() => isOpenModal = false}>✕</button>
        <slot />
	</form>
</dialog>
