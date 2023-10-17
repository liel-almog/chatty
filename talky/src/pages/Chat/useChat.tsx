import useWebSocket from "react-use-websocket";
import {
  CreateMessageDto,
  Message,
  createMessageSchema,
  messageSchema,
} from "../../models/message.model";
import { useQuery, useQueryClient } from "@tanstack/react-query";
import { useCallback, useContext, useEffect, useState } from "react";
import { RoomService } from "../../services/room.service";
import { UsernameContext } from "../../context/Username.context";

interface UseChatProps {
  roomId: number;
}

export const useChat = ({ roomId }: UseChatProps) => {
  const WS_URL = `ws://localhost:8080/ws/chat/${roomId}`;
  const { username } = useContext(UsernameContext);
  const { sendJsonMessage, lastJsonMessage } = useWebSocket<Message>(WS_URL);
  const [newMessage, setNewMessage] = useState<string>("");
  const queryClient = useQueryClient();
  const query = useQuery({
    queryKey: [roomId, "messages"],
    queryFn: () => RoomService.getChatMessages(roomId.toString()),
  });

  useEffect(() => {
    if (lastJsonMessage !== null) {
      const message = messageSchema.parse(lastJsonMessage);
      queryClient.setQueryData<Message[]>([roomId, "messages"], (prev) => {
        return [...(prev ?? []), message];
      });
    }
  }, [lastJsonMessage]);

  const handleSubmit = useCallback(
    (event: React.FormEvent<HTMLFormElement>) => {
      event.preventDefault();
      if (newMessage) {
        const rawMessage: CreateMessageDto = {
          content: newMessage,
          roomId: roomId,
          username: username,
        };

        const result = createMessageSchema.safeParse(rawMessage);
        if (!result.success) {
          throw new Error("הודעה לא תקינה");
        }

        sendJsonMessage<CreateMessageDto>(result.data);

        setNewMessage("");
      }
    },
    [newMessage, roomId, username]
  );

  const handleNewMessageChange = (
    event: React.ChangeEvent<HTMLInputElement>
  ) => {
    setNewMessage(event.target.value);
  };

  return {
    handleSubmit,
    handleNewMessageChange,
    newMessage,
    query,
  };
};
