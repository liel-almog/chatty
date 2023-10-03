import z from "zod";

export const messageSchema = z.object({
  content: z.string().min(1).max(500),
  isMe: z.boolean().default(false),
  roomId: z.number().int(),
});

export type Message = z.infer<typeof messageSchema>;

export const createMessageSchema = messageSchema.pick({
  content: true,
  roomId: true,
});

export type CreateMessageDto = z.infer<typeof createMessageSchema>;
