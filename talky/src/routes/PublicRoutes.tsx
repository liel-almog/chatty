import { BrowserRouter, Route, Routes } from "react-router-dom";
import { Welcome } from "../pages/Welcome";

export interface PublicRoutesProps {}

export const PublicRoutes = () => (
  <BrowserRouter>
    <Routes>
      <Route index element={<Welcome />} />
    </Routes>
  </BrowserRouter>
);
