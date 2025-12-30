import { Navigate } from "react-router-dom";
import { isAuthenticated } from "../store/auth";
import type { JSX } from "react";

type Props = {
  children: JSX.Element;
};

const ProtectedRoute = ({children} : Props) => {
  if (!isAuthenticated()) {return <Navigate to="/login" replace/>}

  return children;
};

export default ProtectedRoute