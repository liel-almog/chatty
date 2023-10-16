import { PropsWithChildren, createContext } from "react";
import { useLocalStorage } from "../hooks/useLocalStorage";

export interface UsernameContextProps {
  username: string;
  setUsername: (username: string) => void;
}

export const UsernameContext = createContext<UsernameContextProps>({
  username: "Anonymous",
  setUsername: () => {},
});

export interface UsernameProviderProps {}

export const UsernameProvider = ({
  children,
}: PropsWithChildren<UsernameProviderProps>) => {
  const [username, setUsername] = useLocalStorage<string>(
    "username",
    "Anonymous"
  );

  const handleSetUsername = (username: string) => {
    if (!username) {
      setUsername("Anonymous");
    }

    setUsername(username);
  };

  return (
    <UsernameContext.Provider
      value={{
        username: username,
        setUsername: handleSetUsername,
      }}
    >
      {children}
    </UsernameContext.Provider>
  );
};
