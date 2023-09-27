import { BrowserRouter, Route, Routes } from "react-router-dom";
import { Welcome } from "../pages/Welcome";
import { Layout } from "../components/Layout";

export interface PublicRoutesProps {}

export const PublicRoutes = () => (
  <BrowserRouter>
    <Routes>
      <Route path="/" element={<Layout />}>
        <Route index element={<Welcome />} />
      </Route>
    </Routes>
  </BrowserRouter>
);
