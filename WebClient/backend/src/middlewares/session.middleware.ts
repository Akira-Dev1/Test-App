import { Request, Response, NextFunction } from "express";

export function sessionMiddleware(
  req: Request,
  _: Response,
  next: NextFunction
) {
  const sessionId = req.cookies?.session_id;

  (req as any).sessionId = sessionId || null;

  next();
}