import type { FC } from "react";
import { BrowserRouter, Route, Routes } from "react-router";
import { HomePage } from "../../pages/home";
import { Layout } from "../../pages/layout";

export const AppRouter: FC = () => {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Layout />}>
          <Route index element={<HomePage />} />
        </Route>
      </Routes>
    </BrowserRouter>
  );
};
