import axios from 'axios'
let base='api/v1'
export const GetCategories = (param) => {return axios.get(`${base}/categories`, {params:param}).then(res=>res.data)}
export const AddCategory = (param, token) => {return axios.post(`${base}/category?token=${token}`, param).then(res=>res.data)}
export const UpdateCategory = (param, token) => {return axios.put(`${base}/category?token=${token}`, param).then(res =>res.data)}
export const DeleteCategory = (param, token) => {return axios.delete(`${base}/category?token=${token}`, {params:param}).then(res=>res.data)}