type options = {
  internal?: boolean;
};

export function defineApi({ internal = false }: options = {}) {
  const { $i18n } = useNuxtApp();
  const { serverUrl } = useRuntimeConfig().public;

  const api = $fetch.create({
    baseURL: internal ? "/api" : serverUrl,
    onRequest({ options }) {
      const headers = (options.headers ||= {} as Headers);
      const authToken =
        getHeader(headers, "Authorization") ??
        `Bearer ${useCookie("token").value}`;

      console.log(authToken);

      addHeader(headers, "Accept-Language", $i18n.locale.value);
      addHeader(headers, "Authorization", authToken);
    },
  });

  return api;
}

export function addHeader(
  headers: Headers | Array<Record<string, any>> | unknown,
  key: string,
  value: string
) {
  if (Array.isArray(headers)) {
    headers.push([key, value]);
  } else if (headers instanceof Headers) {
    headers.set(key, value);
  } else {
    (headers as Record<string, any>)[key] = value;
  }
}

export function getHeader(
  headers: Headers | Array<Record<string, any>> | unknown,
  key: string
) {
  if (Array.isArray(headers)) {
    return headers.find(([k]) => k === key)?.[1];
  } else if (headers instanceof Headers) {
    return headers.get(key);
  } else {
    return (headers as Record<string, any>)[key];
  }
}
