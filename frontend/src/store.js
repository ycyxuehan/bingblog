import Vue from 'vue'
import Vuex from 'vuex'
import axios from 'axios'

axios.interceptors.response.use(
  response => {
    return response;
  },
  error => {
    if (error.response) {
      switch (error.response.status) {
        case 401:
        //跳转登录
        this.$store.dispatch(logout);
      }
    }
  }
);
Vue.use(Vuex)

export default new Vuex.Store({
    // ...
    state: {
        headerTheme: 'light',
        headerTitle: '万花师兄的个人博客',
        headerLogo:'static/img/hualuo.png',
        status: '',
        token: localStorage.getItem('token') || '',
        user: "",
        accesslevel: localStorage.getItem('level') || 0,
        accessList: [
          {id: 0, path: "/", level: 1},
          {id: 1, path: "/articles", level: 1},
          {id: 2, path: "/article", level: 1},
          {id: 3, path: "/admin", level: 3},
        ]
    },
    mutations: {
        auth_request(state) {
          state.status = 'loading';
        },
        auth_success(state, loginData) {
          state.status = 'success';
          state.token = loginData.token;
          state.user = loginData.user;
          state.accesslevel = loginData.level;
        },
        auth_error(state) {
          state.status = 'error';
        },
        logout(state) {
          state.status = ''; 
          state.token = '';
          state.accesslevel = 0;
        },
      },
      actions: {
        Login({commit}, params) {
          return new Promise((resolve, reject) => {
          commit('auth_request')
          // 向后端发送请求，验证用户名密码是否正确，请求成功接收后端返回的token值，利用commit修改store的state属性，并将token存放在localStorage中
          axios.post('api/v1/login', params)
            .then(res => {
              res = res.data;
              if(res.Code != 0){
                commit('auth_error')
                localStorage.removeItem('token')
                localStorage.removeItem('level')
                reject(err)
                return
              }
              console.info(params)
              const timeout = res.data.ResultData.timeout;
              const token = res.Data.token;
              const user = res.Data.username;
              const level = res.Data.accesslevel;
              localStorage.setItem('token', token)
              localStorage.setItem('level', level)
              localStorage.setItem('timeout', timeout)
              // 每次请求接口时，需要在headers添加对应的Token验证
              axios.defaults.headers.common['Authorization'] = token
              // 更新token
              commit('auth_success', {token:token, user:user, level:level})
              resolve(res)
            })
            .catch(err => {
              commit('auth_error')
              localStorage.removeItem('token')
              localStorage.removeItem('level')
              reject(err)
            })
          })
        },
        LogOut({ commit, state }) {
          // return new Promise((resolve, reject) => {
          //   axios.get('Logout')
          //     .then(response => {
          //       removeIsLogin()
          //       localStorage.removeItem('loginUsername');
          //       // 移除之前在axios头部设置的token,现在将无法执行需要token的事务
          //       delete axios.defaults.headers.common['Authorization'];
          //       resolve(response)
          //      })
          //     .catch(error => {
          //       reject(error)
          //     })
          // })
          commit('logout')
          localStorage.removeItem('loginUsername');
          // 移除之前在axios头部设置的token,现在将无法执行需要token的事务
          delete axios.defaults.headers.common['Authorization'];
        }
      },
    
    getters: {
        // !!将state.token强制转换为布尔值，若state.token存在且不为空(已登录)则返回true，反之返回false
        isLogin: state => !!state.token,
        headerTheme: state => state.headerTheme,
        headerTitle: state => state.headerTitle,
        authStatus: state => state.status,
        accessLevel: state => state.accesslevel,
        accessList: state => state.accessList,
        token: state => state.token,
        authStatus: state => state.status,
        headerLogo: state => state.headerLogo,
    }
})