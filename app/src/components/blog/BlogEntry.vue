<template>
    <div class="blog-entry">
        <h1>{{ entry.title }}</h1>
        <h3>
            Posted on
            <a href="#">{{ entry.created_at | date }}</a> by
            <a href="#">{{ entry.username }}</a>
        </h3>
        <img class="blog-promo mb1" alt :src="'/images/blog/' + entry.id + '.jpg'" />
        <div class="blog-contents" v-html="entry.content" />
    </div>
</template>

<script>
export default {
    props: {
        entry: {
            type: Object,
            required: true
        }
    }
};
</script>

<style scoped>
.blog-entry {
    max-width: 627px;
    display: inline-block;
    margin: 10px 0;
    padding: 20px;
    background-color: var(--color-grey-medium);
    vertical-align: top;
    border-radius: 20px;
    border: 1px solid var(--color-grey-light);
}

.blog-entry:first-child {
    margin-top: 0;
}

.blog-entry:last-child {
    margin-bottom: 0;
}

.blog-entry h1 {
    color: white;
    text-align: center;
    margin: 0;
    font-weight: 100;
}

.blog-entry h3 {
    color: white;
    font-size: 16px;
    text-align: center;
    font-weight: 100;
    margin: 0 0 16px 0;
}

.blog-contents >>> p:last-child {
    margin-bottom: 0;
}

.blog-promo {
    width: 100%;
    height: auto;
    padding: 0;
    box-shadow: 0 0 4px #222;
}

.blog hr {
    border-color: #666;
    margin-bottom: 16px;
}

.buttons {
    text-align: right;
}
</style>