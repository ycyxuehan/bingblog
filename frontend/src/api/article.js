import axios from 'axios'
let base = 'api/v1'

export const GetArticles = (param) => {return axios.get(`${base}/articles`, {params:param}).then(res => res.data)}
export const GetArticle = (param) => {return axios.get(`${base}/article`, {params:param}).then(res=> res.data)}
export const AddArticle = (param) => {return axios.post(`${base}/article`, param).then(res=> res.data)}
export const UpdateArticle = (param) => {return axios.put(`${base}/article`, param).then(res => res.data)}
export const DeleteArticle = (param) => {return axios.delete(`${base}/article`, {params:param}).then(res=>res.data)}

