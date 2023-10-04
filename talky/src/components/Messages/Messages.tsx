import clsx from "clsx";
import { Message } from "../../models/message.model";
import classes from "./messages.module.scss";
import { useEffect, useRef } from "react";

export interface MessagesProps {
  messages: Message[];
}

const MessageConatiner = ({ content, isMe }: Message) => {
  return (
    <li className={clsx(classes.messageContainer, { [classes.isMe]: isMe })}>
      <div className={classes.message}>
        <h6 className={classes.username}>A User</h6>
        <span>{content}</span>
      </div>
    </li>
  );
};

export const Messages = ({ messages }: MessagesProps) => {
  const scrollRef = useRef<HTMLDivElement>(null);
  const messagesElems = messages.map((message) => {
    return <MessageConatiner {...message} key={message.id} />;
  });

  useEffect(() => {
    if (scrollRef) {
      scrollRef.current?.scrollIntoView({ behavior: "smooth" });
    }
  }, []);

  return (
    <ul className={classes.container}>
      {messagesElems}
      <div ref={scrollRef} />
    </ul>
  );
};
