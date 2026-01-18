// Состояние авторизации
export type AuthState =
  | "unknown"
  | "anonymous"
	| "need_verify"
  | "authorized";

// Ответ сервиса авторизации при проверках
export type AuthStatusResponse = {
	status: "pending" | "access_denied" | "approved";
};

// Ответ при успешной верификации
export type VerifyResponse = {
  status: "authorized";
};

// Универсальная ошибка API
export type ApiError = {
  statusCode: number;
  message?: string;
};