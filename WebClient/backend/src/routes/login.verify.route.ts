import { Router, Request, Response } from "express";
import { verifyWithAuth } from "../services/auth.service";
import { SessionRepository } from "../redis/session.repository";

const router = Router();

router.post("/login/verify", async (req: Request, res: Response) => {
  try {
    const sessionId = req.cookies?.session_id;
    const { code } = req.body;

    if (!sessionId || !code) {
      return res.status(400).json({ error: "Invalid request" });
    }

    const session = await SessionRepository.get(sessionId);

    if (!session) {
      return res.status(401).json({ status: "anonymous" });
    }

    if (session.status !== "pending" || !session.entryToken) {
      return res.status(400).json({ error: "Session is not pending" });
    }

    const authResponse = await verifyWithAuth(
      session.entryToken,
      code
    );

    if (authResponse.status === "pending") {
      return res.json({ status: "pending" });
    }

    if (authResponse.status === "access_denied") {

      await SessionRepository.delete(sessionId);
      return res.status(401).json({ status: "access_denied" });
    }

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
