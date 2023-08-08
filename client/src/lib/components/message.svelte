<script context="module">
	
</script>
<script>
    /** @type {boolean} */
    export let sender

	/** @type {string} */
	export let content
</script>

<div class="chat {sender ? 'chat-end' : 'chat-start'}">
	<div class="chat-header">
		Obi-Wan Kenobi
		<!-- <time class="text-xs opacity-50">2 hours ago</time> -->
	</div>
	<div class="chat-bubble">{content}</div>
	<!-- <div class="chat-footer opacity-50">Seen</div> -->
</div>
