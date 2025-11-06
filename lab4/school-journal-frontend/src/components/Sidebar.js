import { Link } from "react-router-dom";
import "./Sidebar.css";

export default function Sidebar() {
  return (
    <div className="sidebar">
      <h2>Шкільний Журнал</h2>
      <nav>
        <Link to="/students">Учні</Link>
        <Link to="/teachers">Вчителі</Link>
        <Link to="/classes">Класи</Link>
        <Link to="/subjects">Предмети</Link>
        <Link to="/lessons">Уроки</Link>
        <Link to="/grades">Оцінки</Link>
      </nav>
    </div>
  );
}
