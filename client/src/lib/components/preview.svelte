<script>
	/** @type {string} */
	export let title;

	/** @type {boolean} */
	let selected = false;

	/** Toggle between selected and unselected state*/
	export function toggleSelection() {
		selected = !selected;
	}
</script>

<div
	class="w-full h-20 p-8 {selected === true
		? 'bg-neutral-content text-neutral'
		: 'bg-neutral hover:bg-neutral-focus'}  selection:bg-neutral-content divide-solid"
>
	{title}
</div>
