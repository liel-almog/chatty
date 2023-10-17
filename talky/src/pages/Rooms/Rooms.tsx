import { RoomService } from "../../services/room.service";
import classes from "./rooms.module.scss";
import { useQuery } from "@tanstack/react-query";
import { Chat } from "../Chat";
import { useState } from "react";
import clsx from "clsx";

export interface RoomsProps {}

export const Rooms = () => {
  const [roomId, setRoomId] = useState<number>();

  const query = useQuery({
    queryKey: ["rooms"],
    queryFn: RoomService.getRooms,
  });

  const roomsCard = query.data?.map((room) => {
    return (
      <li
        key={room.id}
        className={clsx(classes.item, {
          [classes.active]: roomId === room.id,
        })}
        onClick={() => setRoomId(room.id)}
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
      {roomId && <Chat roomId={roomId} />}
    </>
  );
};
