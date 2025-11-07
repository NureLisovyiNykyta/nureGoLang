import DataTable from "../components/DataTable";

export default function SubjectsPage() {
  const columns = [
    { label: "Назва предмета", accessor: "name" },
    { label: "Опис", accessor: "description" },
    { label: "Створено", accessor: "created_at" },
    { label: "Змінено", accessor: "updated_at" },
  ];

  return <DataTable endpoint="subjects" columns={columns} />;
}
