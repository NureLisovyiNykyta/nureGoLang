import DataTable from "../components/DataTable";

export default function TeachersPage() {
  const columns = [
    { label: "Ім’я", accessor: "first_name" },
    { label: "Прізвище", accessor: "last_name" },
    { label: "Email", accessor: "email" },
    { label: "ID предмета", accessor: "subject_id" },
  ];

  return <DataTable endpoint="teachers" columns={columns} />;
}
