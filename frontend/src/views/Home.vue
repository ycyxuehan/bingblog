<template>
    <div id='home-container' class="home">
        <!-- <Card  class="home-card" shadow> -->
            <div v-if="categories.length == 0 && recommendArticles.length == 0">这家伙很懒，还没有博客</div>
            <recommend-carousel :articles='recommendArticles' v-if="recommendArticles.length > 0"></recommend-carousel>
            <category-card :categories="categories"></category-card>
        <!-- </Card> -->
    </div>
</template>

<script>
    import CategoryCard from '@/components/category/Card.vue'
    import RecommendCarousel from '@/components/recommend/Carousel.vue'
    import {GetArticles} from '@/api/article.js'
    import {GetCategories} from '@/api/category.js'
    export default {
        name: 'HomeContainer',
        components:{CategoryCard, RecommendCarousel},
        data() {
            return {
                categories:[
                ],
                recommendArticles:[
                ],
            }
        },
        methods: {
            getCategories: function(){
                GetCategories().then(res=>{
                    if(res.Code == 0){
                        this.categories = res.Data
                    }
                })
            },
            getRecommendArticles: function() {
                GetArticles({recommend:true}).then(res=>{
                    if(res.Code == 0){
                        this.recommendArticles = res.Data;
                    }
                })
            }
        },
        mounted() {
            this.getCategories()
            this.getRecommendArticles()
        },
    }
</script>

<style lang="" scoped>
    .home {
        background-color: rgb(245, 241, 241);
        padding: 5px;
        height:780px; 

    }
    .home-card {
        white-space: nowrap;
        overflow-y: auto;
        -webkit-overflow-scrolling:auto;
        height:770px; 
        /* box-shadow:5px 5px 5px rgba(0,0,0,.4); */
    }
</style>