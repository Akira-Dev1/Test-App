import { createContext } from "react";
import type { AuthState } from "./types";

export type AuthContextValue = {
  state: AuthState;
  refreshAuth: () => Promise<void>;
};

export const AuthContext = createContext<AuthContextValue | null>(null);
