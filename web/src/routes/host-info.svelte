<script lang="ts">
	import { writable } from 'svelte/store';

	interface Memo {
		Total: string;
		Free: string;
		Used: string;
		Available: string;
	}

	interface HostInfo {
		Host: string;
		OS: string;
		Arch: string;
		Memo: Memo;
	}

	let messages = writable<HostInfo>({
		Host: '',
		OS: '',
		Arch: '',
		Memo: {
			Total: '',
			Free: '',
			Used: '',
			Available: ''
		}
	});

	let evtSource = new EventSource('http://localhost:9999/sse');

	evtSource.addEventListener('hostInfo', function (event: MessageEvent) {
		var dataobj: HostInfo = JSON.parse(event.data);
		messages.set(dataobj);
	});
</script>

<div class="text-column w-52 content-start text-sm text-white">
	<h2>Host Info</h2>
	<p>OS: {$messages.OS}</p>
	<p>Arch: {$messages.Arch}</p>
	<p>Name: {$messages.Host}</p>
	<p>Memory:</p>
	<ul class="ml-4">
		<li>Total: {$messages.Memo?.Total}</li>
		<li>Free: {$messages.Memo?.Free}</li>
		<li>Used: {$messages.Memo?.Used}</li>
		<li>Available: {$messages.Memo?.Available}</li>
	</ul>
</div>
