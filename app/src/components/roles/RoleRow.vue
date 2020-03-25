<template>
    <div class="button-row" :class="{ alt: index%2==0 }">
        <div class="row-edit" v-if="is_edit">
            <form @submit.prevent="save">
                <input type="text" v-model="name" placeholder="Enter role name..." />
                <input type="submit" value="Save" />
            </form>
        </div>
        <div class="row-read" v-else>
            <div class="role-name">{{ role.name }}</div>
            <div class="role-buttons">
                <button @click="is_edit = true;">
                    <i class="far fa-edit"></i>
                </button>
                <button @click="$emit('remove', role.id);">
                    <i class="far fa-trash-alt"></i>
                </button>
            </div>
        </div>
    </div>
</template>

<script>
export default {
    props: {
        role: {
            type: Object,
            required: true
        },
        index: {
            type: Number,
            required: false
        }
    },
    data() {
        return {
            name: this.role.name,
            is_edit: false
        };
    },
    methods: {
        save() {
            this.$emit("save", { id: this.role.id, name: this.name });
            this.is_edit = false;
        }
    }
};
</script>

<style scoped>
.button-row {
    padding: 10px 0;
    font-size: 24px;
    color: white;
    width: 100%;
    height: 48px;
}

.row-read {
    display: table;
    height: 48px;
    width: 100%;
}
.alt {
    background-color: var(--color-grey-medium);
}

.role-name {
    display: table-cell;
    vertical-align: middle;
    padding-left: 20px;
}

.role-buttons {
    display: table-cell;
    text-align: right;
    vertical-align: middle;
    padding-right: 20px;
}

button {
    color: var(--color-secondary);
    display: inline;
    background: none;
    border: none;
    cursor: pointer;
}

input,
select,
textarea {
    display: inline-block;
    width: calc(75% - 15px);
    margin-left: 10px;
    font-size: 24px;
    padding: 10px;
    border-radius: 8px;
    border: 0;
    -webkit-box-sizing: border-box; /* Safari/Chrome, other WebKit */
    -moz-box-sizing: border-box; /* Firefox, other Gecko */
    box-sizing: border-box; /* Opera/IE 8+ */
    font-family: Nixie One, cursive;
}

input[type="submit"] {
    width: calc(25% - 15px);
    margin: 0 10px;
    color: white;
    background-color: var(--color-secondary);
    border: 0;
    height: 48px;
    cursor: pointer;
    font-family: Nixie One, cursive;
}
</style>
