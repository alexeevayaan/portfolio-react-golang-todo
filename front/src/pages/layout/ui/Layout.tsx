import { NavLink, Outlet } from "react-router";
import "./layout.css";

export function Layout() {
  return (
    <div className="wrapper">
      <header className="header">
        <NavLink to={"/"}>
          <span className="header-item">Главная</span>
        </NavLink>
        <NavLink to={"/"}>
          <span className="header-item">Войти</span>
        </NavLink>
      </header>
      <main>
        <Outlet />
      </main>
      <footer>Footer</footer>
    </div>
  );
}
