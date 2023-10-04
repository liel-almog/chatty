import { messageSchema } from "../models/message.model";
import { roomSchema } from "../models/room.model";
import { axiosInstance } from "./utils/axios";

const PREFIX = "room";
export class RoomService {
  static async getRooms() {
    const rawData = (await axiosInstance.get(`${PREFIX}`)).data;
    const rooms = roomSchema.array().parse(rawData);
    return rooms;
  }

  static async getChatMessages(roomId: string) {
    const rawData = (await axiosInstance.get(`${PREFIX}/${roomId}/message`))
      .data;
    const messages = messageSchema.array().parse(rawData);
    return messages;
  }
}
