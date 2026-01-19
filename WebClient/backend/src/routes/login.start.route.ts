import { Router, Request, Response } from "express";
import { SessionRepository } from "../redis/session.repository";
import { generateToken } from "../utils/token";
import { SESSION_COOKIE_NAME } from "../config/cookies";

const router = Router();

router.get("/login", async (req: Request, res: Response) => {
  const type = req.query.type as string | undefined;
  
  if (!type) {
  return res.sendStatus(200); // SPA через nginx
  }

  // Проверка типа логина
  if (!type || !["github", "yandex", "code"].includes(type)) {
    return res.status(400).send("Invalid login type");
  }

  // Получаем session_id
  let sessionId = req.cookies?.[SESSION_COOKIE_NAME];

  // Если нет сессии — создаём
  if (!sessionId) {
    sessionId = generateToken(16);

    res.cookie(SESSION_COOKIE_NAME, sessionId, {
      httpOnly: true,
      sameSite: "lax",
      path: "/", 
    });
  }

  // Генерируем entry_token
  const entryToken = generateToken(16);

  // Сохраняем pending-сессию
  await SessionRepository.set(sessionId, {
    status: "pending",
    entryToken,
  });

  /**
   * ⚠️ ВАЖНО ПО ТЗ
   * Здесь должен быть вызов Authorization Module
   * Пока делаем заглушку
   */

  if (type === "code") {
    console.log("Login code:", {
      sessionId,
      entryToken,
    });
  }

  // Redirect пользователя
  return res.redirect("/verify");
});

export default router;
