import axios from 'axios'
let base='api/v1'
export const GetCategories = (param) => {return axios.get(`${base}/categories`, {params:param}).then(res=>res.data)}
export const AddCategory = (param) => {return axios.post(`${base}/category`, param).then(res=>res.data)}
export const UpdateCategory = (param) => {return axios.put(`${base}/category`, param).then(res =>res.data)}
export const DeleteCategory = (param) => {return axios.delete(`${base}/category`, {params:param}).then(res=>res.data)}