// See https://svelte.dev/docs/kit/types#app.d.ts
// for information about these interfaces
declare global {

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

	interface App {
		Id: string;
		Name: string;
		Ports: number[];
		Host: string;
		Type: string;
		Volumes: string[];
		Dirs: string[];
		Process: Process;
	}

	interface Process {
		Name: string;
		Pid: number;
		Status: string;
		Cmdline: string;
		Connections: any[];
	}
	interface Connection {
		SourceApp: App;
		TargetApp: App;
		Protocol: string;
		TargetPort: number;
		Status: string;
		StartTimestamp: string;
		EndTimestamp: string;
		Duration: string;
	}

	namespace App {
		// interface Error {}
		// interface Locals {}
		// interface PageData {}
		// interface PageState {}
		// interface Platform {}
	}
}

export {};
