import { BrowserRouter, Navigate, Route, Routes } from "react-router-dom";
import { Layout } from "../components/Layout";
import { Chat } from "../pages/Chat";
import { Rooms } from "../pages/Rooms";
import { Welcome } from "../pages/Welcome";
import { BaseErrorBoundary } from "../components/common/BaseErrorBoundary";

export interface PublicRoutesProps {}

export const PublicRoutes = () => (
  <BrowserRouter>
    <BaseErrorBoundary>
      <Routes>
        <Route path="/" element={<Layout />}>
          <Route index element={<Welcome />} />
          <Route path="room">
            <Route index element={<Rooms />} />

            <Route
              path=":roomId/chat"
              element={
                <BaseErrorBoundary>
                  <Chat />
                </BaseErrorBoundary>
              }
            />
          </Route>
        </Route>
        <Route path="*" element={<Navigate to="/" />} />
      </Routes>
    </BaseErrorBoundary>
  </BrowserRouter>
);
