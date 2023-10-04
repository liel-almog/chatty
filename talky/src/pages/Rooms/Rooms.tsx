import { useNavigate } from "react-router-dom";
import { RoomService } from "../../services/room.service";
import classes from "./rooms.module.scss";
import { useQuery } from "@tanstack/react-query";

export interface RoomsProps {}

export const Rooms = () => {
  const navigate = useNavigate();
  const query = useQuery({
    queryKey: ["rooms"],
    queryFn: RoomService.getRooms,
  });

  const roomsCard = query.data?.map((room) => {
    return (
      <li key={room.id} className={classes.item}>
        <span className={classes.textContent}>{room.name}</span>
        <button onClick={() => navigate(`${room.id}/chat`)}>הצטרף לחדר</button>
      </li>
    );
  });

  return (
    <main className={classes.container}>
      <header className={classes.header}>
        <h2>רשימת חדרים</h2>
      </header>
      <ul className={classes.list}>{roomsCard}</ul>
    </main>
  );
};
