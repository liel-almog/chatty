import { useContext, useState } from "react";
import classes from "./welcome.module.scss";
import { UsernameContext } from "../../context/Username.context";
import { useNavigate } from "react-router-dom";

export interface WelcomeProps {}

export const Welcome = () => {
  const { username, setUsername } = useContext(UsernameContext);
  const [name, setName] = useState<string>(username);
  const [isNameValid, setIsNameValid] = useState<boolean>(true);
  const navigate = useNavigate();
  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    if (name && name.length <= 50) {
      setName("");
      setUsername(name);
      navigate("/room");
    } else {
      setIsNameValid(false);
    }
  };

  return (
    <main className={classes.container}>
      <h1>ברוכים הבאים ל- Chatty ! 🎉</h1>
      <section className={classes.instructions}>
        <p>
          נשמע שאתם כאן לשוחח, נכון? אז בואו נתחיל בצעד קטן - בחירת שם משתמש.
        </p>
        <p>
          שם המשתמש שלכם יהיה הזהות שלכם בצ'אט וילווה אתכם בכל פעילותכם כאן.
          <br />
          אנחנו רק רוצים להזכיר שהמערכת רצה לוקאלית ואיננה מאובטחת, אז נא להימנע
          משימוש במידע רגיש 😁
        </p>
        <article>
          <strong>הערות:</strong>
          <ul>
            <li>השם אינו יכול להכיל יותר מ-50 תווים</li>
            <li>השם חייב להיות מגניב וייחודי</li>
          </ul>
        </article>
      </section>
      <form onSubmit={handleSubmit} className={classes.getStarted}>
        <input
          onChange={(e) => {
            const value = e.target.value;
            if (value && value.length <= 50) {
              setIsNameValid(true);
            } else {
              setIsNameValid(false);
            }

            setName(value);
          }}
          value={name}
          placeholder="בחרו שם משתמש"
          className={classes.username}
        />
        {!isNameValid && <p className={classes.error}>שם המשתמש אינו תקין</p>}
        <br />
        <button className={classes.btn}>בואו נתחיל</button>
      </form>
    </main>
  );
};
