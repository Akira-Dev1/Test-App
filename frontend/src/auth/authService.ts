import { fetchClient } from "../api/fetchClient";
import type { AuthStatusResponse } from "./types";

export async function fetchAuthStatus(): Promise<AuthStatusResponse> {
  return fetchClient<AuthStatusResponse>("/login/check");
}

export async function verifyLogin(code: string): Promise<void> {
  await fetchClient("/login/verify", {
    method: "POST",
    body: JSON.stringify({ code }),
  });
}

export async function logout(): Promise<void> {
  await fetchClient("/logout", {
    method: "POST",
  });
}

export async function logoutAll(): Promise<void> {
  await fetchClient("/logout?all=true", {
    method: "POST",
  });
}
