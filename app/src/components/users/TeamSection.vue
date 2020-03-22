<template>
    <div v-if="group && group.length > 0">
        <h2>{{ roleName }}</h2>
        <div class="team">
            <team-member
                v-for="(user, index) in group"
                :key="index"
                :user="user"
                @remove="$emit('remove', user.id)"
                :isAdmin="isAdmin"
            />
        </div>
    </div>
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
h2 {
    color: white;
    text-align: center;
    margin-bottom: 1rem;
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
    -webkit-justify-content: space-around;
    justify-content: space-around;
    -webkit-align-content: space-around;
    align-content: space-around;
    margin-bottom: 40px;
}
</style>
