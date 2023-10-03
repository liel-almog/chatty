import { messageSchema } from "../models/message.model";
import { axiosInstance } from "./utils/axios";

const PREFIX = "chat";
export class ChatService {
  static async getChatMessages() {
    const rawData = (await axiosInstance.get(`${PREFIX}/message`)).data;
    const messages = messageSchema.array().parse(rawData);
    return messages;
  }
}
