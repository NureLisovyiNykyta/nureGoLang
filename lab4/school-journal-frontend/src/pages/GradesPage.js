import DataTable from "../components/DataTable";

export default function GradesPage() {
  const columns = [
    { label: "ID учня", accessor: "student_id" },
    { label: "ID уроку", accessor: "lesson_id" },
    { label: "Оцінка", accessor: "grade" },
    { label: "Коментар", accessor: "comment" },
  ];

  return <DataTable endpoint="grades" columns={columns} />;
}
