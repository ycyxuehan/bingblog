<template>
  <div id="default-navi" class="navi">
    <Row>
      <Col span="1" style="padding:0;margin:0px;">
        <img :src="logo" alt class="navi-logo" @click="()=>{$router.push('/')}">
      </Col>
      <Col span="5">
        <h1 class="header"  @click="()=>{$router.push('/')}">{{title}}</h1>
      </Col>
      <Col span="1" offset="15"  class="login">
        <div v-if="isLogin">
          <Button type="primary"  @click="()=>{$router.push('/admin')}"><Icon type="ios-construct" />管理</Button>
        </div>
      </Col>
      <Col span="1" class="login">
        <Button type="default" v-if="!isLogin" @click="()=>{$router.push('/login')}">Login</Button>
      </Col>
    </Row>
  </div>
</template>

<script>
export default {
  name: "DefaultNavi",
  props: ["theme", "logo", "title"],
  data() {
    return {
      isLogin: this.$store.getters.isLogin,
    };
  },
};
</script>

<style scoped>
.navi {
  /* height: 100px; */
  padding: 0;
  margin: 0;
  /* background-color: #fff; */
  background-color: rgb(245, 241, 241);

  /* position: fixed; */
  /* width: 80%; */
  height: 84px;
}
.navi-logo {
  padding-top: 2px;
  padding-left: 10px;
  height: 79px;
  width: 70px;
}
.header {
  margin-top: 1em;
}
.login {
  padding-top: 38px;
}
</style>