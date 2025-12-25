import { api } from "./http";

export const getHealth = async () => {
  const { data } = await api.get("/health");
  return data;
};
