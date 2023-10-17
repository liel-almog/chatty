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
      <Route path="room">
        <Route index element={<Rooms />} />
        <Route path=":roomId/chat" element={<Chat />} />
      </Route>
    </Route>
  )
);

export const PublicRoutes = () => {
  return (
    <BaseErrorBoundary>
      <RouterProvider router={routes} />
    </BaseErrorBoundary>
  );
};
