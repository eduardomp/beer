// since there's no dynamic data here, we can prerender

import type { PageLoad } from "./$types";

export const load: PageLoad = async ({ fetch, params }) => {
    const res = await fetch(`/api/v1/apps`);
    const response : App[] = await res.json();

    return { response } ;
};