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
      <article key={room.id} className={classes.item}>
        <span>{room.name}</span>
        <button onClick={() => navigate(`${room.id}/chat`)}>הצטרף לחדר</button>
      </article>
    );
  });

  return (
    <main className={classes.container}>
      <header>
        <h1>Chat Rooms</h1>
      </header>
      <section className={classes.list}>{roomsCard}</section>
    </main>
  );
};
