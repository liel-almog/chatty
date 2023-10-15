import useWebSocket from "react-use-websocket";
import {
  CreateMessageDto,
  Message,
  messageSchema,
} from "../../models/message.model";
import { useQuery, useQueryClient } from "@tanstack/react-query";
import { useCallback, useContext, useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import invariant from "invariant";
import { RoomService } from "../../services/room.service";
import { UsernameContext } from "../../context/Username.context";

export const useChat = () => {
  const { roomId } = useParams<{ roomId: string }>();
  invariant(roomId, "roomId is required");
  const WS_URL = `ws://localhost:8080/ws/chat/${roomId}`;

  const { username } = useContext(UsernameContext);
  const { sendJsonMessage, lastJsonMessage } = useWebSocket<Message>(WS_URL);
  const [newMessage, setNewMessage] = useState<string>("");
  const queryClient = useQueryClient();
  const query = useQuery({
    queryKey: [roomId, "messages"],
    queryFn: () => RoomService.getChatMessages(roomId),
  });

  useEffect(() => {
    if (lastJsonMessage !== null) {
      const message = messageSchema.parse(lastJsonMessage);
      queryClient.setQueryData<Message[]>([roomId, "messages"], (prev) => {
        return [...(prev ?? []), message];
      });
    }
  }, [lastJsonMessage]);

  // We do not use useCallback here because the wsUrl is not changing
  const handleSubmit = useCallback(
    (event: React.FormEvent<HTMLFormElement>) => {
      event.preventDefault();
      if (newMessage) {
        sendJsonMessage<CreateMessageDto>({
          content: newMessage,
          roomId: 1,
          username: username,
        });
        setNewMessage("");
      }
    },
    [newMessage, username]
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