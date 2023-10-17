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

export interface PublicRoutesProps {}

export const routes = createBrowserRouter(
  createRoutesFromElements(
    <Route path="*" element={<Layout />}>
      <Route index element={<Welcome />} />
      <Route path="room" element={<Rooms />}></Route>
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
