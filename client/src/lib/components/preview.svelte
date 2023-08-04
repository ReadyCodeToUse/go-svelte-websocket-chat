<script>
	/** @type {string}*/
	export let id;

	/** @type {string} */
	export let title;

	/** Toggle between selected and unselected state*/
	export function toggleSelection() {}
</script>

<div
	{id}
	class="w-full h-20 p-8 bg-neutral hover:bg-neutral-focus selection:bg-neutral-content divide-solid"
>
	{title}
</div>
