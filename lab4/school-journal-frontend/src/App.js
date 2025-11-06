import { BrowserRouter as Router, Routes, Route, Navigate } from "react-router-dom";
import Sidebar from "./components/Sidebar";
import StudentsPage from "./pages/StudentsPage";
import TeachersPage from "./pages/TeachersPage";
import ClassesPage from "./pages/ClassesPage";
import SubjectsPage from "./pages/SubjectsPage";
import LessonsPage from "./pages/LessonsPage";
import GradesPage from "./pages/GradesPage";
import "./styles.css";

function App() {
  return (
    <Router>
      <div className="app">
        <Sidebar />
        <div className="content">
          <Routes>
            <Route path="/" element={<Navigate to="/students" />} />
            <Route path="/students" element={<StudentsPage />} />
            <Route path="/teachers" element={<TeachersPage />} />
            <Route path="/classes" element={<ClassesPage />} />
            <Route path="/subjects" element={<SubjectsPage />} />
            <Route path="/lessons" element={<LessonsPage />} />
            <Route path="/grades" element={<GradesPage />} />
          </Routes>
        </div>
      </div>
    </Router>
  );
}

export default App;
