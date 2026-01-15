import { tryRefresh } from "./refresh";
import { createApiError } from "./errors";
import { redirectToRoot } from "../auth/redirect";

const API_URL = "http://localhost:8081";

type FetchOptions = RequestInit & {
  retry?: boolean;
};

export async function fetchClient<T>(
  url: string,
  options: FetchOptions = {}
): Promise<T> {
  const response = await fetch(API_URL + url, {
    ...options,
    credentials: "include",
    headers: {
      "Content-Type": "application/json",
      ...options.headers,
    },
  });
  
  if (response.ok) {
    return response.json();
  }

  if (response.status == 401 && !options.retry) {
    const refreshed = await tryRefresh();

    if (refreshed) {
      return fetchClient<T>(url, { ...options, retry: true });
    }

    redirectToRoot();
    throw createApiError(response);
  }

  if (response.status === 403) {
    throw createApiError(response);
  }

  throw createApiError(response);
};