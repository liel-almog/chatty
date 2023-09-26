import classes from "./welcome.module.scss";

export interface WelcomeProps {}

export const Welcome = () => {
  return <div className={classes.container}>Welcome Home</div>;
};
