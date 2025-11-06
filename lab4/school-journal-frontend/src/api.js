const API_URL = process.env.REACT_APP_API_URL;

export async function fetchData(endpoint) {
  const res = await fetch(`${API_URL}/${endpoint}`);
  if (!res.ok) throw new Error("Помилка при отриманні даних");
  return await res.json();
}

export async function createData(endpoint, data) {
  const res = await fetch(`${API_URL}/${endpoint}`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(data),
  });
  if (!res.ok) throw new Error("Помилка при створенні запису");
  return await res.json();
}

export async function updateData(endpoint, id, data) {
  const res = await fetch(`${API_URL}/${endpoint}/${id}`, {
    method: "PUT",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(data),
  });
  if (!res.ok) throw new Error("Помилка при оновленні запису");
  return await res.json();
}

export async function deleteData(endpoint, id) {
  const res = await fetch(`${API_URL}/${endpoint}/${id}`, {
    method: "DELETE",
  });
  if (!res.ok) throw new Error("Помилка при видаленні");
  return await res.json();
}
