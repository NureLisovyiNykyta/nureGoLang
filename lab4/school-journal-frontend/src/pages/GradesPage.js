import DataTable from "../components/DataTable";

export default function GradesPage() {
  const columns = [
    { label: "ID учня", accessor: "student_id" },
    { label: "ID уроку", accessor: "lesson_id" },
    { label: "Оцінка", accessor: "value" },
    { label: "Коментар", accessor: "comment" },
    { label: "Створено", accessor: "created_at" },
    { label: "Змінено", accessor: "updated_at" },
  ];

  return <DataTable endpoint="grades" columns={columns} />;
}
