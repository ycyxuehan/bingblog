<template>
  <div id="recommend-carousel" class>
    <!-- <Card shadow class="card-carousel"> -->
    <div @click="openArticle()">
      <Carousel v-model="value" autoplay loopautoplay-speed="3000" class="carousel" :dots="'none'">
        <CarouselItem v-for="(article, index) in articles" :key="index">
            <h1 class="carousel-title">{{article.title}}</h1>
            <!-- <span class="views"><Icon type="md-eye" />{{article.views}}</span> -->
            <p class="carousel-desc">{{article.outline}}</p>
            
        </CarouselItem>
      </Carousel>
    </div>
    <!-- </Card> -->
  </div>
</template>

<script>
export default {
  name: "RecommendCarousel",
  props: ["articles"],
  data() {
    return {
      value: 1
    };
  },
  methods: {
      openArticle: function(){
          let a = this.articles[this.value];
          this.$router.push("/article?id=" + a.id);
      }
  },
};
</script>

<style lang="" scoped>
.carousel {
  /* position: relative; */
  height: 300px;
  /* margin-bottom: 0.5em; */
  background-color: #5cadff;
  /* margin: 1em; */
  color: #000;
  margin: 0 1em 0 1em;
  box-shadow: 3px 3px 3px rgba(70, 153, 248, 0.4);
  /* width: 100%; */
}
.card-carousel {
  padding: 0.5em;
  background-color: #5cadff;
}
.views {
  /* top: 280px; */
  /* left: 90%; */
}
</style>