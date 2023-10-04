import { PropsWithChildren, createContext } from "react";
import { useLocalStorage } from "../hooks/useLocalStorage";

export interface UsernameContextProps {
  username: string;
  setUsername: React.Dispatch<React.SetStateAction<string>>;
}

export const UsernameContext = createContext<UsernameContextProps>({
  username: "",
  setUsername: () => {},
});

export interface UsernameProviderProps {}

export const UsernameProvider = ({
  children,
}: PropsWithChildren<UsernameProviderProps>) => {
  const [username, setUsername] = useLocalStorage<string>("username", "");
  console.log({ username });

  return (
    <UsernameContext.Provider value={{ username, setUsername }}>
      {children}
    </UsernameContext.Provider>
  );
};
