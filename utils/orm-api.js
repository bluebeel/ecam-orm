import axios from 'axios';

const BASE_URL = 'http://localhost:3000';

export {getAllProjects, getProject, deleteProject, getAllTasks, getTask, deleteTask};

function getAllProjects() {
  const url = `${BASE_URL}/projects`;
  return axios.get(url).then(response => response.data);
}

function getProject(title) {
  const url = `${BASE_URL}/projects` + title;
  return axios.get(url).then(response => response.data);
}

function deleteProject(title) {
  const url = `${BASE_URL}/projects` + title;
  return axios.delete(url).then(response => response.data);
}

function getAllTasks(title) {
  const url = `${BASE_URL}/projects` + title + `/tasks`;
  return axios.get(url).then(response => response.data);
}

function getTask(title, id) {
  const url = `${BASE_URL}/projects` + title + `/tasks/` + id;
  return axios.get(url).then(response => response.data);
}

function deleteTask(title, id) {
  const url = `${BASE_URL}/projects` + title + `/tasks/` + id;
  return axios.delete(url).then(response => response.data);
}

