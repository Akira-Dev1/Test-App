import { Router, Request, Response } from "express";
import { SessionRepository } from "../redis/session.repository";
import { verifyWithAuth } from "../services/auth.service";
import { SESSION_COOKIE_NAME } from "../config/cookies";

const router = Router();

router.post("/login/verify", async (req: Request, res: Response) => {
  const sessionId = req.cookies?.[SESSION_COOKIE_NAME];
  const { code } = req.body;

  // Проверки входных данных
  if (!sessionId || !code) {
    return res.status(400).json({ error: "Invalid request" });
  }

  const session = await SessionRepository.get(sessionId);

  // Проверка сессии
  if (!session || session.status !== "pending" || !session.entryToken) {
    return res.status(401).json({ error: "Session not valid" });
  }

  // Проверка через Auth Module
  const authResponse = await verifyWithAuth(
    session.entryToken,
    code
  );

  // Обработка ответов Auth
  if (authResponse.status === "pending") {
    return res.status(200).json({ status: "pending" });
  }

  if (authResponse.status === "access_denied") {
    await SessionRepository.delete(sessionId);
    return res.status(401).json({ status: "access_denied" });
  }

  // УСПЕХ → authorized
  await SessionRepository.set(sessionId, {
    status: "authorized",
    accessToken: authResponse.access_token,
    refreshToken: authResponse.refresh_token,
    userId: authResponse.user_id,
  });

  return res.status(200).json({ status: "approved" });
});

export default router;
