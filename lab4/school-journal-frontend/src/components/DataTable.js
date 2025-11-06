import { useState, useEffect } from "react";
import { fetchData, createData, updateData, deleteData } from "../api";

export default function DataTable({ endpoint, columns }) {
  const [data, setData] = useState([]);
  const [formData, setFormData] = useState({});
  const [editingId, setEditingId] = useState(null);

  useEffect(() => {
    fetchData(endpoint).then(setData).catch(console.error);
  }, [endpoint]);

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

  return (
    <div className="table-page">
      <h2>{endpoint.toUpperCase()}</h2>

      <form className="form" onSubmit={handleSubmit}>
        {columns.map((col) => (
          <input
            key={col.accessor}
            name={col.accessor}
            placeholder={col.label}
            value={formData[col.accessor] || ""}
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
          {data.map((item) => (
            <tr key={item.id}>
              {columns.map((col) => (
                <td key={col.accessor}>{item[col.accessor]}</td>
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
