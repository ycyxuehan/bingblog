import axios from 'axios'
let base = 'api/v1'

export const GetArticles = (param) => {return axios.get(`${base}/articles`, {params:param}).then(res => res.data)}
export const GetArticle = (param) => {return axios.get(`${base}/article`, {params:param}).then(res=> res.data)}
export const AddArticle = (data, token) => {return axios.post(`${base}/article?token=${token}`, data).then(res=> res.data)}
export const UpdateArticle = (data, token) => {return axios.put(`${base}/article?token=${token}`, data).then(res => res.data)}
export const DeleteArticle = (data, token) => {return axios.delete(`${base}/article?token=${token}`, {params:data}).then(res=>res.data)}

