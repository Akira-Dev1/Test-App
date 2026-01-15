import type { ApiError } from "../auth/types";

export function createApiError(response: Response): ApiError {
  return {
    statusCode: response.status,
  };
}
