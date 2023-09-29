import { BrowserRouter, Route, Routes } from "react-router-dom";
import { Welcome } from "../pages/Welcome";
import { Layout } from "../components/Layout";
import { Chat } from "../pages/Chat";

export interface PublicRoutesProps {}

export const PublicRoutes = () => (
  <BrowserRouter>
    <Routes>
      <Route path="/" element={<Layout />}>
        <Route index element={<Welcome />} />
        <Route path="/chat" element={<Chat/>} />
      </Route>
    </Routes>
  </BrowserRouter>
);
