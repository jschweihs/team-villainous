<template>
    <div class="blog-entry" v-if="entry && user">
        <h1>{{ entry.title }}</h1>
        <h3>
            Posted on
            <a href="#">{{ entry.created_at | date }}</a> by
            <a href="#">{{ user.username }}</a>
        </h3>
        <img class="blog-promo" alt :src="'/images/blog/' + entry.id + '.jpg'" />
        <p v-html="entry.preview" />
        <div class="buttons">
            <!--         	<button @click="$emit('goto', entry.id);">
        		<i class="fas fa-arrow-circle-right"></i>
            </button>-->
            <button @click="$emit('edit', entry.id);">
                <i class="far fa-edit"></i>
            </button>
            <button @click="$emit('remove', entry.id);">
                <i class="far fa-trash-alt"></i>
            </button>
        </div>
    </div>
</template>

<script>
export default {
    props: {
        entry: {
            type: Object,
            required: true
        },
        user: {
            type: Object,
            required: true
        }
    },
    created() {
        if (!this.user) {
            this.$emit("get-user", this.entry.user_id);
        }
    }
};
</script>

<style scoped>
button {
    color: var(--color-secondary);
    display: inline;
    background: none;
    border: none;
    cursor: pointer;
    -webkit-touch-callout: none; /* iOS Safari */
    -webkit-user-select: none; /* Safari */
    -khtml-user-select: none; /* Konqueror HTML */
    -moz-user-select: none; /* Firefox */
    -ms-user-select: none; /* Internet Explorer/Edge */
    user-select: none; /* Non-prefixed version, currently
                                  supported by Chrome and Opera */
}

.blog-entry {
    max-width: 627px;
    display: inline-block;
    margin: 20px 0;
    padding: 20px;
    background-color: var(--color-grey-dark);
    vertical-align: top;
    border-radius: 6px;
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

.blog-entry p {
    text-indent: 16px;
    letter-spacing: 0.25px;
}

.blog-promo {
    width: 100%;
    height: auto;
    padding: 0;
    box-shadow: 0 0 4px #222;
}

.blog hr {
    border-color: var(--color-grey-light);
    margin-bottom: 16px;
}

.buttons {
    text-align: right;
}
</style>