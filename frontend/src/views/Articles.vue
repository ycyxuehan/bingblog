<template>
    <div id="articles-container" class="articles">
        <Row v-for="(article, index) in articles" :key="index" class="article-row">
            <Card class="article-card">
                <div  @click="openArticle(article)">
                    <h2>{{article.title}}</h2>
                    <Row>
                        <Col span="20">
                            <p class="outline">{{article.outline}}</p>
                        </Col>
                        <Col span="4">
                            <span class="views"><Icon type="md-eye" />{{article.views}}</span>
                        </Col>
                    </Row>
                </div>
            </Card>
        </Row>
    </div>
</template>

<script>
    import {GetArticles} from '@/api/article.js'
    export default {
        name: 'ArticlesContainer',
        data() {
            return {
                articles:[
                ],
                categoryId:1,
            }
        },
        methods: {
            openArticle: function(a){
                this.$router.push('/article?id=' + a.id + '&category=' + this.categoryId);
            },
            getArticles: function(){
                GetArticles({category:this.categoryId}).then(res=>{
                    if(res.Code == 0){
                        this.articles = res.Data
                    }
                })
            }
        },
        mounted() {
            this.categoryId = this.$route.query.id;
            if(this.categoryId == undefined || this.categoryId < 1){
                this.categoryId = 1
            }
            this.getArticles()
        },
    }
</script>

<style lang="" scoped>
    .articles {
        white-space: nowrap;
        overflow-y: auto;
        -webkit-overflow-scrolling:auto;
        min-height: 780px;
        height: 780px;
    }
    .article-row {
        padding: 0.25em 0.5em 0.25em 0.5em;
    }
    .article-card {
        text-align: left;
    }
    .article-card p {
        margin-left: 2em;
    }
    .outline {
        overflow:hidden;
        text-overflow:ellipsis;
        -o-text-overflow:ellipsis;
        white-space:nowrap;
    }
</style>