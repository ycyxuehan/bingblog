import Vue from 'vue'
import Router from 'vue-router'
import AdminContainer from '@/views/Admin.vue'
import ArticleContainer from '@/views/Article.vue'
import HomeContainer from '@/views/Home.vue'
import LoginContainer from '@/views/Login.vue'
import ArticlesContainer from '@/views/Articles.vue'

Vue.use(Router)

export default new Router({
  mode: 'history',
  base: process.env.BASE_URL,
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeContainer,
      meta:{requiresAuth:false}
    },
    {
      path: '/login',
      name: 'login',
      component: LoginContainer,
      meta:{requiresAuth:false}
    },
    {
      path: '/admin',
      name: 'admin',
      component: AdminContainer,
      meta:{requiresAuth:true}
    },
    {
      path: '/article',
      name: 'article',
      component: ArticleContainer,
      meta:{requiresAuth:false}
    },
    {
      path: '/articles',
      name: 'articles',
      component: ArticlesContainer,
      meta:{requiresAuth:false}
    },
  ]
})
