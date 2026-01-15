const API_URL = "http://localhost:8081";

export async function tryRefresh(): Promise<boolean> {
  const res = await fetch(`${API_URL}/auth/refresh`, {
    method: "POST",
    credentials: "include",
  });

  return res.ok;
}