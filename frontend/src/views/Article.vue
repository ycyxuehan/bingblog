<template>
    <div id='article-container'>
        <Card dis-hover>
            <div slot="title">
                <Row class="article-title-row">
                    <Col span="1">
                        <Button type="primary" @click="backArticles()">返回列表</Button>
                    </Col>
                    <Col span="22" class="article-title">
                        <h2>{{article.title}}</h2>
                    </Col>
                </Row>
            </div>
            <div class="article-content" v-html="vhtml">
            </div>
        </Card>
    </div>
</template>

<script>
    import {GetArticle} from '@/api/article.js'
    let Base64 = require('js-base64').Base64;
    export default {
        name: 'ArticleContainer',
        data() {
            return {
                article:{
                },
                categoryId: 1,
                articleId: 1,
                vhtml: '',
            }
        },
        methods: {
            backArticles: function(){
                if(this.categoryId == undefined || this.categoryId < 1){
                    this.$router.push('/')
                }else {
                    this.$router.push('/articles?id=' + this.categoryId)
                }

            },
            getArticle: function(){
                GetArticle({id:this.articleId,edit:false}).then(res=>{
                    console.info()
                    if(res.Code == 0 && this.categoryId == res.Data.category){
                        this.article = res.Data
                        this.categoryId = this.article.category
                        this.vhtml = Base64.decode(this.article.content)
                    }
                })
            }
        },
        mounted() {
            // this.categoryId = this.$route.query.category;
            this.articleId = this.$route.query.id
            if(this.articleId == undefined || this.articleId < 1){
                this.$router.push('/')
                return
            }
            this.getArticle()
        },
    }
</script>

<style lang="" scoped>
    .article-title-row {
        height: 30px;
    }
    .article-content {
        height: 670px;
        word-wrap: break-word;
        text-align: left;
    }

</style>