import { useEffect, useState } from "react";
import {
  CreateMessageDto,
  Message,
  messageSchema,
} from "../../models/message.model";
import classes from "./chat.module.scss";
import { Messages } from "../../components/Messages";
import useWebSocket from "react-use-websocket";
import { useQuery, useQueryClient } from "@tanstack/react-query";
import { ChatService } from "../../services/chat.service";
export interface ChatProps {}

export const Chat = () => {
  const queryClient = useQueryClient();
  const [newMessage, setNewMessage] = useState<string>("");
  const WS_URL = "ws://localhost:8080/ws/chat";
  const { sendJsonMessage, lastJsonMessage } = useWebSocket<Message>(WS_URL);
  const query = useQuery({
    queryKey: ["messages"],
    queryFn: ChatService.getChatMessages,
  });

  useEffect(() => {
    if (lastJsonMessage !== null) {
      const message = messageSchema.parse(lastJsonMessage);
      queryClient.setQueryData<Message[]>(["messages"], (prev) => {
        return [...(prev ?? []), message];
      });
    }
  }, [lastJsonMessage]);

  // We do not use useCallback here because the wsUrl is not changing
  const handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    if (newMessage) {
      sendJsonMessage<CreateMessageDto>({ content: newMessage, roomId: 1 });
      setNewMessage("");
    }
  };

  const handleNewMessageChange = (
    event: React.ChangeEvent<HTMLInputElement>
  ) => {
    setNewMessage(event.target.value);
  };

  return (
    <main className={classes.container}>
      <header className={classes.header}>
        <h1>חדר צא'ט</h1>
      </header>
      <Messages messages={query.data ?? []} />
      <form onSubmit={handleSubmit}>
        <section className={classes.input}>
          <input
            onChange={handleNewMessageChange}
            value={newMessage}
            id="newMessage"
            placeholder="כתוב את ההודעה שלך כאן..."
          ></input>
          <button type="submit" id="sendMessage">
            שליחה
          </button>
        </section>
      </form>
    </main>
  );
};
