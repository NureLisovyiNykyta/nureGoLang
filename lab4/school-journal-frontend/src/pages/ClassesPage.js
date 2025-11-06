import DataTable from "../components/DataTable";

export default function ClassesPage() {
  const columns = [
    { label: "Назва класу", accessor: "name" },
    { label: "ID класного керівника", accessor: "teacher_id" },
  ];

  return <DataTable endpoint="classes" columns={columns} />;
}
