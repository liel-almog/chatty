import { RoomService } from "../../services/room.service";
import classes from "./rooms.module.scss";
import { useQuery } from "@tanstack/react-query";
import clsx from "clsx";
import { Outlet, useNavigate, useParams } from "react-router-dom";
import { z } from "zod";

export interface RoomsProps {}

export const Rooms = () => {
  const navigate = useNavigate();
  const { roomId } = useParams();
  const parsedRoomId = z.coerce.number().optional().safeParse(roomId);
  if (!parsedRoomId.success) {
    throw new Error("מזהה חדר לא תקין");
  }

  const query = useQuery({
    queryKey: ["rooms"],
    queryFn: RoomService.getRooms,
  });

  const roomsCard = query.data?.map((room) => {
    return (
      <li
        key={room.id}
        className={clsx(classes.item, {
          [classes.active]: parsedRoomId.data === room.id,
        })}
        onClick={() => navigate(`${room.id}/chat`)}
      >
        <span className={classes.textContent}>{room.name}</span>
      </li>
    );
  });

  return (
    <>
      <aside className={classes.roomsAside}>
        <header className={classes.header}>
          <h2>רשימת חדרים</h2>
        </header>
        <ul className={classes.list}>{roomsCard}</ul>
      </aside>
      <Outlet />
    </>
  );
};
