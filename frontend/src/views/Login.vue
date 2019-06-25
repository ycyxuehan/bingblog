<template>
  <div id="login-container" class="login-div">
    <!-- <h2 class="login-title">万花师兄的个人博客</h2> -->
    <Form ref="loginForm" :model="loginForm" :rules="rules" @keydown.enter.native="handleSubmit">
      <FormItem prop="userName">
        <Input v-model="loginForm.username" placeholder="请输入用户名">
          <span slot="prepend">
            <Icon :size="16" type="ios-person"></Icon>
          </span>
        </Input>
      </FormItem>
      <FormItem prop="password">
        <Input type="password" v-model="loginForm.password" placeholder="请输入密码">
          <span slot="prepend">
            <Icon :size="14" type="md-lock"></Icon>
          </span>
        </Input>
      </FormItem>
      <FormItem>
        <Button @click="handleSubmit" type="primary" long>登录</Button>
      </FormItem>
    </Form>
  </div>
</template>

<script>
export default {
  name: "LoginContainer",
  data() {
    return {
      rulers:{
          userNameRules: {
        type: Array,
        default: () => {
          return [{ required: true, message: "账号不能为空", trigger: "blur" }];
        }
      },
      passwordRules: {
        type: Array,
        default: () => {
          return [{ required: true, message: "密码不能为空", trigger: "blur" }];
        }
      },
      },
      loginForm: {
        username: "",
        password: ""
      }
    };
  },
  methods: {
    handleSubmit: function() {
      let params = {
        username: this.loginForm.username,
        password: this.loginForm.password
      };
      this.$store
        .dispatch("Login", params)
        .then(() => {
          this.$router.push("/");
        })
        .catch(error => {
          console.log(error);
        });
    }
  }
};
</script>

<style scoped>
.login-div {
  width: 30%;
  max-width: 500px;
  margin-left: 35%;
  margin-top: 150px;
}
.login-title {
  margin-bottom: 1em;
}
</style>