import { GlobalProvider } from "./context/GlobalProvider";
import { PublicRoutes } from "./routes/PublicRoutes";
import "./index.scss";

function App() {
  return (
    <GlobalProvider>
      <PublicRoutes />
    </GlobalProvider>
  );
}

export default App;
