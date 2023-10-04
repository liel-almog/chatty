import classes from "./chat.module.scss";
import { Messages } from "../../components/Messages";
import { useChat } from "./useChat";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faBolt } from "@fortawesome/free-solid-svg-icons";
export interface ChatProps {}

export const Chat = () => {
  const { handleNewMessageChange, handleSubmit, newMessage, query } = useChat();
  const messages = query.data ?? [];

  return (
    <main className={classes.container}>
      <header className={classes.header}>
        <h2>חדר צא'ט</h2>
      </header>
      {messages.length > 0 && <Messages messages={messages} />}
      <form className={classes.send} onSubmit={handleSubmit}>
        <input
          onChange={handleNewMessageChange}
          value={newMessage}
          id="newMessage"
          placeholder="כתוב את ההודעה שלך כאן..."
        ></input>
        <button type="submit" id="sendMessage">
          <FontAwesomeIcon icon={faBolt} />
          <span>שלח</span>
        </button>
      </form>
    </main>
  );
};
