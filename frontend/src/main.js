// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import store from './store'
import iView from 'iview';
import 'iview/dist/styles/iview.css';
Vue.use(iView);

Vue.config.productionTip = false
router.beforeEach((to, from, next) => {
  // 检测路由配置中是否有requiresAuth这个meta属性
  if (to.matched.some(record => record.meta.requiresAuth)) {
    // 判断是否已登录
    if (store.getters.isLoggedIn) {
      //判断是否有权限访问
      let canAccess = false;
      store.getters.accessList.forEach(access=>{
        if(access.path == to.path && access.level <= store.getters.accessLevel){
          console.info("you can access this location: ",to.path);
          canAccess = true;
          return
        }
      });
      if(canAccess){
        next();      
      }else{
        alert("您无权访问该页面");
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
