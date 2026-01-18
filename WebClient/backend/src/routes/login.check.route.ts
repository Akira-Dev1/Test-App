import { Router, Request, Response } from "express";
import { SessionRepository } from "../redis/session.repository";

const router = Router();

router.get("/login/check", async (req: Request, res: Response) => {
  const sessionId = (req as any).sessionId as string | null;

  // Нет cookie → anonymous
  if (!sessionId) {
    return res.status(200).json({ status: "anonymous" });
  }

  // Проверяем Redis
  const session = await SessionRepository.get(sessionId);

  // Нет записи → anonymous
  if (!session) {
    return res.status(200).json({ status: "anonymous" });
  }

  // Возвращаем статус из Redis
  return res.status(200).json({
    status: session.status,
  });
});

export default router;
