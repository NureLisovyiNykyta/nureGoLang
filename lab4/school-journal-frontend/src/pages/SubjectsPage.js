import DataTable from "../components/DataTable";

export default function SubjectsPage() {
  const columns = [
    { label: "Назва предмета", accessor: "name" },
    { label: "Опис", accessor: "description" },
  ];

  return <DataTable endpoint="subjects" columns={columns} />;
}
