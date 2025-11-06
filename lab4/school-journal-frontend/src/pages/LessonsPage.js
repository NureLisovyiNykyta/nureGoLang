import DataTable from "../components/DataTable";

export default function LessonsPage() {
  const columns = [
    { label: "ID класу", accessor: "class_id" },
    { label: "ID предмета", accessor: "subject_id" },
    { label: "ID вчителя", accessor: "teacher_id" },
    { label: "Дата", accessor: "date" },
    { label: "Тема уроку", accessor: "topic" },
    { label: "Домашнє завдання", accessor: "homework" },
  ];

  return <DataTable endpoint="lessons" columns={columns} />;
}
