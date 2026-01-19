const AUTH_URL = process.env.AUTH_SERVICE_URL || "http://auth:8081";

type AuthVerifyResponse =
  | {
      status: "approved";
      access_token: string;
      refresh_token: string;
      user_id: string;
    }
  | {
      status: "access_denied";
    }
  | {
      status: "pending";
    };

export async function verifyWithAuth(
  entryToken: string,
  code: string
): Promise<AuthVerifyResponse> {
  const res = await fetch(`${AUTH_URL}/login/verify`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      entry_token: entryToken,
      code,
    }),
  });

  if (!res.ok) {
    throw new Error("Auth service error");
  }

  return res.json();
}
