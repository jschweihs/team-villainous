<template>
    <fieldset v-if="group && group.length > 0">
        <legend>{{ roleName }}</legend>
        <div class="team">
            <team-member
                v-for="(user, index) in group"
                :key="index"
                :user="user"
                @remove="$emit('remove', user.id)"
                :isAdmin="isAdmin"
            />
        </div>
    </fieldset>
</template>

<script>
import TeamMember from "./TeamMember.vue";

export default {
    components: {
        TeamMember
    },
    props: {
        group: {
            type: Array,
            required: true
        },
        roleName: {
            type: String,
            required: true
        },
        isAdmin: {
            type: Boolean,
            required: false
        }
    }
};
</script>

<style scoped>
fieldset {
    padding: 0.5rem 0;
    border: 0;
    border-top: 1px solid var(--color-grey-light);
}

legend {
    color: white;
    font-family: inherit;
    padding: 0 0.5rem;
    text-align: center;
}

.team {
    width: 100%;
    display: -webkit-box;
    display: -moz-box;
    display: -ms-flexbox;
    display: -moz-flex;
    display: -webkit-flex;
    display: flex;
    -webkit-flex-direction: row;
    flex-direction: row;
    flex-wrap: wrap;
    -webkit-justify-content: center;
    justify-content: center;
    -webkit-align-content: space-around;
    align-content: space-around;
    padding-top: 20px;
}
</style>
