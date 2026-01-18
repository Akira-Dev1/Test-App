import { Navigate } from "react-router-dom";
import type { ReactElement } from "react";
import { useAuth } from "../hooks/useAuth";

type PublicRouteProps = {
  children: ReactElement;
};

const PublicRoute = ({ children }: PublicRouteProps) => {
  const { state } = useAuth();

  if (state === "unknown") {
    return null;
  }

  if (state === "authorized") {
    return <Navigate to="/" replace />;
  }

  return children;
};

export default PublicRoute;