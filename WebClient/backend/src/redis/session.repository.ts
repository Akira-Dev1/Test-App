import { redisClient } from "./redis.client";
import { SessionData } from "../types/session";

const SESSION_TTL = Number(process.env.SESSION_TTL || 600);

export class SessionRepository {
  static async get(sessionId: string): Promise<SessionData | null> {
    const raw = await redisClient.get(sessionId);
    return raw ? JSON.parse(raw) : null;
  }

  static async set(
    sessionId: string,
    data: SessionData
  ): Promise<void> {
    await redisClient.set(sessionId, JSON.stringify(data), {
      EX: SESSION_TTL,
    });
  }

  static async delete(sessionId: string): Promise<void> {
    await redisClient.del(sessionId);
  }
}
