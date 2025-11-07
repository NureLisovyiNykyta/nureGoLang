import { useState, useEffect } from "react";
import { fetchData, createData, updateData, deleteData } from "../api";

export default function DataTable({ endpoint, columns }) {
  const [data, setData] = useState([]);
  const [formData, setFormData] = useState({});
  const [editingId, setEditingId] = useState(null);

  useEffect(() => {
    fetchData(endpoint).then(setData).catch(console.error);
  }, [endpoint]);

  const getInputType = (accessor) => {
    if (!accessor) return "text";
    const a = accessor.toLowerCase();
    if (a === "email" || a.endsWith("_email")) return "email";
    if (a === "birth_date" || a === "birthdate" || a === "date" || a.endsWith("_date")) return "date";
    if (a === "created_at" || a === "updated_at" || a.endsWith("_at")) return "datetime-local";
    if (a === "value" || a === "score" || a === "grade") return "number";
    return "text";
  };

  const visibleFormColumns = columns.filter(
    (c) => c && !["id", "created_at", "updated_at"].includes(c.accessor)
  );

  const handleChange = (e) => {
    setFormData({ ...formData, [e.target.name]: e.target.value });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      if (editingId) {
        await updateData(endpoint, editingId, formData);
      } else {
        await createData(endpoint, formData);
      }
      const updated = await fetchData(endpoint);
      setData(updated);
      setFormData({});
      setEditingId(null);
    } catch (err) {
      alert(err.message);
    }
  };

  const handleEdit = (item) => {
    setEditingId(item.id);
    setFormData(item);
  };

  const handleDelete = async (id) => {
    if (!window.confirm("Видалити запис?")) return;
    await deleteData(endpoint, id);
    setData(await fetchData(endpoint));
  };

  const formatValue = (val, accessor) => {
    if (val == null) return "";

    const key = (accessor || "").toLowerCase();

    const looksLikeDateKey = key.endsWith("_at") || key.includes("_date") || key === "date";

    if (looksLikeDateKey) {
      if (typeof val === "string" && /\d{4}-\d{2}-\d{2}T/.test(val)) {
        const d = new Date(val);
        if (!isNaN(d)) return d.toLocaleString("uk-UA");
      }

      if (typeof val === "number") {
        const d = new Date(val);
        if (!isNaN(d)) return d.toLocaleString("uk-UA");
      }
    }

    return val;
  };

  return (
    <div className="table-page">
      <h2>{endpoint.toUpperCase()}</h2>

      <form className="form" onSubmit={handleSubmit}>
        {visibleFormColumns.map((col) => (
          <input
            key={col.accessor}
            name={col.accessor}
            type={getInputType(col.accessor)}
            placeholder={col.label}
            value={formData[col.accessor] ?? ""}
            onChange={handleChange}
          />
        ))}
        <button type="submit">{editingId ? "Оновити" : "Додати"}</button>
      </form>

      <table>
        <thead>
          <tr>
            {columns.map((col) => (
              <th key={col.accessor}>{col.label}</th>
            ))}
            <th>Дії</th>
          </tr>
        </thead>
        <tbody>
          {data.map((item, idx) => (
            <tr key={`${item.id ?? "no-id"}-${idx}`}>
              {columns.map((col) => (
                <td key={col.accessor}>{formatValue(item[col.accessor], col.accessor)}</td>
              ))}
              <td>
                <button onClick={() => handleEdit(item)}>✏️</button>
                <button onClick={() => handleDelete(item.id)}>🗑️</button>
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}
