const API_URL = import.meta.env.VITE_API_URL;

export async function tryRefresh(): Promise<boolean> {
  const res = await fetch(`${API_URL}/auth/refresh`, {
    method: "POST",
    credentials: "include",
  });

  return res.ok;
}