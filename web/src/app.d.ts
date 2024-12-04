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
	namespace App {
		// interface Error {}
		// interface Locals {}
		// interface PageData {}
		// interface PageState {}
		// interface Platform {}
	}
}

export {};
