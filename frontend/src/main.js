// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import store from './store'
import iView from 'iview';
import 'iview/dist/styles/iview.css';
import 'wangeditor/release/wangEditor.min.css'
import hljs from 'highlight.js';
import 'highlight.js/styles/googlecode.css' //样式文件
// import '../static/css/article.css'  

Vue.directive('highlight',function (el) {
    let blocks = el.querySelectorAll('pre code');
    setTimeout(() =>{
        blocks.forEach((block)=>{
        hljs.highlightBlock(block)
        })
    }, 200)
})
Vue.use(iView);

Vue.config.productionTip = false
router.beforeEach((to, from, next) => {
  // 检测路由配置中是否有requiresAuth这个meta属性
  if (to.matched.some(record => record.meta.requiresAuth)) {
    // 判断是否已登录
    if (store.getters.isLoggedIn) {
      //判断是否有权限访问
      let canAccess = false;
      let accessList = store.getters.accessList;
      for(var index in accessList){
        let now = Date.now().valueOf()
        // console.info(now, now/1000 - store.getters.sessionTimeOut, accessList[index].path == to.path, accessList[index].level <= store.getters.accessLevel)
        // 判断session是否已过期
        if(now/1000 - store.getters.sessionTimeOut >= 0){
          console.info("login timeout")
          //执行登出清理
          store.commit('logout')
          next('/login'); 
          return
        }
        if(accessList[index].path == to.path && accessList[index].level <= store.getters.accessLevel){
          console.info("you can access this location: ",to.path);
          canAccess = true;
          break
        }
      };
      if(canAccess){
        if(to.path == '/login'){
          next('/')
        }else {
          next();
        }
      }else{
        alert("您无权访问该页面，");
      }
      return
    }    
    // 未登录则跳转到登录界面
    next('/login'); 
  } else {
    next() 
  }
})


/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  store,
  components: { App },
  template: '<App/>'
})
