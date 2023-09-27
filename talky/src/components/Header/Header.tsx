import classes from "./header.module.scss";

export interface HeaderProps {}

export const Header = () => {
  return (
    <header className={classes.container}>
      <Title />
      <UserSection />
    </header>
  );
};

const Title = () => <h1>Chatty - פרוייקט צא'ט</h1>;

const UserSection = () => {
  return (
    <div className={classes.user}>
      <button>התנתקות</button>
    </div>
  );
};
