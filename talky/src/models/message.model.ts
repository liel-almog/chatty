import z from "zod";

export const messageSchema = z.object({
  content: z.string().min(1).max(500),
  isMe: z.boolean().default(false),
});

export type Message = z.infer<typeof messageSchema>;
