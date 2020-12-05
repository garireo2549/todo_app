const table = document.getElementById("table")
getTODOS();

function getTODOS() {
  fetch("http://localhost:80/todos/get").then((response) => {
    if (response.ok) {
      return response.json();
    }
  }).then(data => {
    // todo一覧を作成
    createTable(data)
  }).catch((err) => {
    console.log(err);
  });
}

function createTable(todos) {
  for(t of todos) {
    console.log(t)
    const tr = document.createElement("tr")
    const name = document.createElement("td")
    const todo = document.createElement("td")
    name.textContent = t.Name
    todo.textContent = t.TODO
    tr.appendChild(name)
    tr.appendChild(todo)
    table.appendChild(tr)
  }
}

document.getElementById("button").onclick = function() {
  const name = document.getElementById("name").value;
  const todo = document.getElementById("todo").value;
  const TODO = {
    Name: name,
    TODO: todo
  }
  fetch("http://localhost:80/todos/post", {
    method: "POST",
    headers: {"Content-Type": "application/json"},
    body: JSON.stringify(TODO)
  }).then((response) => {
    if (response.ok) {
      return response.json();
    }
  }).then(data => {
    document.location.reload()
  }).catch((err) => {
    console.log(err);
  });
};
