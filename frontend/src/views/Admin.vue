<template>
    <div id='admin-container'>
        <Card shadow class="admin-card">
            <Row class="admin-row">
                <Col span="2">
                    <span>类别：</span>
                </Col>
                <Col span="20">
                    <Select v-model="categoryId">
                        <Option v-for="(ca, index) in categories" :value="ca.id" :key="index">{{ca.name}}</Option>
                    </Select>
                </Col>
                <Col span="2">
                    <Button type="primary" v-if="category.id != undefined && category.id != 0" @click="showCategoryEditForm()">修改</Button>
                    <Button type="primary" @click="showCategoryEditForm()">新增</Button>
                </Col>
            </Row>
            <Row  class="admin-row">
                <Col span="2">
                    <span>标题：</span>
                </Col>
                <Col span="22">
                    <Input v-model="article.title" placeholder=""></Input>
                </Col>
            </Row>
            <Row  class="admin-row">
                <Col span="2">
                    <span>简述：</span>
                </Col>
                <Col span="22">
                    <Input v-model="article.outline" type="textarea" placeholder=""></Input>
                </Col>
            </Row>
            <Row  class="admin-row">
                <div id='editor' class="editor"></div>
            </Row>
            <Row class="admin-row">
                <Col span="2" offset="22">
                    <Button type="primary" @click="saveArticle()">保存</Button>
                </Col>
            </Row>
        </Card>
        <Modal
            v-model="editCategoryDlgVisible"
            :title="editCategoryDlgTitle"
            @on-ok="saveCategory()"
            >
            <Form :model="category">
                <FormItem label="Name">
                    <Input v-model="category.name" placeholder=""></Input>
                </FormItem>
                <FormItem label="Description">
                    <Input type="textarea" :autosize="{minRows: 2,maxRows: 5}" v-model="category.description" placeholder=""></Input>
                </FormItem>
            </Form>
        </Modal>
    </div>
</template>

<script>
    import Editor from 'wangeditor'
    import 'wangeditor/release/wangEditor.min.css'
    import {AddArticle, GetArticle, UpdateArticle} from '@/api/article.js'
    import {GetCategories, AddCategory, UpdateCategory, GetCategory} from '@/api/category.js'
    let Base64 = require('js-base64').Base64;

    export default {
        name: 'AdminContainer',
        data() {
            return {
                article:{
                    title:"",
                    content:"",
                },
                categories:[
                    {id:1, name:"test"}
                ],
                categoryId: 1,
                /**
                 * @description 设置change事件触发时间间隔
                 */
                changeInterval: {
                    type: Number,
                    default: 200
                },
                editor: null,
                category:{},
                vhtml:'',
                editCategoryDlgVisible: false,
                editCategoryDlgTitle: '',
            }
        },
        methods: {
            getArticle: function(){
                GetArticle({id:this.articleId, edit:true}).then(res=>{
                    console.info(this.articleId , res.Data.id)
                    if(res.Code == 0 && this.articleId == res.Data.id){
                        this.article = res.Data
                        this.category.id = this.article.category
                        this.vhtml = Base64.decode(this.article.content)
                        if(this.editor){
                            this.editor.txt.html(this.vhtml)
                        }
                    }
                })
            },
            getCategories: function(){
                GetCategories().then(res=>{
                    if(res.Code == 0){
                        this.categories = res.Data
                    }
                })
            },
            updateArticle: function() {
                UpdateArticle(this.article, this.$store.getters.token).then(res=>{
                    if(res.Code == 0){
                        this.$Notice.open({
                            title: 'Succeeded',
                            desc: "博客 " + this.article.title + " 已保存。"
                        });
                    }
                })
            },
            addArticle: function(){
                AddArticle(this.article, this.$store.getters.token).then(res=>{
                    if(res.Code == 0){
                        this.$Notice.open({
                            title: 'Succeeded',
                            desc: "博客 " + this.article.title + " 已保存。"
                        });
                    }
                })
            },
            getCategory: function(){
                GetCategory({id:this.categoryId}).then(res=>{
                    if(res.Code == 0){
                        this.category = res.Data
                    }
                })
            },
            addCategory: function(){
                AddCategory(this.category, this.$store.getters.token).then(res=>{
                    if(res.Code == 0){
                        this.$Notice.open({
                            title: 'Succeeded',
                            desc: "分类 " + this.category.name + " 已保存。"
                        });
                        this.getCategories()
                    }
                })
            },
            updateCategory: function(){
                UpdateCategory(this.category, this.$store.getters.token).then(res=>{
                    if(res.Code == 0){
                        this.$Notice.open({
                            title: 'Succeeded',
                            desc: "类别 " + this.category.name + " 已保存。"
                        });
                        this.getCategories()
                    }

                })
            },
            showCategoryEditForm: function(){
                for(var c in this.categories){
                    if(c.id == this.categoryId){
                        this.category = c
                    }
                }
                this.editCategoryDlgTitle = "修改类别" + this.category.name
                this.editCategoryDlgVisible = true
            },
            showCategoryAddForm: function(){
                this.category = {
                    name:"", description:""
                }
                this.editCategoryDlgTitle = "添加一个类别"
                this.editCategoryDlgVisible = true
            },
            saveCategory: function(){
                if(this.category.id == undefined){
                    this.addCategory()
                }else{
                    this.updateCategory()
                }
            },
            saveArticle: function(){
                this.article.category = this.categoryId
                this.article.content = Base64.encode(this.vhtml)
                // this.article.outline = this.vhtml.substring(0,120)
                if(this.article.id == undefined){
                    this.addArticle()
                }else{
                    this.updateArticle()
                }
            }
        },
        mounted() {
            this.editor = new Editor(`#editor`)
            this.editor.customConfig.onchange = (html) => {
                // let text = this.editor.txt.text()
                this.vhtml = html
            }
            // if (this.cache) localStorage.editorCache = html
            // this.$emit('input', this.valueType === 'html' ? html : text)
            // this.$emit('on-change', html, text)
            // }
            this.editor.customConfig.zIndex = 0;
            // this.editor.customConfig.height = "500";
            this.editor.customConfig.onchangeTimeout = this.changeInterval
            // create这个方法一定要在所有配置项之后调用
            this.editor.create()
            // 如果本地有存储加载本地存储内容
            // let html = this.value || localStorage.editorCache
            // if (html) this.editor.txt.html(html)
            this.getCategories()
            this.articleId = this.$route.query.id
            if(this.articleId != undefined){
                this.getArticle()
            }
        },
    }
</script>

<style lang="" scoped>
    .admin-card{
        margin: 0.5em;
    }
    .admin-row {
        padding: 0.5em;
        /* min-height: 30px; */
    }
    .editor {
        height: 520px;
        max-height: 520px;
        text-align: left;
    }
    
    
</style>