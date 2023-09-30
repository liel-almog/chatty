import { useState } from "react";
import { Message, messageSchema } from "../../models/message.model";
import classes from "./chat.module.scss";
import { Messages } from "../../components/Messages";
export interface ChatProps {}

export const Chat = () => {
  const ws = new WebSocket("ws://localhost:8080");
  const [messages, setMessages] = useState<Message[]>([]);

  ws.onmessage = (event) => {
    const message = messageSchema.parse(JSON.parse(event.data));
    setMessages((prev) => [...prev, message]);
  };

  return (
    <main className={classes.container}>
      <header className={classes.header}>
        <h1>חדר צא'ט</h1>
      </header>
      <Messages messages={messages} />
      <section className={classes.input}>
        <textarea
          id="newMessage"
          placeholder="כתוב את ההודעה שלך כאן..."
          rows={1}
        ></textarea>
        <button id="sendMessage">שליחה</button>
      </section>
    </main>
  );
};
