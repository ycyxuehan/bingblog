<template>
  <div id="app">
    <Layout :style="{background:'rgb(245, 241, 241)'}">
      <default-navi :theme="theme" :title="title" :logo="logo"></default-navi>
      <div class="content">
        <router-view/>
      </div>
      <Footer class="layout-footer-center" :style="{padding:'0 0 0 0'}">
        <footer-container></footer-container>
      </Footer>
    </Layout>
  </div>
</template>

<script>
import DefaultNavi from "@/components/navi/Default.vue";
import FooterContainer from "@/components/footer/Footer.vue";
export default {
  name: "App",
  components: {
    FooterContainer,
    DefaultNavi
  },
  data() {
    return {
      theme: "",
      logo: "",
      title: "",
    };
  },
  mounted() {
    this.theme = this.$store.getters.headerTheme;
    this.title = this.$store.getters.headerTitle;
    this.logo = this.$store.getters.headerLogo;
  }
};
</script>

<style>
#app {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 0.1em;
  width: 80%;
  margin-left: 10%;
  
  background-color: rgb(245, 241, 241);
}
.content {
  min-height: 780px;
  height: 780px;
  
}
.w-e-text-container{
    height: 480px !important;
}
</style>
