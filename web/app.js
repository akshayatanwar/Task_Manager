const KEY = "task-manager.tasks";

const form = document.getElementById("form");
const titleInput = document.getElementById("title");
const priorityInput = document.getElementById("priority");
const listEl = document.getElementById("list");
const emptyEl = document.getElementById("empty");
const summaryEl = document.getElementById("summary");
const clearBtn = document.getElementById("clear");
const filterBtns = document.querySelectorAll("[data-filter]");

let tasks = load();
let filter = "all";

function load() {
  try {
    return JSON.parse(localStorage.getItem(KEY)) || [];
  } catch (e) {
    return [];
  }
}

function save() {
  try {
    localStorage.setItem(KEY, JSON.stringify(tasks));
  } catch (e) {
    /* storage unavailable: tasks still work until the page is closed */
  }
}

function render() {
  const visible = tasks.filter(t =>
    filter === "all" ? true : filter === "done" ? t.done : !t.done
  );

  listEl.innerHTML = "";
  visible.forEach(t => {
    const li = document.createElement("li");
    li.className = t.priority + (t.done ? " done" : "");
    li.dataset.id = t.id;

    const box = document.createElement("input");
    box.type = "checkbox";
    box.checked = t.done;
    box.setAttribute("aria-label", "Mark done");

    const text = document.createElement("span");
    text.className = "text";
    text.textContent = t.title;

    const tag = document.createElement("span");
    tag.className = "tag";
    tag.textContent = t.priority;

    const del = document.createElement("button");
    del.className = "del";
    del.type = "button";
    del.textContent = "×";
    del.setAttribute("aria-label", "Delete task");

    li.append(box, text, tag, del);
    listEl.appendChild(li);
  });

  emptyEl.hidden = visible.length > 0;

  const left = tasks.filter(t => !t.done).length;
  summaryEl.textContent = tasks.length === 0
    ? "No tasks yet"
    : left + " of " + tasks.length + " left to do";

  filterBtns.forEach(b => b.classList.toggle("on", b.dataset.filter === filter));
}

form.addEventListener("submit", e => {
  e.preventDefault();
  const title = titleInput.value.trim();
  if (!title) return;
  tasks.unshift({
    id: Date.now().toString(36) + Math.random().toString(36).slice(2, 6),
    title: title,
    priority: priorityInput.value,
    done: false
  });
  titleInput.value = "";
  save();
  render();
  titleInput.focus();
});

listEl.addEventListener("click", e => {
  const li = e.target.closest("li");
  if (!li) return;
  const id = li.dataset.id;
  if (e.target.classList.contains("del")) {
    tasks = tasks.filter(t => t.id !== id);
  } else if (e.target.type === "checkbox") {
    const t = tasks.find(t => t.id === id);
    if (t) t.done = e.target.checked;
  } else {
    return;
  }
  save();
  render();
});

filterBtns.forEach(b => {
  b.addEventListener("click", () => {
    filter = b.dataset.filter;
    render();
  });
});

clearBtn.addEventListener("click", () => {
  tasks = tasks.filter(t => !t.done);
  save();
  render();
});

render();
