import DataTable from "../components/Datatable";

export default function StudentsPage() {
  const columns = [
    { label: "Ім’я", accessor: "first_name" },
    { label: "Прізвище", accessor: "last_name" },
    { label: "Email", accessor: "email" },
    { label: "Дата народження", accessor: "birth_date" },
  ];

  return <DataTable endpoint="students" columns={columns} />;
}
