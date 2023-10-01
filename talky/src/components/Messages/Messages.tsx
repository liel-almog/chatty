import clsx from "clsx";
import { Message } from "../../models/message.model";
import classes from "./messages.module.scss";

export interface MessagesProps {
  messages: Message[];
}

const MessageConatiner = ({ content, isMe }: Message) => {
  return (
    <article
      className={clsx(classes.messageContainer, { [classes.isMe]: isMe })}
    >
      <div className={classes.message}>
        <span className={classes.username}>A User</span>
        <p>{content}</p>
      </div>
    </article>
  );
};

export const Messages = ({ messages }: MessagesProps) => {
  const messagesElems = messages.map((message) => {
    return <MessageConatiner {...message} key={crypto.randomUUID()} />;
  });

  return <main className={classes.container}>{messagesElems}</main>;
};
