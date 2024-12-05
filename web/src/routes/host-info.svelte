<script lang="ts">
	import { writable } from 'svelte/store';
    import { onMount } from 'svelte';
    
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

    let usedMemoryColor = writable<string>('');
    let freeMemoryColor = writable<string>('');
    let availableMemoryColor = writable<string>('');

    
    function convertToBytes(input: string): number {
        let result = 0;
        let marker = 1024;
        let inputValue = Number.parseFloat(input.replace(/\D/g,''));

        if (input.endsWith('TB')) {
            result = inputValue * marker * marker * marker * marker;
        }

        if (input.endsWith('GB')) {
            result = inputValue * marker * marker * marker;
        }

        if (input.endsWith('MB')) {
            result = inputValue * marker * marker;
        }

        if (input.endsWith('KB')) {
            result = inputValue * marker;
        }

        return result;

    }

    function getMemoryColorNearTotalIsWorse(totalBytes: number, usedBytes: number): string {
        let result = 'text-green-500';
        
        if (usedBytes > (totalBytes * 0.5) && usedBytes <= (totalBytes * 0.75)) {
            result = 'text-yellow-500';
        }

        if (usedBytes > (totalBytes * 0.75)) {
            result = 'text-red-500';
        }

        return result;
    }

    function getMemoryColorNearTotalIsBetter(totalBytes: number, usedBytes: number): string {
        let result = 'text-green-500';
        
        if (usedBytes < (totalBytes * 0.5) && usedBytes >= (totalBytes * 0.25)) {
            result = 'text-yellow-500';
        }
        
        if (usedBytes < (totalBytes * 0.25)) {
            result = 'text-red-500';
        }
        
        return result;
    }

    function getUsedMemoryColor(memo: Memo): string {
        let total = convertToBytes(memo.Total);
        let used = convertToBytes(memo.Used);
        
        return getMemoryColorNearTotalIsWorse(total, used);
    }

    function getFreeMemoryColor(memo: Memo): string {
        let total = convertToBytes(memo.Total);
        let free = convertToBytes(memo.Free);
        
        return getMemoryColorNearTotalIsBetter(total, free);
    }

    function getAvailableMemoryColor(memo: Memo): string {
        let total = convertToBytes(memo.Total);
        let available = convertToBytes(memo.Available);
        
        return getMemoryColorNearTotalIsBetter(total, available);
    }

    onMount(() => {
        //TODO: Externalize the SSE endpoint
	    let evtSource = new EventSource('http://localhost:9999/sse');

        evtSource.addEventListener('hostInfo', function (event: MessageEvent) {
            var dataobj: HostInfo = JSON.parse(event.data);
            hostInfo.set(dataobj);
            usedMemoryColor.set(getUsedMemoryColor(dataobj.Memo));
            freeMemoryColor.set(getFreeMemoryColor(dataobj.Memo));
            availableMemoryColor.set(getAvailableMemoryColor(dataobj.Memo));
        });
    })

</script>

<div class="w-auto text-xs font-extrabold text-white bg-[#2C2C2C] rounded-sm p-1">
 	<p>
        {$hostInfo.Host} | 
        {$hostInfo.OS} {$hostInfo.Arch} | 
        MEM: 
        Total <span class="text-blue-500">{$hostInfo.Memo?.Total}</span> 
        Used  <span class="{$usedMemoryColor}">{$hostInfo.Memo?.Used}</span>
        Free  <span class="{$freeMemoryColor}">{$hostInfo.Memo?.Free}</span>
        Available <span class="{$availableMemoryColor}">{$hostInfo.Memo?.Available}</span>
    </p>
</div>
