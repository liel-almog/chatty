import {
  Route,
  RouterProvider,
  createBrowserRouter,
  createRoutesFromElements,
} from "react-router-dom";
import { Layout } from "../pages/Layout";
import { Rooms } from "../pages/Rooms";
import { Welcome } from "../pages/Welcome";
import { BaseErrorBoundary } from "../components/common/BaseErrorBoundary";
import { Chat } from "../pages/Chat";

export interface PublicRoutesProps {}

export const routes = createBrowserRouter(
  createRoutesFromElements(
    <Route path="*" element={<Layout />}>
      <Route index element={<Welcome />} />
      <Route
        // errorElement={<BaseErrorBoundary />}
        // ErrorBoundary={BaseErrorBoundary}
        // hasErrorBoundary
        path="room"
        element={
          <BaseErrorBoundary>
            <Rooms />
          </BaseErrorBoundary>
        }
      >
        <Route
          path=":roomId/chat"
          handle={{ scrollMode: "pathname" }}
          element={
            <BaseErrorBoundary>
              <Chat />
            </BaseErrorBoundary>
          }
        />
      </Route>
    </Route>
  )
);

export const PublicRoutes = () => {
  return <RouterProvider router={routes} />;
};
