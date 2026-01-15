import React, { useEffect, useState } from "react";
import type { AuthState } from "./types";
import { fetchAuthStatus } from "./authService";
import { redirectToVerify, redirectToRoot } from "./redirect";
import { AuthContext } from "./AuthContext";
import Loader from "../components/Loader";

export const AuthProvider: React.FC<{ children: React.ReactNode }> = ({ children }) => {
  const [state, setState] = useState<AuthState>("unknown");

  const refreshAuth = async () => {
    try {
      const data = await fetchAuthStatus();

      switch (data.status) {
        case "pending":
          setState("need_verify");
          redirectToVerify();
          break;
        case "access_denied":
          setState("anonymous");
          redirectToRoot();
          break;
        case "approved":
          setState("authorized");
          break;
        default:
          setState("unknown");
      }
    } catch (e) {
      const err = e as { statusCode?: number };
      if (err.statusCode === 401) {
        setState("anonymous");
        redirectToRoot();
      }
    }
  };

  useEffect(() => {
    (async () => {
      await refreshAuth();
    })()
  }, []);

  if (state === "unknown") {
    return <Loader />;
  }

  return (
    <AuthContext.Provider value={{ state, refreshAuth }}>
      {children}
    </AuthContext.Provider>
  );
};
