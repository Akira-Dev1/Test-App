import { Navigate } from "react-router-dom";
import type { JSX } from "react";
import { useAuth } from "../hooks/useAuth";

type Props = {
  children: JSX.Element;
};


function ProtectedRoute({ children }: Props) {
  const { state } = useAuth();

  if (state === "unknown") {
    return null;
  }

  if (state === "anonymous") {
    return <Navigate to="/" replace />;
  }

  if (state === "need_verify") {
    return <Navigate to="/verify" replace />;
  }

  return children;
}

export default ProtectedRoute