import DataTable from "../components/DataTable";

export default function ClassesPage() {
  const columns = [
    { label: "Назва класу", accessor: "name" },
    { label: "ID класного керівника", accessor: "teacher_id" },    
    { label: "Створено", accessor: "created_at" },
    { label: "Змінено", accessor: "updated_at" },
  ];

  return <DataTable endpoint="classes" columns={columns} />;
}
