<template>
    <div id="category-card" >
        <Row class="card-row">
            <Col span="8" v-for="(c, index) in categories" :key="index">
            <Card shadow  class="card-col">
            <div @click="openCategory(c)">
                <h3 class="card-title">{{c.name}}</h3>
                <p class="card-desc">{{c.description}}</p>

            </div>
            </Card>
            </Col>
        </Row>
        <!-- <Row v-for="(c, index) in categories" :key="index" class="card-row">
             <Card class="card">
                <Row>
                    <h3 class="card-title">{{c.name}}</h3>
                </Row>
                <p class="card-desc">{{c.description}}</p>
            </Card>
        </Row> -->
        <!-- <div v-for="(c, index) in categories" :key="index" class="card-col">
            <Card >
                <Row>
                    <h3 class="card-title">{{c.name}}</h3>
                </Row>
                <p class="card-desc">{{c.description}}</p>
            </Card>
        </div> -->
    </div>
</template>

<script>
    export default {
        name: 'CategoryCard',
        props: ["categories"],
        data() {
            return {
                
            }
        },
        methods: {
            openCategory: function (c){
                this.$router.push("/articles?id=" + c.id)
            },
        },
    }
</script>

<style lang="" scoped>
    .card {
        background-color: #515a6e;
        color: #fff;
    }
    .card-col {
        padding: 0.5em;
        display:inline-block;
        height: 80px;
        background-color: #2db7f5;
        margin: 0.5em;
        color: #fff;
        width: 98%;
    }
    .card-title {
        /* padding-left: 0.5em; */
        text-align: left;
    }
    .card-desc {
        padding-left: 2em;
        text-align: left;
    }
    .card-row {
        margin: 0.5em;
        margin-bottom: 0.5em;
    }
</style>