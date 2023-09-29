import classes from "./chat.module.scss";

export interface ChatProps {}

export const Chat = () => {
  return (
    <main className={classes.container}>
      <header className={classes.header}>
        <h1>חדר צא'ט</h1>
      </header>
      <section className={classes.messages}>
        <article className={classes.message}>
          <span className={classes.username}>ליאל:</span> שלום, מה שלומך?
        </article>
      </section>
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
