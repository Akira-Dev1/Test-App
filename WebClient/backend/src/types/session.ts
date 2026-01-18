export type SessionStatus =
  | "anonymous"
  | "pending"
  | "authorized";

export type SessionData = {
  status: SessionStatus;

  entryToken?: string;

  accessToken?: string;
  refreshToken?: string;

  userId?: string;
};