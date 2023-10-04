import { z } from "zod";

export const roomSchema = z.object({
  name: z.string().min(1).max(2000),
  id: z.number().int(),
});

export type Room = z.infer<typeof roomSchema>;
