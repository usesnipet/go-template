import { BrowserRouter, Route, Routes } from "react-router-dom";

import { AppLayout } from "./components/app-layout";
import { HomePage } from "./pages/page";
import { SettingsPage } from "./pages/settings/page";

export const Router = () => {
  return (
    <BrowserRouter>
      <Routes>
        <Route element={<AppLayout />}>
          <Route path="/" element={<HomePage />} />
          <Route path="/settings" element={<SettingsPage />} />
        </Route>
      </Routes>
    </BrowserRouter>
  )
}
