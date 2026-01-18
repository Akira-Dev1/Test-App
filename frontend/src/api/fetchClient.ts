import { tryRefresh } from "./refresh";
import { createApiError } from "./errors";

const API_URL = import.meta.env.VITE_API_URL;

type FetchOptions = RequestInit & {
  retry?: boolean;
};

export async function fetchClient<T>(
  path: string,
  options: FetchOptions = {}
): Promise<T> {
  const response = await fetch(`${API_URL}${path}`, {
    ...options,
    credentials: "include",
    headers: {
      "Content-Type": "application/json",
      ...options.headers,
    },
  });
  
  console.log("HTTP status:", response.status);

  if (response.ok) {
    return response.json();
  }

  if (response.status == 401 && !options.retry) {
    const refreshed = await tryRefresh();

    if (refreshed) {
      return fetchClient<T>(path, { ...options, retry: true });
    }

    throw createApiError(response);
  }

  if (response.status === 403) {
    throw createApiError(response);
  }

  throw createApiError(response);
};