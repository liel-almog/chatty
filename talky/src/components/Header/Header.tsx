import { useNavigate } from "react-router-dom";
import classes from "./header.module.scss";

export interface HeaderProps {}

export const Header = () => {
  const navigate = useNavigate();

  return (
    <header className={classes.container}>
      <h1 onClick={() => navigate("/")}>Chatty - פרוייקט צא'ט</h1>
    </header>
  );
};
