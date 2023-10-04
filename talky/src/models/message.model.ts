import z from "zod";

export const messageSchema = z.object({
  content: z.string().min(1).max(500),
  roomId: z.number().int(),
  username: z.string().min(1).max(50),
  createdAt: z.coerce.date(),
  id: z.number().int(),
});

export type Message = z.infer<typeof messageSchema>;

export const createMessageSchema = messageSchema.pick({
  content: true,
  roomId: true,
  username: true,
});

export type CreateMessageDto = z.infer<typeof createMessageSchema>;
