import DataTable from "../components/DataTable";

export default function StudentsPage() {
  const columns = [
    { label: "Ім’я", accessor: "first_name" },
    { label: "Прізвище", accessor: "last_name" },
    { label: "Email", accessor: "email" },
    { label: "ID класу", accessor: "class_id" },
    { label: "Дата народження", accessor: "birth_date" },
    { label: "Створено", accessor: "created_at" },
    { label: "Змінено", accessor: "updated_at" },
  ];

  return <DataTable endpoint="students" columns={columns} />;
}
