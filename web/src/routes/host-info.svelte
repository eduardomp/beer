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

	let hostInfo = writable<HostInfo>({
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
		hostInfo.set(dataobj);
	});
</script>

<div class="text-column w-40 p-5 content-start text-xs text-white bg-[#2C2C2C] rounded-sm">
	<p>{$hostInfo.Host}</p>
	<p>{$hostInfo.OS} - {$hostInfo.Arch}</p>
	<p>Memory:</p>
	<ul class="ml-4">
		<li>Total: {$hostInfo.Memo?.Total}</li>
		<li>Free: {$hostInfo.Memo?.Free}</li>
		<li>Used: {$hostInfo.Memo?.Used}</li>
		<li>Available: {$hostInfo.Memo?.Available}</li>
	</ul>
</div>
