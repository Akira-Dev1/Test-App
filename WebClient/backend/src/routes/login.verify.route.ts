import { Router, Request, Response } from "express";
import { verifyWithAuth } from "../services/auth.service";
import { SessionRepository } from "../redis/session.repository";

const router = Router();

/**
 * POST /login/verify
 * Body: { code: string }
 * Cookie: session_id
 */
router.post("/login/verify", async (req: Request, res: Response) => {
  try {
    const sessionId = req.cookies?.session_id;
    const { code } = req.body;

    // Базовая валидация
    if (!sessionId || !code) {
      return res.status(400).json({ error: "Invalid request" });
    }

    // Получаем сессию из Redis
    const session = await SessionRepository.get(sessionId);

    if (!session) {
      return res.status(401).json({ status: "anonymous" });
    }

    // Проверяем, что мы реально в pending
    if (session.status !== "pending" || !session.entryToken) {
      return res.status(400).json({ error: "Session is not pending" });
    }

    // Вызываем Auth Module
    const authResponse = await verifyWithAuth(
      session.entryToken,
      code
    );

    // Обрабатываем ответы Auth
    if (authResponse.status === "pending") {
      return res.json({ status: "pending" });
    }

    if (authResponse.status === "access_denied") {
      // пользователь отказал
      await SessionRepository.delete(sessionId);
      return res.status(401).json({ status: "access_denied" });
    }

    // УСПЕХ → сохраняем JWT в Redis
    await SessionRepository.set(sessionId, {
      status: "authorized",
      accessToken: authResponse.access_token,
      refreshToken: authResponse.refresh_token,
      userId: authResponse.user_id,
    });

    // Сообщаем SPA об успехе
    return res.json({ status: "approved" });

  } catch (error) {
    console.error("LOGIN VERIFY ERROR:", error);
    return res.status(500).json({ error: "Internal server error" });
  }
});

export default router;
